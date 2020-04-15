package productExceptSelf

func productExceptSelf(nums []int) []int {
	prod := make([]int, len(nums))

	prod[0] = 1
	for i := 1; i < len(nums); i++ {
		prod[i] = prod[i-1] * nums[i-1]
	}

	p := 1
	for i := len(nums) - 2; i >= 0; i-- {
		p *= nums[i+1]
		prod[i] *= p
	}

	return prod
}
