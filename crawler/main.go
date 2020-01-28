/*
  author='du'
  date='2020/1/23 16:12'
*/
package main

import (
	"golang20200117/crawler/cnblogs/parser"
	"golang20200117/crawler/engine"
	"golang20200117/crawler/engine/simpleconcurrent"
	"golang20200117/crawler/scheduler"
	"golang20200117/crawler/scheduler/simplescheduler"
)

func main() {
	simpleConcurrent()
	//singleTask()
	//concurrentBlogList()
}

func concurrentBlogList() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "https://www.cnblogs.com",
		ParserFunc: parser.ParseBlogList,
	})
}

//单任务版本爬虫
func singleTask() {
	e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:        "https://www.cnblogs.com",
		ParserFunc: parser.ParseBlogList,
	})
}

//简单调度器版本爬虫
func simpleConcurrent() {
	e := simpleconcurrent.SimpleConcurrentEngine{
		Scheduler:   &simplescheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "https://www.cnblogs.com",
		ParserFunc: parser.ParseBlogList,
	})
}
