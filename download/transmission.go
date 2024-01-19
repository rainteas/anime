package download

import (
	"anme/dto"
	"anme/model"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Transmission 下载器实现
type TransmissionDownloader struct {
	Config         *DownloaderConfig
	SessionID      string // 存储 Transmission 会话 ID
	ItemRepository model.ItemRepository
}

// NewTransmissionDownloader 创建一个新的 TransmissionDownloader 实例并执行初始化操作。
func NewTransmissionDownloader(config *DownloaderConfig, itemRepository model.ItemRepository) (Downloader, error) {
	downloader := &TransmissionDownloader{
		Config:         config,
		ItemRepository: itemRepository,
	}
	// 在初始化时获取 Transmission 会话 ID
	sessionID, err := downloader.getSessionID()
	if err != nil {
		return nil, err
	}
	downloader.SessionID = sessionID

	return downloader, nil
}

func (td *TransmissionDownloader) AnimeByTorrents(torrent []byte, path string) (map[string]string, error) {
	encodedTorrent := base64.StdEncoding.EncodeToString(torrent)

	transmission, err := td.sendRequestToTransmission("torrent-add", map[string]interface{}{
		"metainfo":     encodedTorrent,
		"download-dir": path, // 指定下载目录
	})

	if err != nil {
		log.Printf("发送请求时出错: %v", err)
		return nil, err
	}
	if transmission.Result != "success" {
		log.Printf("添加种子时出错: %s", transmission.Result)
		return nil, fmt.Errorf("添加种子时出错: %s", transmission.Result)
	}
	torrentInfo := make(map[string]string, 5)
	torrentInfo["hashString"] = transmission.Arguments.TorrentAdded.HashString
	torrentInfo["id"] = fmt.Sprintf("%d", transmission.Arguments.TorrentAdded.Id)
	torrentInfo["name"] = transmission.Arguments.TorrentAdded.Name
	torrentInfo["result"] = transmission.Result
	return torrentInfo, nil
}

func (td *TransmissionDownloader) AnimeByMagnet(magnet string, path string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (td *TransmissionDownloader) TorrentByPath(path string) ([]byte, error) {
	var client *http.Client
	if td.Config.ProxyEnable {
		// 解析代理 URL
		proxy, err := url.Parse(td.Config.ProxyURL)
		if err != nil {
			return nil, err
		}
		// 配置 HTTP 客户端使用代理
		transport := &http.Transport{Proxy: http.ProxyURL(proxy)}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}

	// 发送 HTTP 请求
	resp, err := client.Get(path)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("TorrentByPath body流未正常关闭", err)
		}
	}(resp.Body)

	// 读取响应体
	return io.ReadAll(resp.Body)
}

// Response 代表从 Transmission API 收到的响应。
type Response struct {
	Result    string `json:"result"`
	Arguments struct {
		TorrentAdded struct {
			HashString string `json:"hashString" db:"hash_string"`
			Id         int    `json:"id"`
			Name       string `json:"name"`
		} `json:"torrent-added"`
		Torrents []struct {
			ID     int               `json:"id"`
			Name   string            `json:"name"`
			Status dto.TorrentStatus `json:"status"`
		} `json:"torrents"`
	} `json:"arguments"`
}

// sendRPCRequest 发送请求到 Transmission 的 JSON-RPC API。
func (td *TransmissionDownloader) sendRequestToTransmission(method string, arguments map[string]interface{}) (Response, error) {
	client := &http.Client{}

	params := map[string]interface{}{
		"method":    method,
		"arguments": arguments,
	}

	payload, err := json.Marshal(params)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", td.Config.URL, bytes.NewBuffer(payload))
	if err != nil {
		return Response{}, err
	}

	req.SetBasicAuth(td.Config.Username, td.Config.Password)
	req.Header.Set("Content-Type", "application/json")

	// 如果存在会话 ID，将其添加到请求头
	if td.SessionID != "" {
		req.Header.Set("X-Transmission-Session-Id", td.SessionID)
	}

	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("sendRPCRequest body流未正常关闭", err)
		}
	}(resp.Body)

	// 处理响应头，提取会话 ID
	if td.SessionID == "" {
		td.SessionID = resp.Header.Get("X-Transmission-Session-Id")
	}

	// 处理响应，根据需要进行处理
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

// getSessionID 在初始化时获取 Transmission 会话 ID。
func (td *TransmissionDownloader) getSessionID() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", td.Config.URL, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(td.Config.Username, td.Config.Password)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("getSessionID body流未正常关闭", err)
		}
	}(resp.Body)

	// 提取会话 ID
	return resp.Header.Get("X-Transmission-Session-Id"), nil
}

func (td *TransmissionDownloader) StatusByTorrentID(torrentID string) (string, error) {
	transmission, err := td.sendRequestToTransmission("torrent-get", map[string]interface{}{
		"ids":    []string{torrentID},
		"fields": []string{"status"},
	})
	if err != nil {
		log.Println("Error sending request:", err)
		return "", err
	}
	if transmission.Result != "success" {
		log.Println("Error getting torrent status:", transmission.Result)
		return "", fmt.Errorf("Error getting torrent status: %s", transmission.Result)
	}
	return fmt.Sprintf("%s", transmission.Arguments.Torrents[0].Status), nil
}
