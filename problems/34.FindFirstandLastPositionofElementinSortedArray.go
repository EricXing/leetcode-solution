package problems

// Given an array of integers nums sorted in ascending order,
// find the starting and ending position of a given target value.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// If the target is not found in the array, return [-1, -1].
//
// Example 1:
//
//	Input: nums = [5,7,7,8,8,10], target = 8
//	Output: [3,4]
//
// Example 2:
//
//	Input: nums = [5,7,7,8,8,10], target = 6
//	Output: [-1,-1]
func searchRange(nums []int, target int) []int {
	mid := binarySearch(nums, target)
	if mid == -1 {
		return []int{-1, -1}
	}
	left := findIndex(nums, target, 0, mid, true)
	if left == -1 {
		left = mid
	}
	right := findIndex(nums, target, mid, len(nums)-1, false)
	if right == -1 {
		right = mid
	}
	return []int{left, right}
}

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func findIndex(nums []int, target, left, right int, isLeft bool) int {
	mid := left + (right-left)/2
	if isLeft {
		switch {
		case nums[mid] == target && (mid == left || mid > left && nums[mid-1] != target):
			return mid
		case (nums[mid] == target && mid > 0 && nums[mid-1] == target) || nums[mid] > target:
			right = mid - 1
			//left index is not in the left part and the right part is not empty
		case nums[right] < target && target < nums[len(nums)-1]:
			left = right + 1
			right = len(nums) - 1
		case nums[mid] < target:
			left = mid + 1
		case left == mid && nums[mid] != target:
			return -1
		}
	} else {
		switch {
		//mid is the last target
		case nums[mid] == target && (mid == right || mid < right && nums[mid+1] != target):
			return mid
			//mid is not the last target
		case (nums[mid] == target && mid < len(nums)-1 && nums[mid+1] == target) || nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		case target < nums[left] && target > nums[0]:
			right = left - 1
			left = 0
		case mid == right && nums[mid] != target:
			return -1
		}
	}
	return findIndex(nums, target, left, right, isLeft)
}
