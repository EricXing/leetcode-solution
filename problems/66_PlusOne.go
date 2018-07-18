package problems

// Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
//
// The digits are stored such that the most significant digit is at the head of the list,
// and each element in the array contain a single digit.
//
// You may assume the integer does not contain any leading zero, except the number 0 itself.
//
// Example 1:
//
//	Input: [1,2,3]
//	Output: [1,2,4]
//	Explanation: The array represents the integer 123.
//
// Example 2:
//
//	Input: [4,3,2,1]
//	Output: [4,3,2,2]
//	Explanation: The array represents the integer 4321.
func PlusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	var proto = 0
	for i := len(digits) -1; i >= 0; i-- {
		digits[i] += proto
		if digits[i] >= 10 {
			digits[i] -= 10
			proto = 1
		} else {
			proto = 0
		}
	}
	if proto > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}