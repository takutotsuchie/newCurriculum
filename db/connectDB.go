package db

import (
	"database/sql"
)

var (
	DB *sql.DB
)

// DBインスタンスをセット
func SetDB(db *sql.DB) {
	DB = db
}

// DBインスタンスを取得
func GetDB() *sql.DB {
	return DB
}
