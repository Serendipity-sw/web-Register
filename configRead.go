package main

import (
	"github.com/guotie/config"
	"strings"
)

//配置文件读取
func configRead() {
	serverListeningPort = config.GetIntDefault("serverListeningPort", 8080)
	logsDir = config.GetString("logsDir")
	forwardingDomain = strSplit(config.GetString("forwardingDomain"))
	rootPrefix = strSplit(strings.TrimSpace(config.GetStringDefault("rootPrefix", "")))
}

func strSplit(str string) string {
	if len(str) != 0 {
		if strings.HasSuffix(str, "/") {
			str = str[0 : len(str)-1]
		}
	}
	return str
}
