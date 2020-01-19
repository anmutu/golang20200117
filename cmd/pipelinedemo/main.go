/*
  author='du'
  date='2020/1/17 21:28'
*/
package main

import (
	"bufio"
	"fmt"
	"golang20200117/pipeline"
	"os"
)

func main() {
	const filename = "large.in"
	const n = 100000000

	//创建文件
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//随机生成n个数并且写入
	p := pipeline.RandomSource(n)
	pipeline.WriteSink(bufio.NewWriter(file), p) //用到bufio明显会快很多。

	//打开文件
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//读取文件
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}

}

func MergeDemo() {
	//p1:=pipeline.InMemSort(pipeline.ArraySource(35,21,13,8,5))
	//p2:=pipeline.InMemSort(pipeline.ArraySource(3,5,1,9,7,11))
	//p:=pipeline.Merge(p1,p2)

	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(35, 21, 13, 8, 5)),
		pipeline.InMemSort(pipeline.ArraySource(3, 5, 1, 9, 7, 11)))

	//第一种写法
	//for{
	//	if num,ok:=<-p;ok{
	//		fmt.Println(num)
	//	}else{
	//		break
	//	}
	//}

	//第二种写法
	for v := range p {
		fmt.Println(v)
	}
}
