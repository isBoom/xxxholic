package model

import "github.com/jinzhu/gorm"

// Comment 评论模型
type Comment struct {
	gorm.Model
	UserId       uint   // `sql:"type:integer REFERENCES users(id) on update no action on delete no action"`
	VideoId      uint   // `sql:"type:integer REFERENCES videos(id) on update no action on delete no action"`
	ParentId     uint   `gorm:"not null`
	FirstId      uint   `gorm:"not null`
	Content      string `gorm:"not null`
	ParentUserId uint   `gorm:"not null`
}
