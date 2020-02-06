/*
  author='du'
  date='2020/2/4 22:51'
*/
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "129.211.78.6:6379")
	if err != nil {
		fmt.Println("连接redis失败，失败原因是:", err)
	}
	defer conn.Close()

	item, err := redis.String(conn.Do("lpop", "hbdw:fruit:apple:n500:1"))
	if err != nil {
		fmt.Println("从redis中pop数据失败")
		return
	}
	fmt.Println("成功从redis中pop出数据，数值是:", item)
}
