package happyNumber

import (
	"math"
	"strconv"
	"strings"
)

// A happy number is a number defined by the following process:
// Starting with any positive integer,
// replace the number by the sum of the squares of its digits,
// and repeat the process until the number equals 1 (where it will stay),
// or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy numbers.

// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1

func isHappy(n int) bool {

	seen := make(map[int]bool)
	return checkHappy(n, seen)
}

func checkHappy(n int, seen map[int]bool) bool {

	// check if already seen to close recursive loop
	if seen[n] {
		return false
	}

	// if not add to map
	seen[n] = true

	digits := getDigits(n)

	// must be positive integer
	if n < 0 {
		return false
	}

	// math
	var out float64
	for _, d := range digits {
		out += math.Pow(d, 2)
	}

	// check
	if int(out) == 1 {
		return true
	}

	// recursion
	return checkHappy(int(out), seen)
}

func getDigits(n int) []float64 {

	// convert int to string
	str := strconv.Itoa(n)

	// split string into individual digits
	strs := strings.Split(str, "")

	// convert []string into []float64
	var digits []float64
	for _, s := range strs {
		num, _ := strconv.ParseFloat(s, 64)
		digits = append(digits, num)
	}
	return digits
}
