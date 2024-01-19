package download

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Aria2Downloader Aria2 下载器实现
type Aria2Downloader struct {
	Config DownloaderConfig
	Secret string
}

// JSONRPCRequest JSON-RPC 请求结构
type JSONRPCRequest struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      string        `json:"id"`
	JsonRPC string        `json:"jsonrpc"`
}

// JSONRPCResponse JSON-RPC 响应结构
type JSONRPCResponse struct {
	Result interface{} `json:"result"`
	Error  *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID      string `json:"id"`
	JsonRPC string `json:"jsonrpc"`
}

func (ad *Aria2Downloader) AnimeByTorrents(torrent []byte, path string) (map[string]string, error) {
	// 在这里实现通过 Aria2 下载动漫的逻辑
	// 使用 a.Config 可以获取下载器配置

	// 示例：调用 Aria2 的 addTorrent 方法
	method := "aria2.addTorrent"
	params := []interface{}{torrent, []string{}, map[string]interface{}{"dir": path}}
	response, err := ad.sendRPCRequest(method, params)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"id": response.Result.(string),
	}, nil
}

func (ad *Aria2Downloader) AnimeByMagnet(magnet string, path string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (ad *Aria2Downloader) TorrentByPath(path string) ([]byte, error) {
	// 使用 a.Config 可以获取下载器配置

	// 示例：调用 Aria2 的 addUri 方法
	method := "aria2.addUri"
	params := []interface{}{[]string{path}, map[string]interface{}{}}
	response, err := ad.sendRPCRequest(method, params)
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("Download started: GID %s", response.Result.(string))), nil
}
func (ad *Aria2Downloader) sendRPCRequest(method string, params []interface{}) (*JSONRPCResponse, error) {
	// 构建 JSON-RPC 请求
	request := JSONRPCRequest{
		Method:  method,
		Params:  params,
		ID:      "1",
		JsonRPC: "2.0",
	}

	// 将请求结构体转为 JSON 字符串
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// 发送 HTTP POST 请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", ad.Config.URL, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	// 添加 Aria2 RPC 密钥
	req.Header.Set("Authorization", "Bearer "+ad.Secret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 响应
	var response JSONRPCResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// 处理错误
	if response.Error != nil {
		return nil, fmt.Errorf("aria2 RPC error: %d - %s", response.Error.Code, response.Error.Message)
	}

	return &response, nil
}
func (ad *Aria2Downloader) StatusByTorrentID(torrentID string) (string, error) {
	return "nil", nil
}
