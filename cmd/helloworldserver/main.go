/*
  author='du'
  date='2020/1/17 20:21'
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>hello world %s", request.FormValue("name"))
	})
	http.ListenAndServe("8888", nil)
	fmt.Println("ok")
}
