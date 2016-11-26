package main

import (
	"connector"
	. "entry"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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
	conn := connector.Connector{ctx}

	logline := fmt.Sprintf("listening to %s:%s", ctx.GetHost(), strconv.Itoa(ctx.GetPort()))
	logger.Write("INFO", logline)
	fmt.Println(logline)

	go signalListen(ctx)

	conn.Run()
}

func signalListen(ctx *Context) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR2)
	for {
		fmt.Println("listen signal")
		s := <-c
		if s == syscall.SIGUSR2 {
			//重新加载配置文件
			LoadConfigFile(ctx)
		}
	}
}

func getCtx() *Context {
	ctx := new(Context)
	hostp := flag.String("host", HOST, "input host")
	portp := flag.Int("port", PORT, "input port")
	configfilep := flag.String("config", CONFIGFILE, "input config file")
	flag.Parse()

	host := *hostp
	port := *portp
	configfile := *configfilep

	ctx.SetHost(host)
	ctx.SetPort(port)
	ctx.SetConfigFile(configfile)
	ctx.SetLogFile(LOGFILE)

	//加载配置文件
	LoadConfigFile(ctx)
	return ctx
}

func LoadConfigFile(ctx *Context) {
	configfile := ctx.GetConfigFile()
	// 配置文件优先
	if configfile != "" {
		c, err := goconfig.LoadConfigFile(configfile)
		if err == nil {
			host, err1 := c.GetValue("", "host")
			port, err2 := c.Int("", "port")
			logfile, err3 := c.GetValue("", "logfile")
			if err1 == nil {
				ctx.SetHost(host)
			}
			if err2 == nil {
				ctx.SetPort(port)
			}
			if err3 == nil {
				ctx.SetLogFile(logfile)
			}
		}
	}
}
