package firstUnique

type FirstUnique struct {
	Keys   []int
	Counts map[int]int
}

func Constructor(nums []int) FirstUnique {
	m := mapValues(nums)
	return FirstUnique{Counts: m, Keys: nums}
}

func (this *FirstUnique) ShowFirstUnique() int {
	if len(this.Counts) == 0 {
		return -1
	}

	for _, k := range this.Keys {
		if this.Counts[k] == 1 {
			return k
		}
	}
	return -1
}

func (this *FirstUnique) Add(value int) {
	this.Keys = append(this.Keys, value)
	this.Counts[value]++
}

func mapValues(nums []int) map[int]int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	return m
}
