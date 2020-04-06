package defangIPaddr

import "testing"

var tts = []struct {
	in  string
	out string
}{
	{in: "1.1.1.1", out: "1[.]1[.]1[.]1"},
	{in: "255.100.50.0", out: "255[.]100[.]50[.]0"},
}

func TestDefangIPaddr(t *testing.T) {
	for _, tt := range tts {
		got := defangIPaddr(tt.in)
		if got != tt.out {
			t.Errorf("Expected: %s, Got: %s", tt.out, got)
		}
	}
}
