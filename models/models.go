package models

import (
	"time"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Author struct {
	Id         uint         `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	Name       string       `gorm:"type:varchar(255)" json:"name"`
	Email      string       `gorm:"type:varchar(255);unique_index" json:"email"`
	Password   string       `gorm:"type:varchar(255)" json:"password"`
	Profile    string       `json:"profile"`
	ShortStory []ShortStory `json:"shortStory"`
}

type Category struct {
	Id         uint         `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	Name       string       `gorm:"type:varchar(255)" json:"name"`
	ShortStory []ShortStory `json:"shortStory"`
}

type ShortStory struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Title      string    `gorm:"type:varchar(255)" json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint      `json:"authorID"`
	CategoryID uint      `json:"categoryID"`
}
