/*
  author='du'
  date='2020/1/24 11:16'
*/
package engine

import (
	"golang20200117/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//如果requests里有就一直请求，如果其中有错，注意不要panic,要纪录日志。
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher失败，fetch的url是%s:,错误信息是:%v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...) //有了这行，可能会有后面的解析报错

		for _, item := range parseResult.Items {
			log.Printf("得到itme:%s", item)
		}
	}
}
