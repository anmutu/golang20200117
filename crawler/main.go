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
	//blogList()
}

//func blogList(){
//	engine.SimpleEngine.Run(engine.Request{
//		Url:        "https://www.cnblogs.com",
//		ParserFunc: parser.ParseBlogList,
//	})
//}

//func concurrentBlogList() {
//	e := engine.ConcurrentEngine{
//		Scheduler:   &scheduler.SimpleScheduler{},
//		WorkerCount: 10,
//	}
//
//	e.Run(engine.Request{
//		Url:        "https://www.cnblogs.com",
//		ParserFunc: parser.ParseBlogList,
//	})
//}

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

func concurrentBlogList1() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url:        "https://www.cnblogs.com",
		ParserFunc: parser.ParseNextBlogs,
	})
}

//<a href="/sitehome/p/2" class="p_2 current" onclick="aggSite.loadCategoryPostList(2,20);buildPaging(2);return false;">2</a>

//func excellentUser(){
//	engine.Run(engine.Request{
//		Url:"https://www.cnblogs.com",
//		ParserFunc:parser.ParseExcellentUser,
//	})
//}

func cnblogsTest() {
	resp, err := http.Get("https://www.cnblogs.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("错误，状态号", resp.StatusCode)
		return
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	printBlogList(contents)
}

func printBlogList(contents []byte) {
	//<h3><a class="titlelnk" href="https://www.cnblogs.com/ITnoteforlsy/p/12228149.html" target="_blank">B-Tree 和 B+Tree 结构及应用，InnoDB 引擎， MyISAM 引擎</a></h3>
	//用"[a-zA-Z0-9]","[0-9]","[^<]"匹配。
	re := regexp.MustCompile(`<h3><a class="titlelnk" href="https://www.cnblogs.com/[a-zA-Z0-9]+/p/[0-9]+.html" target="_blank">[^<]+</a></h3>`)

	mathes := re.FindAll(contents, -1)
	fmt.Println(mathes)
	for _, m := range mathes {
		fmt.Printf("%s\n", m)
	}
	fmt.Println(len(mathes))
}
