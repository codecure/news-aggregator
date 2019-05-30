package database

import (
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // noqa
)

// GetDB used to connect to the database
func GetDB(dbDir string) *gorm.DB {
	if dbDir == "" {
		panic("Database directory not provided")
	}
	db, err := gorm.Open("sqlite3", path.Join(dbDir, "aggregator.db"))
	if err != nil {
		panic("Failed to connect to the database")
	}
	return db
}
