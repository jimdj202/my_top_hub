package config

import (
	"fmt"
	"os"

)

var Server struct {
	DBUrl string
}

func init() {
	pwd,_ := os.Getwd()
	fmt.Println(pwd)
}
