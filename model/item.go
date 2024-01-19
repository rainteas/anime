package model

// Item 是 'items' 表的领域对象
type Item struct {
	ID                int    `db:"id"`
	Title             string `db:"title"`
	Link              string `db:"link"`
	Description       string `db:"description"`
	PubDate           string `db:"pub_date"`
	Download          int    `db:"download"`
	TorrentHashString string `db:"torrent_hash_string"`
	TorrentID         int    `db:"torrent_id"`
	TorrentName       string `db:"torrent_name"`
}

// ItemRepository 定义了 'items' 表的数据访问接口
type ItemRepository interface {
	GetAll() ([]Item, error)
	GetByID(id int) (*Item, error)
	GetByTitle(title string) ([]Item, error)
	Insert(item *Item) (int64, error)
}
