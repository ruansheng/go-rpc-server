package handler

import (
	"conf"
	"encoding/json"
	"entry"
	"fmt"
	"net"
	"reflect"
	//	"strconv"
	"strings"
	"utils"
)

func HandleConnection(conn *net.Conn) {
	defer (*conn).Close()
	buf := make([]byte, 1024)
	for {
		n, err := (*conn).Read(buf)
		if err == nil {
			lines := utils.ParseProtocal(string(buf[:n]))
			if len(lines) == 4 && strings.EqualFold(lines[2], "PING") { // ping command
				pong := "+PONG\r\n"
				(*conn).Write([]byte(pong))
			} else if len(lines) == 6 { // get command
				fmt.Println(lines[4])
				var request entry.Request
				cmd_bytes := []byte(lines[4])
				json_err := json.Unmarshal(cmd_bytes, &request)
				if json_err == nil {
					ret := getReflectInvokeRet(request)
					resp := entry.Response{"1", 200, "success", ret}
					ret_cmd, _ := json.Marshal(resp)
					redis_cmd := utils.MakeGetProtocal(string(ret_cmd))
					(*conn).Write([]byte(redis_cmd))
				} else {
					fmt.Println(json_err)
				}
			} else if len(lines) == 2 && strings.EqualFold(lines[1], "QUIT") { // QUIT recommand

			} else {
				fmt.Println("command is error:", string(buf[:n]))
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
		return ret
	} else {
		return nil
	}
}

func getArgs(args []interface{}) []reflect.Value {
	in := make([]reflect.Value, len(args))
	for k, param := range args {
		type_str := reflect.TypeOf(param).String()
		fmt.Println(type_str)
		if strings.EqualFold(type_str, "string") {
			in[k] = reflect.ValueOf(param)
		} else if strings.EqualFold(type_str, "float64") {
			pv := reflect.ValueOf(param).Float()
			in[k] = reflect.ValueOf(int(pv))
		} else {
			in[k] = reflect.ValueOf(param)
		}
	}
	return in
}
