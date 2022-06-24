package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id      uint   `gorm:"primarykey;AUTO_INCREMENT"`
	Title   string `gorm:"type:varchar(40)" json:"title"`
	Content string `gorm:"type:varchar(40)" json:"content"`
}
