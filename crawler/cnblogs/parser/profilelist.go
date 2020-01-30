/*
  author='du'
  date='2020/1/29 22:45'
*/
package parser

import (
	"golang20200117/crawler/engine"
	"regexp"
	"strconv"
)

//<a href="https://www.cnblogs.com/SSummerZzz/" class="lightblue">SSummerZzz</a>
const profileListRe = `<a href="(https://www.cnblogs.com/[^/]+/)" class="lightblue">([a-zA-Z0-9]+)</a>`
const preSeedUrl = "https://www.cnblogs.com/#p"

func ParseProfileList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(profileListRe)
	result := engine.ParseResult{}
	matches := re.FindAllSubmatch(contents, -1)
	index := 1
	for _, m := range matches {
		requestUrl := preSeedUrl + strconv.Itoa(index)
		result.Items = append(result.Items, string(m[1])+" "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: requestUrl, ParserFunc: engine.NilParser})
		index++
	}
	return result
}
