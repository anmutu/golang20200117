/*
  author='du'
  date='2020/1/24 12:10'
*/
package main

import (
	"fmt"
	"regexp"
)

const text = "may i have ur email? my email is hbdw@du.com and i think u can have my wechat"

//"."是匹配任何一个字符
//"+"就是一个或者多个
//"*"就是0个或者多个
//"[a-zA-Z0-9]+"表示匹配的只能是一个或多个的大小的字母或者数字
//有了"`"则在表达式里面不用转义再转义了。
func main() {
	simpleMatch()
	simpleMatch1()
	emailMatch()
}

func simpleMatch() {
	re := regexp.MustCompile("hbdw@du.com")
	match := re.FindString(text)
	fmt.Println(match)
}

func simpleMatch1() {
	re := regexp.MustCompile(`.+@.+\..+`)
	match := re.FindString(text)
	fmt.Println(match)
}

func emailMatch() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	fmt.Println(match)
}
