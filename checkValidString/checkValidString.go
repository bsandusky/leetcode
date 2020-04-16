package checkValidString

func checkValidString(s string) bool {

	var left, right []int

	for i, char := range s {
		if char == 40 {
			left = append(left, i)
		} else if char == 42 {
			right = append(right, i)
		} else {
			if len(left) > 0 {
				left = left[:len(left)-1]
			} else if len(right) > 0 {
				right = right[:len(right)-1]
			} else {
				return false
			}
		}
	}

	for len(left) > 0 && len(right) > 0 {
		if left[len(left)-1] > right[len(right)-1] {
			return false
		}
		left = left[:len(left)-1]
		right = right[:len(right)-1]
	}
	return len(left) == 0
}
