package groupAnagrams

import (
	"sort"
)

type bytes []byte

func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {

	var output [][]string
	groups := make(map[string][]string)

	for _, str := range strs {
		key := bytes(str)
		sort.Sort(key)
		tmp := string(key)
		groups[tmp] = append(groups[tmp], str)
	}

	for _, v := range groups {
		output = append(output, v)
	}

	return output
}
