package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var arr []int
	if root == nil {
		return arr
	}
	arr = append(arr, root.Val)
	for _, node := range root.Children {
		arr = append(arr, preorder(node)...)
	}
	return arr
}

func main() {

}
