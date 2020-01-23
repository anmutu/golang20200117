/*
  author='du'
  date='2020/1/23 16:12'
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.cnblogs.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("错误，状态号", resp.StatusCode)
		return
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s \n",contents)

	printBlogList(contents)

}

func printBlogList(contents []byte) {
	//<h3><a class="titlelnk" href="https://www.cnblogs.com/wotxdx/p/12230486.html" target="_blank">漫画 | 什么是散列表（哈希表）？</a></h3>

	//https://careers.tencent.com/search.html?pcid=40001
	re := regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	mathes := re.FindAll(contents, -1)
	fmt.Println(mathes)
	for _, m := range mathes {
		fmt.Printf("%s\n", m)
		//fmt.Println(11)
	}
	fmt.Println(len(mathes))
}
