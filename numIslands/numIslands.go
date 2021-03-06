package numIslands

import (
	"bytes"
)

// This works perfectly fine locally, but refuses to work on the Leetcode runner.
// It seems that land is never assigned when running in Leetcode so none of logic is
// actually being checked, as there are no values to compare. I am not sure why this is.

func numIslands(grid [][]byte) int {

	var count int
	land := make(map[int][]byte)

	for i := range grid {
		for index, value := range grid[i] {
			if value == 1 {
				land[i] = append(land[i], byte(index))
			}
		}
	}

	if land[0] != nil {
		count++
	}

	for j := 1; j < len(land); j++ {
		for k := 0; k < len(land[j]); k++ {
			if bytes.ContainsRune(land[j-1], rune(land[j][k])) || bytes.ContainsRune(land[j], rune(land[j][k]-1)) {
				continue
			}
			count++
		}
	}

	return count
}
