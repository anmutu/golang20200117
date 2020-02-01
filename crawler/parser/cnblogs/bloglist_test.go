/*
  author='du'
  date='2020/1/25 17:05'
*/
package cnblogs

import (
	"golang20200117/crawler/parser"
	"io/ioutil"
	"testing"
)

func TestParseBlogList(t *testing.T) {
	//contents,err:=fetcher.Fetch("https://www.cnblogs.com")
	contents, err := ioutil.ReadFile("bloglist_test_data")
	if err != nil {
		panic(err)
	}
	res := parser.ParseBlogList(contents)
	//fmt.Printf("%s",contents)
	const resultSize = 15
	if len(res.Requests) != resultSize {
		t.Errorf("结果的值是%d,期待的值是:%d", len(res.Requests), resultSize)
	}
}
