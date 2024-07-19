package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"randomRocket"
	"strconv"

	"github.com/phayes/freeport"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// 优先从环境变量中获取监听端口
	port, err := strconv.Atoi(os.Getenv("LIGHTSOCKS_SERVER_PORT"))
	// 服务端监听端口随机生成
	if err != nil {
		port, err = freeport.GetFreePort()
	}
	if err != nil {
		// 随机端口失败就采用 7448
		port = 7448
	}
	// 默认配置
	passwd := randomRocket.RandPassword()
	//passwd := "123123"
	// 启动 server 端并监听
	lsServer, err := randomRocket.NewLsServer(passwd, fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsServer.Listen(func(listenAddr *net.TCPAddr) {
		log.Println(fmt.Sprintf(`
lightsocks-server:%s 启动成功，配置如下：
服务监听地址：
%s
密码：
%s`, "V1", listenAddr, passwd))
	}))
}
