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
	Id        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	Profile   string    `json:"profile"`
	Role      int       `json:"role"`
}

type Category struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
}

type ShortStory struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Title      string    `gorm:"type:varchar(255)" json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint      `json:"authorID"`
	CategoryID uint      `json:"categoryID"`
	Author     Author    `gorm:"foreignKey:AuthorID;references:Id"`
	Category   Category  `gorm:"foreignKey:CategoryID;references:Id"`
}
