package main

import (
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "192.168.1.103:7448")
	defer conn.Close()
	if err != nil {
		return
	}
	for true {
		conn.Write([]byte("Hello!"))
		time.Sleep(1 * time.Second)
	}

}
