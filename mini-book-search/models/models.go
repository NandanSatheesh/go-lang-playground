package models

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	AuthorName        string
	AuthorNationality string
}

type Publisher struct {
	gorm.Model
	PublisherName string
}

type Book struct {
	gorm.Model
	BookTitle   string
	BookYear    uint
	Author      Author `gorm:"foreignkey:AuthorId"`
	AuthorId    uint
	Publisher   Publisher `gorm:"foreignkey:PublisherId"`
	PublisherId uint
}
