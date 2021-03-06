package main

import (
	"container/list"
	"fmt"
	"hub/src/app/db/model"
	"hub/src/app/spiders"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	_ "testing"
)

func Test_Init(t *testing.T) {
	SpiderNames = list.New()
	fmt.Printf("%T",SpiderNames)
	pwd,_ := os.Getwd()
	pwd = pwd + "/spiders"
	fmt.Println(pwd)
	//获取文件或目录相关信息
	fileInfoList,err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		if !fileInfoList[i].IsDir(){
			fileName := fileInfoList[i].Name()
			if strings.HasSuffix(fileName,"spider"){
				continue
			}
			fileName = strings.TrimSuffix(fileName,".go")
			SpiderNames.PushBack(fileName)
		}
	}
	fmt.Println(SpiderNames)

	for i := SpiderNames.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
		reflectValue := reflect.ValueOf(spiders.Spider{})
		dataType := reflectValue.MethodByName("Get" + i.Value.(string))
		data := dataType.Call(nil)
		fmt.Printf("%T",data)
	}

}

func Test_Migrate(t *testing.T)  {
		db1 := NewClient("tophub:hWZpDMhBsRMWHDWc@tcp(192.168.176.128:3306)/tophub?charset=utf8&parseTime=True&loc=Local")
		db1.GetGormDB().AutoMigrate(&model.Item{})
		fmt.Println(db1)

}
