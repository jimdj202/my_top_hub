package main

import (
	"container/list"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"hub/src/app/db"
	//"hub/src/app/spiders"
	//"io/ioutil"
	//"log"
	//"os"
	//"reflect"
	//"strings"
)

var(
	SpiderNames *list.List
)

func main(){
	db.NewCient()

	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	fmt.Println(db,err)

	//SpiderNames = list.New()
	//
	//pwd,_ := os.Getwd()
	//pwd = pwd + "/spiders"
	//fmt.Println(pwd)
	////获取文件或目录相关信息
	//fileInfoList,err := ioutil.ReadDir(pwd)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i := range fileInfoList {
	//	if !fileInfoList[i].IsDir(){
	//		fileName := fileInfoList[i].Name()
	//		if strings.HasSuffix(fileName,"spider"){
	//			continue
	//		}
	//		fileName = strings.TrimSuffix(fileName,".go")
	//		SpiderNames.PushBack(fileName)
	//	}
	//}
	//
	//for i := SpiderNames.Front(); i != nil; i = i.Next() {
	//	fmt.Println(i.Value)
	//	reflectValue := reflect.ValueOf(&spiders.Sipder{})
	//	//reflectValueEle := reflectValue.Elem()
	//	dataType := reflectValue.MethodByName("Get" + i.Value.(string))
	//	data := dataType.Call(nil)
	//	fmt.Printf("%T",data)
	//}
}

func init(){

}