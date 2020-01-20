/*
  author='du'
  date='2020/1/20 10:51'
*/
package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(div(13, 4))
	fmt.Println(funcAsArg(add, 100, 200))

	a, b := 2, 3
	swap(&a, &b)
	fmt.Printf("交换后的值为%d和%d\n", a, b)

	c, d := 4, 5
	c, d = swap1(c, d)
	fmt.Printf("交换后的值为%d和%d\n", c, d)
}

//方法返回两个值，这里返回商和余数
func div(a, b int) (int, int) {
	return a / b, a % b
}

//函数是一等公民，把函数当作参数传递
func funcAsArg(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("调用函数名称为%s,参数为%d和%d\n", opName, a, b)
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}

//这个表示了如果不用指针交换是没有用的,当然这样看起来有点麻烦，其实可以返回两个值就可以了，也就是swap1()函数
func swap(a, b *int) {
	*a, *b = *b, *a
}

//这种才是更好的
func swap1(a, b int) (int, int) {
	return b, a
}
