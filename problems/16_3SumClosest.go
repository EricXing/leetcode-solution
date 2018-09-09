package problems

import "math"

// Given an array nums of n integers and an integer target, find three integers in nums such
// that the sum is closest to target. Return the sum of the three integers.
// You may assume that each input would have exactly one solution.
//
// Example:
//
//	Given array nums = [-1, 2, 1, -4], and target = 1.
//
//	The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
func ThreeSumClosest(nums []int, target int) int {
	var result int
	startFlag := true
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if startFlag {
					result = sum
					startFlag = false
				} else {
					if math.Abs(float64(target-result)) > math.Abs(float64(target-sum)) {
						result = sum
					}
				}
			}
		}
	}
	return result
}
