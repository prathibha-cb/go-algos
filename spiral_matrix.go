package main

import (
	"log"
)

func main() {
	traversal := make([]int, 0)
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	rows, cols := 0, 0
	rows = len(matrix)
	if rows >= 1 {
		cols = len(matrix[0])
	}

	verticalCount := rows
	horizontalCount := cols

	i, j := 0, 0
	first := true
	// Direction: 0 L->R, 1 T->B, 2: R->L, 3: B->T
	direction := 0

	for {
		if direction == 0 { // Left to Right
			if !first {
				horizontalCount--
			} else {
				first = false
			}
			if horizontalCount == 0 {
				break
			}

			for count := 1; count <= horizontalCount; count++ {
				traversal = append(traversal, matrix[i][j])
				j++
			}
			j--

		} else if direction == 1 { // Top to Botton
			verticalCount--
			if verticalCount == 0 {
				break
			}

			for count := 1; count <= verticalCount; count++ {
				traversal = append(traversal, matrix[i][j])
				i++
			}
			i--

		} else if direction == 2 { // Right to Left
			horizontalCount--
			if horizontalCount == 0 {
				break
			}

			for count := 1; count <= horizontalCount; count++ {
				traversal = append(traversal, matrix[i][j])
				j--
			}
			j++

		} else if direction == 3 { // Bottom to Top
			verticalCount--
			if verticalCount == 0 {
				break
			}

			for count := 1; count <= verticalCount; count++ {
				traversal = append(traversal, matrix[i][j])
				i--
			}
			i++
		}
		// Direction change
		direction, i, j = makeTurn(direction, i, j)
	}

	log.Printf("Traversal: %v", traversal)
}

func makeTurn(direction, i, j int) (int, int, int) {

	direction++
	if direction == 4 {
		direction = 0
	}

	switch direction {
	case 0:
		return direction, i, j + 1
	case 1:
		return direction, i + 1, j
	case 2:
		return direction, i, j - 1
	case 3:
		return direction, i - 1, j
	}
	panic("Unknown direction")
}
