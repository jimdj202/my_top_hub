package model

import (
	"fmt"
	"hub/src/app/db"
	"regexp"
	"strconv"
	"testing"
)

func Test_AutoMigrate(t *testing.T) {
	db1 := db.NewClient("tophub:hWZpDMhBsRMWHDWc@tcp(192.168.176.128:3306)/tophub?charset=utf8&parseTime=True&loc=Local")
	db1.GetGormDB().AutoMigrate(&Item{})
	fmt.Println(db1)

}

func Test_Re(t *testing.T){
	reg, _ := regexp.Compile("\\d+")
	comNum2 := reg.Find([]byte("10万+"))
	comNum3, _ := strconv.Atoi(string(comNum2))
	t.Log(comNum3)
}