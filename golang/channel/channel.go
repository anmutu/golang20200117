/*
  author='du'
  date='2020/1/21 22:11'
*/
package main

import (
	"fmt"
	"time"
)

//channel是goroutine与goroutine的交互，必须用另外一个goroutine去接收。
//channel也可以做为参数传递
//"chan<-"表示只能发数据。

func main() {
	//一个失败的例子
	//failedChannelExample()
	//successChanExample()
	//channelAsArg()
	//channelAsArg1()
	//manyChannel()
	manyChannel1() //这里的channel是没有表明方向的
	manyChannel2() //这里调用的时候channel是有方向的
}

//fatal error: all goroutines are asleep - deadlock!
//因为channel是goroutine与goroutine的交互，必须用另外一个goroutine去接收。所以这里错了。
func failedChannelExample() {
	c := make(chan int)
	c <- 1
	c <- 2
	n := <-c
	fmt.Println(n)
}

//成功的channel的案例。用一个goroutine一直去接收。
func successChanExample() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

//channel也可以做为参数传递
func channelAsArg() {
	c := make(chan int)
	go worker(c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func workerWithId(id int, c chan int) {
	for {
		fmt.Printf("worker的id是%d,接收值是%v \n", id, <-c)
	}
}

func channelAsArg1() {
	c := make(chan int)
	go workerWithId(0, c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func workersWithIds(id int, c chan int) {
	for {
		fmt.Printf("worker的id是%d,接收值是%c \n", id, <-c)
	}
}

func manyChannel() {
	//创建10个channel
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go workersWithIds(i, channels[i])
	}

	//向channel里面一直给数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //注意这里给到的值会是a,b,c,d,e,f,g,h,i,j这10个。
	}

	//再向channel里面给数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i //注意这里给到的值会是A,B,C,D,E,F,G,H,I,J这10个。
	}
	time.Sleep(time.Millisecond * 2)

}

func createWorker(id int) chan int {
	c := make(chan int)
	go workersWithIds(id, c)
	return c
}

func manyChannel1() {
	//创建10个channel
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i) //make(chan int)
	}

	//向channel里面一直给数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //注意这里给到的值会是a,b,c,d,e,f,g,h,i,j这10个。
	}

	//再向channel里面给数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i //注意这里给到的值会是A,B,C,D,E,F,G,H,I,J这10个。
	}
	time.Sleep(time.Millisecond * 2)

}

//特别注意：这里的"chan<-"表示只能发数据。这里和上面createWorker的函数只有返回值这里多了表示方向的"<-"
func createWorkerWithSendChannel(id int) chan<- int {
	c := make(chan int)
	go workersWithIds(id, c)
	return c
}

func manyChannel2() {
	//创建10个channel
	var channels [10]chan<- int //这里用的时候也要表明方向
	for i := 0; i < 10; i++ {
		channels[i] = createWorkerWithSendChannel(i) //make(chan int)
	}

	//向channel里面一直给数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //注意这里给到的值会是a,b,c,d,e,f,g,h,i,j这10个。
	}

	//再向channel里面给数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i //注意这里给到的值会是A,B,C,D,E,F,G,H,I,J这10个。
	}
	time.Sleep(time.Millisecond * 2)

}
