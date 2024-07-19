package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"randomRocket"
)

//var remoteAddr = "127.0.0.1:49808"

func main() {
	passwd := flag.String("p", "12345", "")
	remoteAddr := flag.String("h", "", "")
	flag.Parse()
	rc, err := randomRocket.NewLsLocal(*passwd, ":7448", *remoteAddr)
	if err != nil {
		return
	}
	log.Fatalln(rc.Listen(func(listenAddr *net.TCPAddr) {

		log.Println(fmt.Sprintf(`
lightsocks-local:%s 启动成功，配置如下：
本地监听地址：
%s
远程服务地址：
%s
密码：
%s`, "V1", listenAddr, remoteAddr, *passwd))
	}))

}
