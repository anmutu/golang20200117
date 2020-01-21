/*
  author='du'
  date='2020/1/20 14:20'
*/
package main

import "fmt"

func main() {
	getSlicesValue()
	slicesIsViewOfArray()
	extendingSlices()
	append2Slices()
	createSlices()
	copyDelPop()
}

//切片取值
func getSlicesValue() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])
}

//证明下slices里array的一个view,也就是修改slices里面的值也就是修改array里面的值了。
//slices本身是没有数据的，是对底层array的一个view
func slicesIsViewOfArray() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:]
	fmt.Println("update前的s1:")
	fmt.Println(s1)
	fmt.Println("update过后的s1:")
	updateSlices(s1)
	fmt.Println(s1)
	fmt.Println("现在arr的值是:")
	fmt.Println(arr)
}

//进入extendingSlices()函数
//arr最初的值是: [0 1 2 3 4 5 6 7 8]
//s1的值，即arr[2:6]的值是: [2 3 4 5]
//s2的值，即s1[3:5]的值是: [5 6]
//为什么s2的值是会有，因为slices是可以扩展的。
// 一个slices里面有三个变量：
// 1.ptr。指向第一个，
// 2.len。表示这个slices的长度，
// 3.cap。指从ptr到array的最后一个的长度，只要不超过这个cap的长度，就都可以扩展。
//slice是可以向后扩展的，但是不能向前扩展。
func extendingSlices() {
	fmt.Println("进入extendingSlices()函数")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("arr最初的值是:%v,len(arr)=%d,cap(arr)=%d \n", arr, len(arr), cap(arr))
	fmt.Printf("s1的值，即arr[2:6]的值是:%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2的值，即s1[3:5]的值是:%v,len(s1)=%d,cap(s1)=%d\n", s2, len(s2), cap(s2))
}

//进入append2Slices()函数
//s1,s2,s3,s4= [6 7 8] [6 7 8 9] [6 7 8 9 10] [6 7 8 9 10 11]
//arr= [0 1 2 3 4 5 6 7 8]
//超过cap会分配更大的底层数组，已经不是同一个了。
func append2Slices() {
	fmt.Println("进入append2Slices()函数")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[6:]
	s2 := append(s1, 9)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	fmt.Println("s1,s2,s3,s4=", s1, s2, s3, s4)
	fmt.Println("arr=", arr)
}

func createSlices() {
	fmt.Println("进入createSlices()函数")
	var s []int
	for i := 0; i < 50; i++ {
		printSlices(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
}

func copyDelPop() {
	fmt.Println("进入copyDelPop()函数")
	s1 := make([]int, 10, 32)
	s2 := []int{0, 1, 2, 3}
	printSlices(s1)
	printSlices(s2)
	fmt.Println("将s2copy到s1的结果为：")
	copy(s1, s2)
	printSlices(s1)
	//把现在下标为3的数值delete掉
	fmt.Println("将下标为3的delete掉的结果为：")
	s2 = append(s1[:3], s1[4:]...)
	printSlices(s1)
	fmt.Println("从头pop一个数据出去：")
	front := s1[0]
	s1 = s1[1:]
	fmt.Printf("pop出去的头的数值为:%v,pop后的结果为:%v", front, s1)
	printSlices(s1)
	fmt.Println("从尾部pop一个数据出去：")
	tail := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	fmt.Printf("pop出去的尾部的数值为:%v,pop后的结果为:%v", tail, s1)
	printSlices(s1)
}

func printSlices(s []int) {
	fmt.Printf("%v,len(s)=%d,cap(s)=%d \n", s, len(s), cap(s))
}

func updateSlices(s []int) {
	s[0] = 1000
}
