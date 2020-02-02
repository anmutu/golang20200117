/*
  author='du'
  date='2020/1/31 14:49'
*/
package xcar

import (
	"fmt"
	"golang20200117/crawler/engine"
	"regexp"
)

//http://newcar.xcar.com.cn/car/0-0-0-0-1-0-0-0-0-0-0-0/

const host = "http://newcar.xcar.com.cn"

var carListRe = regexp.MustCompile(`<a href="(//newcar.xcar.com.cn/car/[\d+-]+\d+/)"`)
var carBrandListRe = regexp.MustCompile(`<a href="/car/[^/]+/" [^>]*>[^>]+([^>]) </a>`)

//<a href="/car/0-0-0-0-1-0-0-0-0-0-0-0/" data-id="1"><span class="sign"><img src="//img1.xcarimg.com/PicLib/logo/pl1_160s.png-40x30.jpg"></span>奥迪(57) </a>
func ParseCarList(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}

	matches := carBrandListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        "http:" + string(m[1]),
				ParserFunc: engine.NilParser,
			})
		fmt.Printf("得到车品牌:%s", string(m[1]))
	}

	return result
}
