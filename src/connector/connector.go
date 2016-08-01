package connector

import (
	"fmt"
	"handler"
	"net"
	"strconv"
	"utils/logger"
)

type Connector struct {
	Host string
	Port int
}

func (conn *Connector) Run() {
	server_conn, err := net.Listen("tcp", conn.Host+":"+strconv.Itoa(conn.Port))
	if err != nil {
		fmt.Println(err)
		logger.Write("ERROR", err.Error())
	}
	for {
		client_conn, err := server_conn.Accept()
		if err != nil {
			logger.Write("ERROR", err.Error())
			continue
		}
		go handler.HandleConnection(&client_conn)
	}
}
