/*
  author='du'
  date='2020/1/31 14:49'
*/
package xcar

import (
	"golang20200117/crawler/engine"
	"regexp"
)

//http://newcar.xcar.com.cn/car/0-0-0-0-1-0-0-0-0-0-0-0/

const host = "http://newcar.xcar.com.cn"

var carListRe = regexp.MustCompile(`<a href="(//newcar.xcar.com.cn/car/[\d+-]+\d+/)"`)

func ParseCarList(
	contents []byte, _ string) engine.ParseResult {
	//matches := carModelRe.FindAllSubmatch(contents, -1)
	//
	//result := engine.ParseResult{}
	//for _, m := range matches {
	//	result.Requests = append(
	//		result.Requests, engine.Request{
	//			Url: host + string(m[1]),
	//			Parser: engine.NewFuncParser(
	//				ParseCarModel, config.ParseCarModel),
	//		})
	//}

	matches := carListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: "http:" + string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCarList, config.ParseCarList),
			})
	}

	return result
}
