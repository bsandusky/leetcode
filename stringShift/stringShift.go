package stringShift

func stringShift(s string, shift [][]int) string {

	var move int
	for _, v := range shift {
		if v[0] == 0 {
			move -= v[1]
		} else {
			move += v[1]
		}
	}

	move = move % len(s)

	if move == 0 {
		return s
	} else if move > 0 {
		return s[len(s)-move:] + s[:len(s)-move]
	} else {
		return s[-move:] + s[:-move]
	}

}
