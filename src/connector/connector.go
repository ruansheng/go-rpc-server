package connector

import (
	"fmt"
	"handler"
	"net"
	"strconv"
)

type Connector struct {
	Host string
	Port int
}

func (conn *Connector) Run() {
	server_conn, err := net.Listen("tcp", conn.Host+":"+strconv.Itoa(conn.Port))
	if err != nil {
		fmt.Println(err)
	}
	for {
		client_conn, err := server_conn.Accept()
		if err != nil {
			continue
		}
		go handler.HandleConnection(&client_conn)
	}
}
