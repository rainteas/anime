package dto

type DownloadByTorrentsBean struct {
	Arguments struct {
		TorrentAdded struct {
			HashString string `json:"hashString" db:"hash_string"`
			Id         int    `json:"id"`
			Name       string `json:"name"`
		} `json:"torrent-added"`
	} `json:"arguments"`
	Result string `json:"result"`
}
