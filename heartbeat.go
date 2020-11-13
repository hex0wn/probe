package main

import (
	"net"
	"strings"
	"time"
	"github.com/sirupsen/logrus"
)

func getClientIp(prefix string) (string) {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && strings.HasPrefix(ipnet.IP.String(), prefix) {
				return ipnet.IP.String()
			}
		}
	}

	return ""

}

func beat(addr string) {
	accessLog.WithFields(logrus.Fields{
		"probe_ip": addr,
	}).Info("start")
	for range time.Tick(time.Minute * 5){
		accessLog.WithFields(logrus.Fields{
			"probe_ip": addr,
		}).Info("heartbeat")
	}
}