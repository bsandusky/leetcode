package twoSum

func twoSum(nums []int, target int) []int {
	var sum int
	var out []int

	for i := range nums {
		if nums[i] > target {
			continue
		}

		sum += nums[i]

		if sum <= target {
			out = append(out, i)
		}
	}

	return out
}
