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

func (conn *Connector) Run(host string, port int) {
	fmt.Println("Run")
	server_conn, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
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
