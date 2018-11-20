package problems

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
// You are given a target value to search. If found in the array return its index, otherwise return -1.
//
// You may assume no duplicate exists in the array.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// Example 1:
//
//	Input: nums = [4,5,6,7,0,1,2], target = 0
//	Output: 4
//
// Example 2:
//
//	Input: nums = [4,5,6,7,0,1,2], target = 3
//	Output: -1
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	mid := 0
	right := len(nums) - 1
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[right] {
			if target < nums[mid] && target >= nums[left] {
				right -= 1
			} else {
				left += 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return -1
}
