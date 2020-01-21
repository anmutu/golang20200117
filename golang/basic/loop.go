/*
  author='du'
  date='2020/1/19 23:50'
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(convert2Binary(2))
	fmt.Println(convert2Binary(3))
	fmt.Println(convert2Binary(4))
	fmt.Println(convert2Binary(5))
	fmt.Println(convert2Binary(6))
	printFile("du.txt")
}

//十进制转换成二进制（用模二除法，然后将数倒过来）
func convert2Binary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//这里的for相当于其他语言的while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
