/*
  author='du'
  date='2020/4/25 15:07'
*/
package main

import (
	"golang20200117/golang/rpcdemo"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Add", rpcdemo.Args{10, 20}, &result)
	if err != nil {
		panic(err)
	} else {
		log.Printf("结果是：%v", result)
	}

}
