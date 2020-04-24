package main

import (
	"fmt"
)

type NodeDg struct {
	PointsTo   []*NodeDg
	Id         int
	Identifier string
}

type Graph []NodeDg

var visited Graph
var sorted Graph

func main() {
	g := BuildGraph()
	for x := 0; x < len(g); x++ {
		g.topological(x)
		fmt.Printf("len %v\n", sorted)
		if len(g) == len(sorted) {

			break
		}
	}
	fmt.Printf("sorted\n: %v\n", sorted)
	fmt.Printf("**visit\n: %v\n", visited)
}

func BuildGraph() Graph {
	var graph Graph
	g := NodeDg{PointsTo: []*NodeDg{nil}, Id: 0, Identifier: "G"}
	f := NodeDg{PointsTo: []*NodeDg{&g}, Id: 1, Identifier: "F"}
	e := NodeDg{PointsTo: []*NodeDg{&f}, Id: 2, Identifier: "E"}
	c := NodeDg{PointsTo: []*NodeDg{&e}, Id: 3, Identifier: "C"}
	d := NodeDg{PointsTo: []*NodeDg{&f}, Id: 4, Identifier: "D"}
	b := NodeDg{PointsTo: []*NodeDg{&c, &d}, Id: 4, Identifier: "B"}
	a := NodeDg{PointsTo: []*NodeDg{&c}, Id: 5, Identifier: "A"}
	graph = append(graph, g)
	graph = append(graph, f)
	graph = append(graph, e)
	graph = append(graph, c)
	graph = append(graph, d)
	graph = append(graph, b)
	graph = append(graph, a)
	return graph
}

func (g Graph) topological(index int) string {

	fmt.Printf("numern %v\n", index)
	g.visitChildren(g[index])
	return fmt.Sprintf("%v", visited)
}

func (g Graph) visitChildren(n NodeDg) (c []*NodeDg) {
	if visited.inList(n.Id) {
	} else {
		visited = append(visited, n)
	}
	for _, value := range n.PointsTo {
		if value == nil {
			sorted = append(sorted, n)
			return nil
		}
		if visited.inList(n.Id) {
			continue
		}
		g.visitChildren(*value)
	}
	return n.PointsTo
}

func (g Graph) inList(id int) bool {
	for _, node := range g {
		if node.Id == id {
			return true
		}
	}
	return false
}
