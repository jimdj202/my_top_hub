package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB

}

var MyDB *DB

func InitDB(url string){
	MyDB = NewClient(url)
}

func NewClient(url string) *DB{
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", url)
	if err != nil {
		fmt.Println("Connect server error",err)
		return nil
	}

	return &DB{db}
}

func GetMyDB() *DB{
	return MyDB
}

func (d *DB) GetGormDB() *gorm.DB{
	return MyDB.db
}



