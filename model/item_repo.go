package model

// itemRepository 实现了 ItemRepository 接口
type itemRepository struct {
	da *DataAccess
}

func NewItemRepository(da *DataAccess) ItemRepository {
	return &itemRepository{da: da}
}

// GetAll 检索所有 'items' 表中的记录
func (ir *itemRepository) GetAll() ([]Item, error) {
	var items []Item
	err := ir.da.db.Select(&items, "SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	return items, nil
}

// GetByID 通过 ID 从 'items' 表检索单个记录
func (ir *itemRepository) GetByID(id int) (*Item, error) {
	var item Item
	err := ir.da.db.Get(&item, "SELECT * FROM items WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// 通过title从'items'表检索多个记录
func (ir *itemRepository) GetByTitle(title string) ([]Item, error) {
	var items []Item
	err := ir.da.db.Select(&items, "SELECT * FROM items WHERE title=?", title)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (ir *itemRepository) Insert(item *Item) (int64, error) {
	result, err := ir.da.db.NamedExec("INSERT INTO items (title, link, description, pub_date, download, torrent_hash_string, torrent_id, torrent_name) VALUES (:title, :link, :description, :pub_date, :download, :torrent_hash_string, :torrent_id, :torrent_name)", item)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
