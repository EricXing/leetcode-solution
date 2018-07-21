package problems

import "strconv"

// Given two binary strings, return their sum (also a binary string).
//
// The input strings are both non-empty and contains only characters 1 or 0.
//
// Example 1:
//
//	Input: a = "11", b = "1"
//	Output: "100"
//
// Example 2:
//
//	Input: a = "1010", b = "1011"
//	Output: "10101"
func AddBinary(a string, b string) string {
	length := len(a)
	if len(b) > len(a) {
		length = len(b)
	}
	var result = make([]int, length)
	var proto = 0
	i := len(a) - 1
	j := len(b) - 1
	for idx := length - 1; idx >= 0; idx-- {
		var tmp = 0
		m := 0
		n := 0
		if i >= 0 {
			m, _ = strconv.Atoi(string(a[i]))
		}
		if j >= 0 {
			n, _ = strconv.Atoi(string(b[j]))
		}
		tmp = m + n + proto
		if tmp >= 2 {
			proto = 1
			result[idx] = tmp % 2
		} else {
			proto = 0
			result[idx] = tmp
		}
		i--
		j--
	}
	var s = ""
	if proto > 0 {
		s = "1"
	}
	for _, val := range result {
		s += string(strconv.Itoa(val))
	}
	return s
}
