package rss

import (
	"anme/cfg"
	"anme/dto"
	"anme/model"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync/atomic"
)

type SubscriptionTools struct {
	Config *cfg.Config
}

func NewSubscriptionTools(config *cfg.Config) *SubscriptionTools {
	return &SubscriptionTools{Config: config}
}

// FetchRSS 获取 RSS 数据
func (r *SubscriptionTools) FetchRSS(animationSubscriptionAddress string) (*dto.RSS, error) {
	var client *http.Client
	if r.Config.ProxyEnable {
		// 解析代理 URL
		proxy, err := url.Parse(r.Config.ProxyUrl)
		if err != nil {
			return nil, err
		}
		// 配置 HTTP 客户端使用代理
		transport := &http.Transport{Proxy: http.ProxyURL(proxy)}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}

	var resp *http.Response
	var err error
	// 重试次数
	retry := 3
	for retry > 0 {
		resp, err = client.Get(animationSubscriptionAddress)
		if err != nil {
			log.Println("Error fetching RSS:", err, "retry:", retry)
			retry--
			continue
		} else {
			break
		}
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("Error closing response body:", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss dto.RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, err
	}

	return &rss, nil
}

// FetchRSSList 批量获取订阅信息
func (r *SubscriptionTools) FetchRSSList(subscriptionAddress []model.RssMeta) ([]*dto.RssAnimeInfo, error) {
	var rssAnimeInfoList []*dto.RssAnimeInfo
	readSubscriptionProgress := 0.0
	for _, meta := range subscriptionAddress {
		rssList, err := r.FetchRSS(meta.Url)
		if err != nil {
			log.Println("name:"+meta.AnimeName+"url:"+meta.Url+"Error fetching RSS:", err)
		} else {
			readSubscriptionProgress = readSubscriptionProgress + 1
			log.Printf("订阅获取进度: %.2f \n", readSubscriptionProgress/float64(len(subscriptionAddress))*100)
		}

		info := r.ConvertRssToAnimeInfo(meta.AnimeName, meta.Season, rssList)

		// 将rssAnimeInfoList和info合并成新的切片
		rssAnimeInfoList = append(rssAnimeInfoList, info...)

	}
	return rssAnimeInfoList, nil
}

// FetchRSSListByGoroutine 开启线程获取订阅
func (r *SubscriptionTools) FetchRSSListByGoroutine(subscriptionAddress []model.RssMeta) ([]*dto.RssAnimeInfo, error) {
	var rssAnimeInfoList []*dto.RssAnimeInfo
	numberOfTimesToObtainData := 0
	readSubscriptionProgress := int32(0)
	ch := make(chan []*dto.RssAnimeInfo, 20)
	for _, meta := range subscriptionAddress {
		go func(meta model.RssMeta) {
			rssList, err := r.FetchRSS(meta.Url)
			if err != nil {
				log.Printf("name: %s url: %s Error fetching RSS: %v", meta.AnimeName, meta.Url, err)
				ch <- nil
				return
			}
			atomic.AddInt32(&readSubscriptionProgress, 1)
			log.Printf("订阅获取进度: %.2f \n", float64(readSubscriptionProgress)/float64(len(subscriptionAddress))*100)
			ch <- r.ConvertRssToAnimeInfo(meta.AnimeName, meta.Season, rssList)
		}(meta)
	}

	for range subscriptionAddress {
		info := <-ch
		if info != nil {
			rssAnimeInfoList = append(rssAnimeInfoList, info...)
			numberOfTimesToObtainData++
		}
	}
	close(ch)
	return rssAnimeInfoList, nil
}

// ConvertRssToAnimeInfo *dto.RSS转换为*dto.RssAnimeInfo
func (r *SubscriptionTools) ConvertRssToAnimeInfo(animeName string, season string, rssList *dto.RSS) []*dto.RssAnimeInfo {
	var animeInfo []*dto.RssAnimeInfo
	for _, rssItem := range rssList.Channel.Items {
		animeInfo = append(animeInfo, &dto.RssAnimeInfo{
			AnimeName:          animeName,
			AnimeNameAndNumber: rssItem.Title,
			Season:             season,
			DownloadUrl:        rssItem.Enclosures[0].URL,
			Description:        rssItem.Description,
			PubDate:            rssItem.PubDate,
		})
	}
	return animeInfo
}
