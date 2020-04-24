package main

import (
	"fmt"
)

const columns = "abcdaf" // x
const rows = "acbcf"     // y

// dynamic programming example
// longest common subsequence
// given two strings
// how do we find the longest common subsequence?
// abcdaf
// acbcf
// => abcf

func main() {

	data := initializeTable()
	answer := process(data)
	fmt.Println(answer)
}

func initializeTable() (solutionTable [][]int) {
	// initialize rows first
	solutionTable = make([][]int, len(rows))

	// initialize colunns next
	for row := 0; row < len(rows); row++ {
		solutionTable[row] = make([]int, len(columns))
	}

	// do any other setup here
	// for index, _ := range solutionTable {
	//	solutionTable[index][0] = 1
	// }
	return solutionTable
}

func process(solutionTable [][]int) string {

	// assume a function builds these []bytes

	for rowIndex, row := range solutionTable {
		for columnIndex, _ := range row {
			solutionTable = longestCommonSubsequence(solutionTable, rowIndex, columnIndex)
		}
	}
	fmt.Println(solutionTable)
	return backTrace(solutionTable)
}

func longestCommonSubsequence(solutionTable [][]int, rowIndex, columnIndex int) [][]int {

	columnLetters := []byte{'a', 'b', 'c', 'd', 'a', 'f'}
	rowLetters := []byte{'a', 'c', 'b', 'c', 'f'}

	if rowLetters[rowIndex] == columnLetters[columnIndex] {
		// lcs this could be a function
		// skip if indexes are 0
		if rowIndex == 0 || columnIndex == 0 {
			solutionTable[rowIndex][columnIndex] = 1
			return solutionTable
		}
		solutionTable[rowIndex][columnIndex] = solutionTable[rowIndex-1][columnIndex-1] + 1

	} else {
		// max is slightly complex, it simply fills in the max(lookup, lookback) on the 2d table
		solutionTable[rowIndex][columnIndex] = max(solutionTable, rowIndex, columnIndex)
	}
	return solutionTable

}

// max has a lot of error conditions
// since we are "looking backwards"
func max(solutionTable [][]int, rowIndex, columnIndex int) int {

	// first catch both
	if columnIndex == 0 && rowIndex == 0 {
		return solutionTable[rowIndex][columnIndex]
	}

	// catch column index 0
	if columnIndex == 0 {
		return solutionTable[rowIndex-1][columnIndex]
	}

	//catch row index 0
	if rowIndex == 0 {
		return solutionTable[rowIndex][columnIndex-1]
	}

	if solutionTable[rowIndex-1][columnIndex] > solutionTable[rowIndex][columnIndex-1] {
		return solutionTable[rowIndex-1][columnIndex]
	}
	return solutionTable[rowIndex][columnIndex-1]
}

func backTrace(solutionTable [][]int) string {

	rowLetters := []byte{'a', 'c', 'b', 'c', 'f'}
	var answer []byte

	var tuple [][]int
	var lookBack, lookUp int
	for row := len(solutionTable) - 1; row >= 0; row-- {
		for col := len(solutionTable[row]) - 1; col >= 0; col-- {
			// look back guard
			//fmt.Println(i)
			lookBack, lookUp = 0, 0

			if col-1 >= 0 {
				lookBack = solutionTable[row][col-1]

			}

			// look up guard
			if row-1 >= 0 {
				lookUp = solutionTable[row-1][col]
			}

			// continue the column loop, if lookback is greater
			if lookBack > lookUp {
				continue
				// break the column loop, if lookup is greater
			} else if lookBack < lookUp {
				break
			}

			// this represents moving diagnally in a 2d chart
			// jumps up one row, and continue to look backwards from the same place
			// (diagnal move)
			if lookUp == lookBack {
				pair := []int{row, col}
				tuple = append([][]int{pair}, tuple...)
				answer = append([]byte{rowLetters[row]}, answer...)
				row--
				continue
			}
		}
	}
	// fmt.Printf("%v\n", tuple)
	return string(answer)
}
