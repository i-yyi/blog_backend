package model

import "time"

type Post struct {
	ID         uint      `gorm:"primaryKey" json:"id,omitempty"`
	Title      string    `gorm:"not null" json:"title,omitempty"`
	Content    string    `gorm:"not null" json:"content,omitempty"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
}

type Comment struct {
	ID         uint      `gorm:"primaryKey" json:"id,omitempty"`
	PostID     uint      `gorm:"not null" json:"post_id,omitempty"`
	Content    string    `gorm:"not null" json:"content,omitempty"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
}
