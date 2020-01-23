/*
  author='du'
  date='2020/1/22 22:33'
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	printMazeFile()
}

//读取迷宫文件，返回一个二维的数组
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d,%d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

//第一步，读取出迷宫文件的内容，看是否是对的。
func printMazeFile() {
	maze := readMaze("golang/maze/maze.in")
	//fmt.Println(1)
	for _, row := range maze {
		//fmt.Println(2)
		for _, val := range row {
			//fmt.Println(3)
			fmt.Printf("%d", val)
		}
		fmt.Println()
	}
}
