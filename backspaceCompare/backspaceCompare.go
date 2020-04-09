package backspaceCompare

func backspaceCompare(S string, T string) bool {
	s := trimBackspace(S)
	t := trimBackspace(T)
	return s == t
}

func trimBackspace(str string) string {
	tmp := []byte{}
	for i := 0; i < len(str); i++ {
		if str[i] != '#' {
			tmp = append(tmp, str[i])
			continue
		}
		if len(tmp) > 0 {
			tmp = tmp[:len(tmp)-1]
		}
	}
	return string(tmp)
}
