package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_Init(t *testing.T) {
	pwd,_ := os.Getwd()
	fmt.Println(pwd)
	//获取文件或目录相关信息
	fileInfoList,err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		fmt.Println(fileInfoList[i].Name())  //打印当前文件或目录下的文件或目录名
	}

}