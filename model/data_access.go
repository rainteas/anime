package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// DataAccess encapsulates database operations using sqlx
type DataAccess struct {
	db *sqlx.DB
}

// NewDataAccess creates a new instance of DataAccess with a given database connection
func NewDataAccess(driverName string, dataSourceName string) (*DataAccess, error) {
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DataAccess{db: db}, nil
}
