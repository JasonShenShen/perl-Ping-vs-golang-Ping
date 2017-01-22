package main

import (
	"fmt"
	"github.com/ping-master"
	"strconv"
	"time"
)

var fin = make(chan string)
var limits = make(chan int, 100)

func doit(ip string) {
	limits <- 1
	defer func() { <-limits }()
	pingr(ip, 5)
}

func pingr(host string, timeout int) {
	t1 := time.Now().UnixNano()
	alive, err := ping.Ping(host, timeout)
	//alive,err:=true,0
	t2 := time.Now().UnixNano()
	fmt.Println(host, alive, (t2-t1)/1000000, err)
	fin <- host
}

func main() {

	//定义iplist
	var iplist [253]string

	//数组赋值1..254
	for i := 0; i < len(iplist); i++ {
		iplist[i] = "192.168.6." + strconv.Itoa(i+1)
		go doit(iplist[i])
	}

	//等待结束
	for i := 0; i < len(iplist); i++ {
		<-fin
		//fmt.Println(ip,"joined")
	}

}
