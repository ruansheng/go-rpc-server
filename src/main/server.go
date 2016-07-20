package main

import (
	"connector"
	"fmt"
)

const (
	host = "127.0.0.1"
	port = 8080
)

func main() {
	fmt.Println("main")
	conn := connector.Connector{host, port}
	conn.Run(host, port)
}
