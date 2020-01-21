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
