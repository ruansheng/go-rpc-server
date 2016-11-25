package main

import (
	"connector"
	. "entry"
	"flag"
	"fmt"
	"strconv"
	"utils/goconfig"
	"utils/logger"
)

const (
	HOST       = "127.0.0.1"
	PORT       = 8080
	CONFIGFILE = ""
	LOGFILE    = "./go-rpc-server.log"
)

func main() {
	ctx := getCtx()
	conn := connector.Connector{ctx.GetHost(), ctx.GetPort()}

	logline := fmt.Sprintf("listening to %s:%s", ctx.GetHost(), strconv.Itoa(ctx.GetPort()))
	logger.Write("INFO", logline)
	fmt.Println(logline)
	conn.Run()
}

func getCtx() Ctx {
	ctx := new(Ctx)
	hostp := flag.String("host", HOST, "input host")
	portp := flag.Int("port", PORT, "input port")
	configfilep := flag.String("config", CONFIGFILE, "input config file")
	flag.Parse()

	host := *hostp
	port := *portp
	configfile := *configfilep
	logfile := LOGFILE

	// 配置文件优先
	if configfile != "" {
		c, err := goconfig.LoadConfigFile(configfile)
		if err == nil {
			hostc, err1 := c.GetValue("", "host")
			portc, err2 := c.Int("", "port")
			logfilec, err3 := c.GetValue("", "logfile")
			if err1 == nil {
				host = hostc
			}
			if err2 == nil {
				port = portc
			}
			if err3 == nil {
				logfile = logfilec
			}
		}
	}
	ctx.SetHost(host)
	ctx.SetPort(port)
	ctx.SetConfigFile(configfile)
	ctx.SetLogFile(logfile)
	return ctx
}
