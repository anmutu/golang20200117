/*
  author='du'
  date='2020/1/26 21:49'
*/

//并发版的engine

package engine

import (
	"golang20200117/crawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler    Scheduler        //用来调度
	WorkerCount  int              //worker的数量
	SaveItemChan chan interface{} //item的channel，用于存储数据
}

type Scheduler interface {
	Submit(Request)
	ConfigerMsaterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//目前所有的worker共用一个输入
	//in := make(chan Request)//这里的in是旧的createWorker用得到
	out := make(chan ParseResult)

	//e.Scheduler.ConfigerMsaterWorkerChan(in)
	e.Scheduler.Run()

	//建worker
	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out)
		newCreateWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//要收out
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("得到#%d ,item:%v\n", itemCount, item)
			itemCount++
			//在这里将拿到的items放到SaveItemChan里
			go func() { e.SaveItemChan <- item }()
		}

		//把所有的result送给scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

//创建worker，简单调度器
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func newCreateWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

//将fetch的动作提取成work成work函数。
func worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher失败，fetch的url是%s:,错误信息是:%v", r.Url, err)
		return ParseResult{}, err
	}
	parseResult := r.ParserFunc(body)
	return parseResult, nil
}
