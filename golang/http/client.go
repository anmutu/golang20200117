/*
  author='du'
  date='2020/1/22 23:20'
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	simpleClient()
}

func simpleClient() {
	resp, err := http.Get("https://www.cnblogs.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true) //把body也dupm下来
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
