/*
  author='du'
  date='2020/1/22 12:14'
*/
package main

func main() {
	//
	// failedChannel()
}

//fatal error: all goroutines are asleep - deadlock!
//最多只能2个。
func failedChannel() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	c <- 3
}
