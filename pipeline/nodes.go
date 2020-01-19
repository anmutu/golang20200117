/*
  author='du'
  date='2020/1/17 21:16'
*/
package pipeline

import (
	"encoding/binary"
	"io"
	"sort"
)

//其中"<-"表示只能出东西，表示输出。
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	//读进内存
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		//排序
		sort.Ints(a)
		//输出
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

//归并节点
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out) //我关了，你别再发了
	}()
	return out
}

//读数据
func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil {
				break
			}
		}
		close(out)
	}()
	return out

}
