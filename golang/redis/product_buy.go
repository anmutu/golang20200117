/*
  author='du'
  date='2020/2/5 17:41'
*/
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	REDISKEY = "hbdw:fruit:apple:n500:1"
)

func main() {
	conn, err := redis.Dial("tcp", "129.211.78.6:6379")
	if err != nil {
		fmt.Println("连接redis失败，失败的error是：", err)
	}
	defer conn.Close()

	_, err = conn.Do("lpush", REDISKEY, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if err != nil {
		fmt.Println("向redis里push数据失败，失败原因是：", err)
		return
	}
	fmt.Println("成功向redis里push了数据")
}
