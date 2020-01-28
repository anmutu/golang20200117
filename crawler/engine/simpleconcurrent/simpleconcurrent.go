/*
  author='du'
  date='2020/1/28 16:37'
*/
package simpleconcurrent

import (
	"fmt"
	"golang20200117/crawler/engine"
)

//简单调度器,

type SimpleConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(engine.Request)
	ConfigerMasterWorkerChan(chan engine.Request)
}

//这里是所有的worker共用一个输入
func (e *SimpleConcurrentEngine) Run(seeds ...engine.Request) {
	in := make(chan engine.Request)
	out := make(chan engine.ParseResult)

	//把in给到workerChan
	e.Scheduler.ConfigerMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		//把request送进workerChan
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		//将拿到的值打印出来
		for _, item := range result.Items {
			fmt.Printf("拿到item的值是%s:\n", item)
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
