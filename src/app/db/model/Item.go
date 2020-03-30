package model

import (
	"time"
)

type Item struct{
	ID	uint32 `gorm:"primary_key;AUTO_INCREMENT;"`
	Title string `gorm:type:varchar(100)`
	Url string `gorm:type:varchar(200)`
	ImageUrl string `gorm:type:varchar(200)`
	TypeDomain string `gorm:type:varchar(20)`
	TypeFilter string `gorm:type:varchar(20)`
	CommentNum int16
	Date     time.Time `gorm:"index"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"default:current_time"`
	DeletedAt *time.Time
}
