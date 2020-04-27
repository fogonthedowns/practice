package main

import "fmt"

// This is solved with a stack data structure
type Stack []int
type Histogram []int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str)
}

// Remove and return top element of stack.
// Returns false if empty
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Peek() int {
	index := len(*s) - 1
	return (*s)[index]
}

func newHistogram() Histogram {
	return Histogram{2, 1, 2}
}

func (h *Histogram) maxHistogram() {
	maxArea := 0
	area := 0
	var stack Stack
	stack = make(Stack, 1)
	var i int
	for i = 0; i < len(*h); i++ {
		if (stack.IsEmpty()) || stack.Peek() <= (*h)[i] {
			stack.Push((*h)[i])
		} else {
			top, empty := stack.Pop()
			if empty == false {
				panic("this shouldn't have happened")
			}
			if stack.IsEmpty() {
				area = top * i
			} else {
				area = top * (i - stack.Peek() - 1)
			}
			if area > maxArea {
				maxArea = area
			}
		}
	}

	for !stack.IsEmpty() {
		top, empty := stack.Pop()
		if empty == false {
			panic("this should not happen")
		}
		if stack.IsEmpty() {
			area = top * i
		} else {
			area = top * (i - stack.Peek() - 1)
		}
		if area > maxArea {
			maxArea = area
		}
	}
	fmt.Printf("max Area: %v\n", maxArea)
	fmt.Printf("stack: %v", stack)
}

func main() {
	h := newHistogram()
	h.maxHistogram()
}
