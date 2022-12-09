package main

import (
	"fmt"
	"strconv"
)

func main() {
	data, _ := ReadInput()

	matrix := [][]int{}

	for rowIndex, rowString := range data {
		matrix = append(matrix, []int{})
		for _, columnRune := range rowString {
			column, _ := strconv.Atoi(string(columnRune))
			matrix[rowIndex] = append(matrix[rowIndex], column)
		}
	}

	visibleTrees := 0
	highestScore := 0

	for x := range matrix {
		for y := range matrix[x] {
			horizontal, scoreH := checkHorizontal(matrix, x, y)
			vertical, scoreV := checkVertical(matrix, y, x)
			if horizontal || vertical {
				visibleTrees += 1
				if scoreH*scoreV > highestScore {
					highestScore = scoreH * scoreV
				}
			}
		}
	}

	fmt.Println(visibleTrees, highestScore)

}

func checkVertical(matrix [][]int, x int, intersect int) (bool, int) {
	trees := []int{}
	for _, row := range matrix {
		trees = append(trees, row[x])
	}

	visible, up, down := checkArray(trees, intersect)
	return visible, up * down
}

func checkHorizontal(matrix [][]int, y int, intersect int) (bool, int) {
	visible, left, right := checkArray(matrix[y], intersect)
	return visible, left * right
}

func checkArray(trees []int, position int) (bool, int, int) {
	revTrees := make([]int, position+1)
	copy(revTrees, trees)

	ReverseInts(revTrees)

	visibleBelow, bScore := isVisible(trees, revTrees[len(revTrees)-position:], position)
	visibleAbove, aScore := isVisible(trees, trees[position+1:], position)
	score := map[string]int{"a": aScore, "b": bScore}

	if visibleBelow || visibleAbove {
		return true, score["a"], score["b"]
	}

	return false, score["a"], score["b"]
}

func isVisible(trees []int, slicedTrees []int, position int) (bool, int) {
	visible := true
	score := 0
	for _, tree := range slicedTrees {
		if trees[position] <= tree {
			visible = false
			score += 1
			break
		} else if trees[position] > tree {
			score += 1
		}
	}

	if position == 0 {
		score = 0
	}

	return visible, score
}
