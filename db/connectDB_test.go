package db

import (
	"database/sql"
	"testing"
)

func TestSetDB(t *testing.T) {
	mockDB := &sql.DB{}
	SetDB(mockDB)
	if DB != mockDB {
		t.Error("DB is not set correctry")
	}
}

func TestGetDB(t *testing.T) {
	db := GetDB()
	if DB != db {
		t.Error("DB is not set correctry")
	}
}
