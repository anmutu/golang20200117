/*
  author='du'
  date='2020/2/4 22:26'
*/
package main

import (
	"fmt"
)

func Sum(var1, var2 int, quit chan bool) {
	s := var1 + var2
	fmt.Println(s)
	quit <- true
}

func main() {
	channels := make([]chan bool, 12)
	for i := 0; i < 12; i++ {
		channels[i] = make(chan bool)
		go Sum(i, i, channels[i])
	}

	//如果不用time.Sleep也要显示出来，那么就需要把channel里的值送出去。
	//time.Sleep(1*time.Second)

	for _, v := range channels {
		<-v
	}
}
