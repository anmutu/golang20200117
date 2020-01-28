/*
  author='du'
  date='2020/1/28 17:20'
*/
package engine

import (
	"golang20200117/crawler/fetcher"
	"log"
)

//将fetch的动作提取成work成work函数。
func Worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher失败，fetch的url是%s:,错误信息是:%v", r.Url, err)
		return ParseResult{}, err
	}
	parseResult := r.ParserFunc(body)
	return parseResult, nil
}
