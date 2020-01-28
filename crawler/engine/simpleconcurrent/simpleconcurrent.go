/*
  author='du'
  date='2020/1/28 16:37'
*/
package simpleconcurrent

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"golang20200117/crawler/engine"
)

//简单调度器,

type SimpleConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(engine.Request)
}

//这里是所有的worker共用一个输入
func (e SimpleConcurrentEngine) Run(seeds ...engine.Request) {
	in := make(chan engine.Request)
	out := make(chan engine.ParseResult)
	for i := 0; i < e.WorkCount; i++ {
		createWorker(in, out)
	}

	for {
		result := <-out

		//将拿到的值打印出来
		for _, item := range result.Items {
			fmt.Printf("拿到item的值是%s:", item)
		}

		//将拿到的seed给到scheduler去提交
		for _, item := range result.Requests {
			e.Scheduler.Submit(item)
		}
	}

}

func createWorker(in chan engine.Request, out chan engine.ParseResult) {
	go func() {
		for {
			request := <-in
			res, err := engine.Worker(request)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
