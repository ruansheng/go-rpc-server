package main

import (
	"conf"
	"fmt"
)

func main() {
	b := conf.GetConf()
	fmt.Println(b["test"].ServiceUri)
}
