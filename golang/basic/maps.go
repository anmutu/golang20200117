/*
  author='du'
  date='2020/1/21 11:10'
*/
package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "du",
		"job":  "engineer",
		"age":  "it is a secret",
	}

	//发现打印出来的值是无序的
	for k, v := range m {
		fmt.Println(k, v)
	}

	//得到值,判断是否存在key
	if wrongName, ok := m["wrongKey"]; ok {
		fmt.Println(wrongName)
	} else {
		fmt.Println("没有这个key值")
	}

	if rightName, ok := m["name"]; ok {
		fmt.Println(rightName)
	} else {
		fmt.Println("没有这个key值")
	}

	//删除
	name, ok := m["name"]
	fmt.Println("删除前的值是：", name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("删除后的值是：", name, ok) //没有了，就false了。

}
