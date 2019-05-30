package database

import "github.com/jinzhu/gorm"

// NewsItem Common item model for all parsed feeds
type NewsItem struct {
	gorm.Model
	Title   string `gorm:"type:text; not null"`
	Content string `gorm:"type:text"`
	Link    string `gorm:"type:text; not null"`
	Date    string `gorm:"type:datetime"`
	Author  string
	GUID    string `gorm:"unique;not null"`
}
