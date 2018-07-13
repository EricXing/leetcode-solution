package problems

//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//	Input: 123
//	Output: 321
//
// Example 2:
//
//	Input: -123
//	Output: -321
//
// Example 3:
//
//	Input: 120
//	Output: 21
//
//Note:
//Assume we are dealing with an environment which could only
// store integers within the 32-bit signed integer range: [−231,  231 − 1].
// For the purpose of this problem, assume that your function returns 0 when
// the reversed integer overflows.
func Reverse(x int) int {
	var intMax = 2147483647
	var intMin = -2147483648
	var reversed = 0
	for ; x != 0;  {
		var pop = x % 10
		x = x / 10
		if reversed >intMax / 10 || (reversed == intMax / 10 && pop > 7){
			return 0
		}
		if reversed < intMin / 10 || (reversed == intMin / 10 && pop < -8){
			return 0
		}
		reversed = reversed * 10 + pop
	}

	return reversed
}
