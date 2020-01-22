/*
  author='du'
  date='2020/1/22 21:02'
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//使用select来调度数据

func main() {
	var c1, c2 = generator(), generator()
	for {
		select {
		case n := <-c1:
			fmt.Println("从c1的channel：", n)
		case n := <-c2:
			fmt.Println("从c2的channel:", n)
		}
	}
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
