/*
  author='du'
  date='2020/1/19 23:36'
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	simpleif()
}

func simpleif() {
	const filename = "du.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
}
