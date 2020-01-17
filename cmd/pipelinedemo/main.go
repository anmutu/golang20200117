/*
  author='du'
  date='2020/1/17 21:28'
*/
package main

import (
	"fmt"
	"golang20200117/pipeline"
)

func main() {

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
