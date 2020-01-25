/*
  author='du'
  date='2020/1/23 20:49'
*/
package engine

//请求的结构体，包含一个url和解析这个url的一个的函数
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
