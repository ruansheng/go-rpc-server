package main

import (
	"connector"
	"flag"
	"fmt"
	"strconv"
	//	"strings"
	"utils/goconfig"
	"utils/logger"
)

const (
	host   = "127.0.0.1"
	port   = 8080
	config = ""
)

func main() {
	hostp := flag.String("host", host, "input host")
	portp := flag.Int("port", port, "input port")
	configp := flag.String("config", config, "input config")
	flag.Parse()

	// 配置文件优先
	if *configp != "" {
		chost, cport := getHostAndPort(*configp)
		fmt.Println(chost, cport)
		if chost != "" && cport != 0 {
			*hostp = chost
			*portp = cport
		}
	}

	conn := connector.Connector{*hostp, *portp}
	logger.Write("INFO", "listening to "+*hostp+":"+strconv.Itoa(*portp))
	conn.Run()
}

func getHostAndPort(file string) (string, int) {
	c, err := goconfig.LoadConfigFile(file)
	if err == nil {
		host, err1 := c.GetValue("", "host")
		port, err2 := c.Int("", "port")
		if err1 == nil && err2 == nil {
			return host, port
		}
	}
	return "", 0
}
