package main

import (
	"connector"
	"flag"
	"strconv"
	"utils/logger"
)

const (
	host = "127.0.0.1"
	port = 8080
)

func main() {
	hostp := flag.String("host", host, "input host desc")
	portp := flag.Int("port", port, "input port desc")
	flag.Parse()
	conn := connector.Connector{*hostp, *portp}
	logger.Write("INFO", "listening to "+*hostp+":"+strconv.Itoa(*portp))
	conn.Run()
}
