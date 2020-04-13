/*
  author='du'
  date='2020/1/19 23:24'
*/
package main

import (
	"fmt"
)

func main() {
	enums()
	//0 2 3 4 5
	//1 1024 1048576 1073741824 1099511627776
}

func enums() {
	const (
		csharp = iota
		_
		java
		python
		golang
		javascript
	)
	fmt.Println(csharp, java, python, golang, javascript)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)
}
