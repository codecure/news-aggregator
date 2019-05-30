package database

import (
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // noqa
)

// GetDB used to connect to the database
func GetDB(dbPath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", path.Join(dbPath, "aggregator.db"))
	if err != nil {
		panic("Failed to connect to the database")
	}
	return db
}
