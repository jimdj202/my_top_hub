package db

import (
	"fmt"
	"hub/src/app/db/model"
	"testing"
)

func Test_Connect(t *testing.T) {
	db1 := InitDB("tophub:hWZpDMhBsRMWHDWc@tcp(192.168.176.128:3306)/tophub?charset=utf8&parseTime=True&loc=Local")
	db1.GetGormDB().AutoMigrate(&model.Item{})
	fmt.Println(db1)

}

