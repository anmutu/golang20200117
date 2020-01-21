/*
  author='du'
  date='2020/1/21 16:17'
*/
package main

import (
	"golang20200117/golang/retriever/mock"
	"golang20200117/golang/retriever/real"
)

func main() {
	var r Retriever
	r = mock.Retriever{"这里是模拟的retriever"}
	r = real.Retriever{}
	println(download(r))

}

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.cnblogs.com")
}
