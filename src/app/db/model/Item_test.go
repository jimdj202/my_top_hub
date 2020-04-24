package model

import (
	"fmt"
	"hub/src/app/db"
	"testing"
)

func Test_AutoMigrate(t *testing.T) {
	db1 := db.NewClient("tophub:hWZpDMhBsRMWHDWc@tcp(192.168.176.128:3306)/tophub?charset=utf8&parseTime=True&loc=Local")
	db1.GetGormDB().AutoMigrate(&Item{})
	fmt.Println(db1)

}
