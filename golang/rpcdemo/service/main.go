/*
  author='du'
  date='2020/4/25 14:56'
*/
package main

import (
	"golang20200117/golang/rpcdemo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener accept出错：%v", err)
			continue
		}
		//成功的情况
		jsonrpc.ServeConn(conn)
	}
}
