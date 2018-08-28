package problems

import (
	"math"
	"strings"
)

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
//	Symbol       Value
//	I             1
//	V             5
//	X             10
//	L             50
//	C             100
//	D             500
//	M             1000
//
// For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII,
// which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for
// four is not IIII. Instead, the number four is written as IV. Because the one is before the five we
// subtract it making four. The same principle applies to the number nine, which is written as IX.
// There are six instances where subtraction is used:
// 	 I can be placed before V (5) and X (10) to make 4 and 9.
// 	 X can be placed before L (50) and C (100) to make 40 and 90.
// 	 C can be placed before D (500) and M (1000) to make 400 and 900.
// Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
//
// Example 1:
//
//	Input: 3
//	Output: "III"
//
// Example 2:
//
//	Input: 4
//	Output: "IV"
//
// Example 3:
//
//	Input: 9
//	Output: "IX"
//
// Example 4:
//
//	Input: 58
//	Output: "LVIII"
//	Explanation: C = 100, L = 50, XXX = 30 and III = 3.
//
// Example 5:
//
// Input: 1994
//	Output: "MCMXCIV"
//	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
func IntToRoman(num int) string {
	result := ""
	dict10 := map[int]string{
		3: "M",
		2: "C",
		1: "X",
		0: "I",
	}
	dict5 := map[int]string{
		2: "D",
		1: "L",
		0: "V",
	}
	for ; num > 0; {
		// how many digits in the number
		n := int(math.Trunc(math.Log10(float64(num))))
		p := int(math.Pow10(n))
		// first digit of the number
		digit := num / p
		switch {
		case digit < 4:
			result = result + strings.Repeat(dict10[n], digit)
		case digit == 4:
			result = result + dict10[n] + dict5[n]
		case digit >= 5 && digit < 9:
			result = result + dict5[n] + strings.Repeat(dict10[n], digit-5)
		case digit == 9:
			result = result + dict10[n] + dict10[n+1]
		}
		num -= digit * p
	}
	return result
}
