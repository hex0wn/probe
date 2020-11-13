package main

import (
	"fmt"
	"sync"
)

const (
	VERSION = "2.0"
)

func main() {
	fmt.Printf("switcher %s\n", VERSION)
	wg := &sync.WaitGroup{}
	for _, v := range config.Rules {
		wg.Add(1)
		go listen(v, wg)
	}
	addr := getClientIp(config.IPrefix)
	go beat(addr)  //heartbeat
	wg.Wait()
	fmt.Printf("program exited\n")
}
