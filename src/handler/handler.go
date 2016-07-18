package handler

import (
	"fmt"
	"net"
	//"conf"
)

func HandleConnection(conn *net.Conn) {
	fmt.Println(conn)
	defer (*conn).Close()
	buf := make([]byte, 512)
	for {
		n, err := (*conn).Read(buf)
		if err == nil {
			fmt.Println(string(buf[:n]))
		} else {
			break
		}
	}
	//b := conf.GetConf()
	//fmt.Println(b["test"].ServiceUri)
}
