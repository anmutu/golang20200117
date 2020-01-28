/*
  author='du'
  date='2020/1/28 16:37'
*/
package simpleconcurrent

import (
	"fmt"
	"golang20200117/crawler/engine"
)

//简单调度器

type SimpleConcurrentEngine struct {
	Scheduler   Scheduler //调度器
	WorkerCount int       //worker的数量
}

type Scheduler interface {
	Submit(engine.Request)                        //用于将Request送给worker
	ConfigerMasterWorkerChan(chan engine.Request) //配置workChan,实现也就是个赋值操作。
}

//这里是所有的worker共用一个输入，但是会有多个去处理送过来的输入。
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

//创建worker，传入Request的channel和
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
