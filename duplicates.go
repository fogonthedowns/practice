package main

import (
	"fmt"
)

type List []int

var visited List

func main() {
	var newList List
	list := List{100, 2, 2, 4, 13, 100, 2, 9}
	for _, v := range list {
		newList = removeDuplicates(v)
	}
	fmt.Printf("%v\n", newList)
}

func removeDuplicates(v int) List {
	for _, value := range visited {
		if v == value {
			return visited
		}
	}
	visited = append(visited, v)
	return visited
}
