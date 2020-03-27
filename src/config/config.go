package config

import (
	"fmt"
	"os"

)

var(
	SpiderNames []string
)

func init() {
	pwd,_ := os.Getwd()
	fmt.Println(pwd)
}
