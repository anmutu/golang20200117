/*
  author='du'
  date='2020/1/23 21:13'
*/
package parser

//
//import (
//	"golang20200117/crawler/engine"
//	"regexp"
//)

//func ParseCarList(
//	contents []byte, _ string) engine.ParseResult {
//	matches := carModelRe.FindAllSubmatch(contents, -1)
//
//	result := engine.ParseResult{}
//	for _, m := range matches {
//		result.Requests = append(
//			result.Requests, engine.Request{
//				Url: host + string(m[1]),
//				Parser: engine.NewFuncParser(
//					ParseCarModel, config.ParseCarModel),
//			})
//	}
//
//	matches = carListRe.FindAllSubmatch(contents, -1)
//	for _, m := range matches {
//		result.Requests = append(
//			result.Requests, engine.Request{
//				Url: "http:" + string(m[1]),
//				Parser: engine.NewFuncParser(
//					ParseCarList, config.ParseCarList),
//			})
//	}
//
//	return result
//}
