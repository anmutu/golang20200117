/*
  author='du'
  date='2020/1/23 16:12'
*/
package main

import (
	"fmt"
	"golang20200117/crawler/cnblogs/parser"
	"golang20200117/crawler/engine"
	"golang20200117/crawler/scheduler"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	concurrentBlogList()
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
