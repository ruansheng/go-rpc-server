package handler

import (
	"conf"
	"encoding/json"
	"entry"
	"fmt"
	"net"
	"reflect"
	"strings"
	"utils"
	"utils/logger"
)

type Handler struct {
	Ctx *entry.Context
}

func (h *Handler) HandleConnection(conn *net.Conn) {
	fmt.Println("config file:", h.Ctx.GetLogFile())
	defer (*conn).Close()
	defer func() {
		if r := recover(); r != nil {
			logger.Write("ERROR", r.(string))
		}
	}()
	buf := make([]byte, 1024)
	for {
		n, err := (*conn).Read(buf)
		if err == nil {
			lines := utils.ParseProtocal(string(buf[:n]))
			if len(lines) == 4 && strings.EqualFold(lines[2], "PING") { // ping command
				pong := "+PONG\r\n"
				(*conn).Write([]byte(pong))
			} else if len(lines) == 6 { // get command
				logger.Write("INFO", lines[4])
				var request entry.Request
				cmd_bytes := []byte(lines[4])
				json_err := json.Unmarshal(cmd_bytes, &request)
				if json_err == nil {
					ret := getReflectInvokeRet(request)
					//logger.Write("INFO", ret.(string))
					resp := entry.Response{"1", 200, "success", ret}
					ret_cmd, _ := json.Marshal(resp)
					redis_cmd := utils.MakeGetProtocal(string(ret_cmd))
					(*conn).Write([]byte(redis_cmd))
				} else {
					panic(json_err)
				}
			} else if len(lines) == 2 && strings.EqualFold(lines[0], "QUIT") { // QUIT recommand
				panic("QUIT command")
			} else {
				panic("command is error:" + string(buf[:n]))
			}
		} else {
			break
		}
	}
}

func getReflectInvokeRet(request entry.Request) interface{} {
	// get conf
	serviceUris := conf.GetConf()
	confItem, ok := serviceUris[request.Action]
	if ok {
		implClass := confItem.InterClass
		mtv := reflect.ValueOf(implClass)
		in := getArgs(request.Params.Args)
		ret := mtv.MethodByName(request.Params.M).Call(in)
		return ret[0].Interface()
	} else {
		return nil
	}
}

func getArgs(args []interface{}) []reflect.Value {
	fmt.Println(args)
	in := make([]reflect.Value, len(args))
	for k, param := range args {
		in[k] = reflect.ValueOf(param)
	}
	return in
}
