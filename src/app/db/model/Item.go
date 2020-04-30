package model

import (
	"github.com/jinzhu/gorm"
	"hub/src/app/db"
	"time"
)

type Item struct{
	//ID	uint32 `gorm:"type:BIGINT"`
	Index int `gorm:"type:SMALLINT;comment:'类型内部排序'"`
	Title string `gorm:"type:varchar(100);comment:'标题'"`
	Url string `gorm:"type:varchar(200);primary_key;comment:'文章链接'"`
	ImageUrl string `gorm:"type:varchar(200);comment:'图片链接'"`
	TypeDomainID string `gorm:"type:varchar(20);index;comment:'分类ID'"`
	TypeDomain string `gorm:"type:varchar(20);comment:'分类显示'"`
	TypeFilter string `gorm:"type:varchar(20);comment:'综合,科技,娱乐,社区,购物,财经,博客等'"`
	CommentNum int `gorm:"type:BIGINT;comment:'评论数量'"`
	Desc string `gorm:"type:varchar(400);comment:'描述'"`
	Extra string `gorm:"type:varchar(400);comment:'额外'"`
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
