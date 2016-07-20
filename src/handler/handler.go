package handler

import (
	"conf"
	"encoding/json"
	"entry"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
)

func HandleConnection(conn *net.Conn) {
	fmt.Println(conn)
	defer (*conn).Close()
	buf := make([]byte, 1024)
	var cmd string
	for {
		n, err := (*conn).Read(buf)
		if err == nil {
			cmd = string(buf[:n])
			var lines []string
			fmt.Println(cmd)
			lines = strings.Split(cmd, "\r\n")
			fmt.Println(lines)
			fmt.Println(len(lines))
			if len(lines) == 6 { // get command
				fmt.Println(lines[4])
				var request entry.Request
				cmd_bytes := []byte(lines[4])
				jerr := json.Unmarshal(cmd_bytes, &request)
				if jerr == nil {
					fmt.Println(request.Id, request.Action, request.Params)
					serviceUris := conf.GetConf()
					fmt.Println(serviceUris, request.Action)
					if confItem, ok := serviceUris[request.Action]; ok {
						implClass := confItem.InterClass
						fmt.Println(confItem, implClass)
						mtv := reflect.ValueOf(implClass)
						ret := mtv.MethodByName(request.Params.M).Call(nil)
						resp := entry.Response{"1", 200, "success", ret}
						ret_cmd, _ := json.Marshal(resp)
						ret_cmd_str := string(ret_cmd)
						redis_cmd := joinRedisCmd(ret_cmd_str)
						fmt.Println(redis_cmd)
						(*conn).Write([]byte(redis_cmd))
					} else {
						fmt.Println(request.Action + "not exists !")
					}
				} else {
					fmt.Println(jerr)
				}
			} else if len(lines) == 2 {

			}
		} else {
			break
		}
	}
}

func joinRedisCmd(resp string) string {
	slen := len(resp)
	lines := make([]string, 5, 5)
	lines = append(lines, "$")
	lines = append(lines, strconv.Itoa(slen))
	lines = append(lines, "\r\n")
	lines = append(lines, resp)
	lines = append(lines, "\r\n")
	return strings.Join(lines, "")
}
