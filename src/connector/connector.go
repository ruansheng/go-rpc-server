package connector

import (
	. "entry"
	"fmt"
	"handler"
	"net"
	"strconv"
	"utils/logger"
)

type Connector struct {
	Ctx *Context
}

func (conn *Connector) Run() {
	server_conn, err := net.Listen("tcp", conn.Ctx.GetHost()+":"+strconv.Itoa(conn.Ctx.GetPort()))
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
		go handler.HandleConnection(&client_conn, conn.Ctx)
	}
}
