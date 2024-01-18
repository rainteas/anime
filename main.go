package main

import (
	"anme/cfg"
	"anme/download"
	"anme/dto"
	"anme/model"
	"anme/rss"
	"strconv"
	"strings"

	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	name := "./config.ini"
	if len(args) != 0 {
		name = args[0]
	}
	config := cfg.NewConfig(name)
	log.SetPrefix("二次元欢迎你:")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	//写入到指定日志文件
	f, err := os.OpenFile(config.LogUrl, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个多写入器，将日志同时写入文件和控制台
	multiWriter := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multiWriter)

	// ==初始化数据库===
	access, err := model.NewDataAccess("mysql", config.MysqlUrl)
	rssMeta := model.NewRssMetaRepo(access)
	itemRepository := model.NewItemRepository(access)
	// ==数据库===

	// ==初始化下载驱动==
	downloader, err := download.NewTransmissionDownloader(&download.DownloaderConfig{
		URL:         config.TransmissionUrl,
		Username:    config.TransmissionUser,
		Password:    config.TransmissionPasswd,
		ProxyEnable: config.ProxyEnable,
		ProxyURL:    config.ProxyUrl,
	}, itemRepository)

	if err != nil {
		log.Fatalln("下载驱动初始化失败", err)
	}
	// ==下载驱动初始化完成==

	seasons := strings.Split(config.Season, ",")

	// ==通过数据库查询获取订阅地址==
	subscriptionAddress, err := rssMeta.GetBySessions(seasons)
	if err != nil {
		log.Fatal("获取订阅地址失败:", err)
		return
	}

	// ==订阅驱动==
	newRss := rss.NewSubscriptionTools(config)

	// ==获取订阅动漫数据==
	subscribeToDataOnline, err := newRss.FetchRSSList(subscriptionAddress)

	if err != nil {
		log.Fatalln("获取订阅信息失败:", err)
	}

	// 标记已经更新的动漫
	var updatedAnime string

	for _, info := range subscribeToDataOnline {
		listOfDatabaseAnime, err := itemRepository.GetByTitle(info.AnimeNameAndNumber)
		if err != nil {
			log.Fatalln("查询失败:", err)
			return
		}
		if len(listOfDatabaseAnime) != 0 {
			if info.AnimeName != updatedAnime {
				log.Printf("%s 没有更新 再等等吧", info.AnimeName)
			}
			updatedAnime = info.AnimeName
			continue
		}
		dir := config.TransmissionDownloadDir + info.Season + "/" + info.AnimeName + "/"
		torrent, err := downloader.TorrentByPath(info.DownloadUrl)
		if err != nil {
			log.Fatal("TorrentByPath : ", err)
		}
		torrents, err := downloader.AnimeByTorrents(torrent, dir)
		if err != nil {
			log.Fatalf("下载失败: 目录：%s 文件：%s 错误：%v", dir, info.AnimeName, err)
		}
		log.Printf("下载成功: 目录：%s 文件：%s", dir, info.AnimeNameAndNumber)

		toItem := rssItemToItem(info)
		toItem.TorrentHashString = torrents["hashString"]
		toItem.TorrentID, _ = strconv.Atoi(torrents["id"])
		toItem.TorrentName = torrents["name"]
		_, err = itemRepository.Insert(toItem)
		if err != nil {
			log.Fatalln("数据库挂了", err, info.AnimeName)
			return
		}

	}

}

// rss中的Item结构体转换成mode中的Item
func rssItemToItem(dto *dto.RssAnimeInfo) *model.Item {
	return &model.Item{
		ID:          0,
		Title:       dto.AnimeNameAndNumber,
		Link:        dto.DownloadUrl,
		Description: dto.Description,
		PubDate:     dto.PubDate,
		Download:    0,
	}
}
