package lastStoneWeight

import (
	"sort"
)

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	sort.Ints(stones)
	var result int

	for i := len(stones) - 1; i > 0; i-- {
		result = stones[i] - stones[i-1]
		stones[i-1] = result
		sort.Ints(stones)
	}

	return stones[0]
}
