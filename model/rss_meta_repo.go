package model

import "github.com/jmoiron/sqlx"

type rssMetaRepository struct {
	da *DataAccess
}

func NewRssMetaRepo(da *DataAccess) RssMetaRepository {
	return &rssMetaRepository{da: da}
}

// GetAll 检索所有 'rss_meta' 表中的记录
func (ir *rssMetaRepository) GetAll() ([]RssMeta, error) {
	var rssMetas []RssMeta
	err := ir.da.db.Select(&rssMetas, "SELECT * FROM rss_meta")
	if err != nil {
		return nil, err
	}
	return rssMetas, nil
}

// GetByID 通过 ID 从 'rss_meta' 表检索单个记录
func (ir *rssMetaRepository) GetByID(id int) (*RssMeta, error) {
	var rssMeta RssMeta
	err := ir.da.db.Get(&rssMeta, "SELECT * FROM rss_meta WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return &rssMeta, nil
}

// 通过animeName从'rss_meta'表检索多个记录
func (ir *rssMetaRepository) GetByAnimeName(animeName string) ([]RssMeta, error) {
	var rssMetas []RssMeta
	err := ir.da.db.Select(&rssMetas, "SELECT * FROM rss_meta WHERE anime_name =?", animeName)
	if err != nil {
		return nil, err
	}
	return rssMetas, nil
}

// 通过session从'rss_meta'表检索多个记录
func (ir *rssMetaRepository) GetBySession(season string) ([]RssMeta, error) {
	var rssMetas []RssMeta
	err := ir.da.db.Select(&rssMetas, "SELECT * FROM rss_meta WHERE season = ?", season)
	if err != nil {
		return nil, err
	}
	return rssMetas, nil
}

// 通过多个session从'rss_meta'表检索多个记录 使用sqlx.in
func (ir *rssMetaRepository) GetBySessions(seasons []string) ([]RssMeta, error) {
	var rssMetas []RssMeta
	query, args, err := sqlx.In("SELECT * FROM rss_meta WHERE season IN (?)", seasons)
	if err != nil {
		return nil, err
	}
	err = ir.da.db.Select(&rssMetas, query, args...)
	if err != nil {
		return nil, err
	}
	return rssMetas, nil
}
