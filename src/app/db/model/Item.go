package model

import (
	"github.com/jinzhu/gorm"
	"hub/src/app/db"
	"time"
)

type Item struct{
	//ID	uint32 `gorm:"type:BIGINT"`
	Title string `gorm:"type:varchar(100)"`
	Url string `gorm:"type:varchar(200);primary_key"`
	ImageUrl string `gorm:"type:varchar(200)"`
	TypeDomain string `gorm:"type:varchar(20)"`
	TypeFilter string `gorm:"type:varchar(20)"`
	CommentNum int16 `gorm:"type:BIGINT"`
	//Date     time.Time `sql:"index"`
	CreatedAt time.Time `sql:"index"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time
}

func (i *Item)Save() *gorm.DB{
	return db.GetMyDB().GetGormDB().Save(i)
}

//func (i *Item)UpdateOrCreate() *gorm.DB{
//	record := &Item{Url: i.Url}
//	db.GetMyDB().GetGormDB().Find(record)
//}
