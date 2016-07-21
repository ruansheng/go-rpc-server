package utils

import (
	"strconv"
	"strings"
)

func ParseProtocal(cmd string) []string {
	lines := strings.Split(cmd, "\r\n")
	return lines
}

func MakeGetProtocal(resp string) string {
	llen := len(resp)
	lines := make([]string, 5, 5)
	lines = append(lines, "$")
	lines = append(lines, strconv.Itoa(llen))
	lines = append(lines, "\r\n")
	lines = append(lines, resp)
	lines = append(lines, "\r\n")
	return strings.Join(lines, "")
}
