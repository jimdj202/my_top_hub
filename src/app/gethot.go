package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hub/src/app/db"
	"hub/src/app/db/model"
	"hub/src/app/spiders"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	//"hub/src/app/spiders"
	//"io/ioutil"
	//"log"
	//"os"
	//"reflect"
	//"strings"
)

var(
	SpiderNames []string
)

func main(){
	myDB := db.NewClient("tophub:hWZpDMhBsRMWHDWc@tcp(192.168.176.128:3306)/tophub?charset=utf8&parseTime=True&loc=Local")
	defer myDB.Close()

	pwd,_ := os.Getwd()
	pwd = pwd + "/spiders"
	fmt.Println(pwd)
	//获取文件或目录相关信息
	fileInfoList,err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	for i := range fileInfoList {
		if !fileInfoList[i].IsDir(){
			fileName := fileInfoList[i].Name()
			if strings.HasPrefix(fileName,"spider"){
				continue
			}
			fileName = strings.TrimSuffix(fileName,".go")
			//SpiderNames.PushBack(fileName)
			SpiderNames = append(SpiderNames,fileName)
		}
	}
	SpiderNames = []string{"DouBanMovie"}
	for _,funcName := range SpiderNames {
		reflectValue := reflect.ValueOf(&spiders.Spider{Name: funcName})
		//reflectValueEle := reflectValue.Elem()
		dataType := reflectValue.MethodByName("Get" + funcName)
		data := dataType.Call(nil)
		fmt.Printf("%T",data)
		//fmt.Println(data[0].Type(),data[0].Kind(),data[0].CanInterface())
		originData := data[0].Interface().([]model.Item)
		//fmt.Println(originData)
		//items := []model.Item{}
		for _,v := range originData{
			(&v).Save()
		}
	}

}

func init(){

}