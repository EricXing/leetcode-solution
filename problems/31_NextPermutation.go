package problems

// Implement next permutation, which rearranges numbers into the
// lexicographically next greater permutation of numbers.
//
// If such arrangement is not possible, it must rearrange
// it as the lowest possible order (ie, sorted in ascending order).
//
// The replacement must be in-place and use only constant extra memory.
//
// Here are some examples. Inputs are in the left-hand column and
// its corresponding outputs are in the right-hand column.
//
//	1,2,3 → 1,3,2
//	3,2,1 → 1,2,3
//	1,1,5 → 1,5,1
func NextPermutation(nums []int) []int {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i -= 1
	}
	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j -= 1
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for m, n := i+1, len(nums)-1; m < n; m, n = m+1, n-1 {
		nums[m], nums[n] = nums[n], nums[m]
	}
	return nums
}
