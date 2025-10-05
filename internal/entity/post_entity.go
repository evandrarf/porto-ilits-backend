package entity

import "time"

type Post struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

func (Post) TableName() string {
	return "posts"
}