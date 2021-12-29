package models

import "github.com/jinzhu/gorm"

type Articles struct {
	gorm.Model
	Title    string `gorm:"not null" form:"title" json:"title"`
	Content  string `gorm:"not null" form:"content" json:"content"`
	Category string `gorm:"not null" form:"category" json:"category"`
	Status   string `gorm:"not null" form:"status" json:"status"`
}

func (Articles) TableName() string {
	return "tbl_article"
}
