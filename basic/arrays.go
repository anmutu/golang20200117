/*
  author='du'
  date='2020/1/20 13:57'
*/
package main

import "fmt"

func main() {
	var0 := [5]int{}
	//不定长的数据
	var1 := [...]int{2, 4, 6, 8, 10}
	//4行5列的数据
	var2 := [4][5]int{}

	fmt.Println(var1)
	fmt.Println(var2)

	//可以用range去遍历
	for i, v := range var1 {
		fmt.Println(i, v)
	}

	//证明下数组是值类型
	fmt.Println("开始打印var0:")
	printArray(var0)
	fmt.Println("开始打印var1:")
	printArray(var1)
	fmt.Println("开始打印var0和var1:")
	fmt.Println(var0, var1)
}

func printArray(arr [5]int) {
	arr[0] = 1000
	for _, v := range arr {
		fmt.Println(v)
	}
}
