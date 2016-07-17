package main

import (
	"conf"
	"fmt"
)

const (
	host = "127.0.0.1"
	port = 9999
)

func main() {
	b := conf.GetConf()
	fmt.Println(b["test"].ServiceUri)
}
