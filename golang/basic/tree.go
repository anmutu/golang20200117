/*
  author='du'
  date='2020/1/21 13:41'
*/
package main

import "fmt"

//只有指针才可以改变结构的内容
//nil可以调用函数

func main() {
	root := treeNode{Value: 10}
	root.Left = &treeNode{27, nil, nil}
	root.Right = &treeNode{20, nil, nil}
	root.Left.Left = &treeNode{}
	root.Left.Left.SetValue(712)
	root.Left.Right = &treeNode{}
	root.Left.Right.SetValue(713)
	//用工厂函数创建
	root.Right.Left = createNode(755)
	root.Right.Right = createNode(756)
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Left.Right)
	fmt.Println(root.Right)
	fmt.Println(root.Right.Left)
	fmt.Println(root.Right.Right)
}

type treeNode struct {
	Value       int
	Left, Right *treeNode
}

//只有指针才可以改变结构的内容
func (node *treeNode) SetValue(value int) {
	if node != nil {
		node.Value = value
	} else {
		fmt.Println("没有此node节点")
		return
	}
}

//工厂函数
func createNode(value int) *treeNode {
	return &treeNode{Value: value}
}
