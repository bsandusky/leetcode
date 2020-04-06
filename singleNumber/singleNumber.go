package singleNumber

func singleNumber(nums []int) int {

	var out int

	for i := 0; i < len(nums); i++ {
		count := tally(nums, nums[i])
		if count > 1 {
			continue
		}
		out = nums[i]
	}

	return out
}

func tally(slice []int, val int) int {
	var count int
	for _, item := range slice {
		if item == val {
			count++
		}
	}
	return count
}
