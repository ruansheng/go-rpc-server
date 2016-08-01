package logger

import (
	"log"
	"os"
)

func Write(mode string, line string) bool {
	var filename = "go-rpc-server.log"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	logger := log.New(f, mode+":", log.LstdFlags|log.Lshortfile)
	logger.Println(line)

	return true
}
