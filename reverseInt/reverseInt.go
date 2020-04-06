package reverseInt

import (
	"log"
	"math"
	"strconv"
)

func reverse(x int) int {
	// parse input into string to review values
	str := strconv.Itoa(x)

	// new string to be converted to int
	// handle negative sign if needed
	var rev string
	if str[0] == '-' {
		rev = "-"
	} else {
		rev = ""
	}

	// reverse int string into new string
	for x := len(str); x > 0; x-- {
		if str[x-1] == '-' {
			continue
		}
		rev += string(str[x-1])
	}

	// convert output into int
	out, err := strconv.Atoi(rev)
	if err != nil {
		log.Fatal(err)
	}

	if out > math.MaxInt32 || out < math.MinInt32 {
		return 0
	}
	return out
}
