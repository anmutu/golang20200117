/*
  author='du'
  date='2020/1/17 20:54'
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{21, 351, 2, 3, 5, 8, 13}
	sort.Ints(s)
	for i, v := range s {
		fmt.Println(i, v)
	}
}
