package download

// 下载配置结构
type DownloaderConfig struct {
	URL         string
	Username    string
	Password    string
	ProxyURL    string
	ProxyEnable bool
	// 其他配置项
}

// 下载器接口
type Downloader interface {
	// 通过种子下载动漫
	AnimeByTorrents(torrent []byte, path string) (map[string]string, error)

	// 通过磁力连接下载动漫
	AnimeByMagnet(magnet string, path string) (bool, error)

	// 通过地址下载种子
	TorrentByPath(path string) ([]byte, error)

	// 通过torrentID获取任务状态
	StatusByTorrentID(torrentID string) (string, error)
}
