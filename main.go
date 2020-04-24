package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type NodeInit map[int][]int

func main() {
	ds := make(NodeInit, 0)
	ds[0] = []int{8, -1, -1}
	ds[1] = []int{3, -1, -1}
	ds[2] = []int{5, 0, 1}
	ds[3] = []int{2, 4, -1}
	ds[5] = []int{10, 3, 2}
	ds[4] = []int{6, -1, -1}

	fmt.Println(ds)
	nodes := createNodes(ds)
	last := nodes[len(nodes)-1]

	fmt.Println(height(&last))
	fmt.Println("**************** Depth First:")
	depthFirstBinaryTree(&last)
}

func createNodes(n NodeInit) (nodes []Node) {
	for i := 0; i < len(n); i++ {
		node := Node{}
		nodes = append(nodes, node)
	}

	for idx, _ := range n {
		nodes[idx].Value = n[idx][0]
		if n[idx][1] == -1 {
			nodes[idx].Left = nil
		} else {
			nodes[idx].Left = &nodes[n[idx][1]]
		}
		if n[idx][2] == -1 {
			nodes[idx].Right = nil
		} else {
			nodes[idx].Right = &nodes[n[idx][2]]
		}
	}
	return nodes

}

func height(n *Node) int {
	if n == nil {
		return 0
	}

	leftHeight := height(n.Left)
	rightHeight := height(n.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func depthFirstBinaryTree(n *Node) {
	if n != nil {
		fmt.Println(n.Value)
		depthFirstBinaryTree(n.Left)
		depthFirstBinaryTree(n.Right)
	}
}
