package model

// rss下载历史
type RssMeta struct {
	Id        int    `db:"id"`
	AnimeName string `db:"anime_name"`
	Season    string `db:"season"`
	Url       string `db:"url"`
}

// RssMetaRepository 定义了 'rss_meta' 表的数据访问接口
type RssMetaRepository interface {
	GetAll() ([]RssMeta, error)
	GetByID(id int) (*RssMeta, error)
	GetByAnimeName(animeName string) ([]RssMeta, error)
	GetBySession(season string) ([]RssMeta, error)
	GetBySessions(seasons []string) ([]RssMeta, error)
}
