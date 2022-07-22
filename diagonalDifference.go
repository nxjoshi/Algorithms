package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	var lsum, rsum int32
	size := len(arr)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				rsum = rsum + arr[i][j]
			}
			if i+j == size-1 {
				lsum = lsum + arr[i][j]
			}
		}
	}
	diagdiff := math.Abs(float64(lsum - rsum))
	return int32(diagdiff)

}

func main() {
	type Matrix [][]int32
	oneMatrix := Matrix{{11, 2, 4}, {4, 10, 6}, {10, 8, 12}}
	number := diagonalDifference(oneMatrix)
	fmt.Print(number)
}
