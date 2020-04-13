package contiguousArray

func findMaxLength(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	ret, sum := 0, 0
	m := make(map[int]int)
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum]; ok {
			ret = max(ret, i-v)
		} else {
			m[sum] = i
		}
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
