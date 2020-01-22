/*
  author='du'
  date='2020/1/22 15:03'
*/
package main

import (
	"fmt"
	"sync"
)

//想去掉原来的暂停的时间的代码。
//如何通知你，我发送完了。

func main() {
	manyChannel()
}

type worker struct {
	in   chan int
	done func()
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker的id是%d,接收值是%c \n", id, n)
		w.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func manyChannel() {
	var wg sync.WaitGroup

	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		//wg.Add(1)
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//wg.Add(1)
	}

	wg.Wait()

}
