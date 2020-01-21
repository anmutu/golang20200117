/*
  author='du'
  date='2020/1/21 12:56'
*/
package main

import "fmt"

func main() {
	fmt.Println(maxNoRepeating("abcdefg"))
	fmt.Println(maxNoRepeating("aaaaa"))
	fmt.Println(maxNoRepeating("safa"))
}

func maxNoRepeating(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastOccured[ch] >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
