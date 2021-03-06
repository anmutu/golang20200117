/*
  author='du'
  date='2020/1/23 20:49'
*/
package engine

//请求的结构体，包含一个url和解析这个url的一个的函数
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
	//Parser Parser
}

//返回的结构体，其中"interface{}"表示任何表示任何类型，有点类似c#里的泛型T
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	PayLoad interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type ParserFunc func(contents []byte, url string) ParseResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
