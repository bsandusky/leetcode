package countingElements

func countElements(arr []int) int {

	var tally int

	for _, val := range arr {
		for _, check := range arr {
			if val+1 == check {
				tally++
				break
			}
		}
	}
	return tally
}
