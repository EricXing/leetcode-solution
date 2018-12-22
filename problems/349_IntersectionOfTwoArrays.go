package problems

// Given two arrays, write a function to compute their intersection.
//
// Example 1:
//
//	Input: nums1 = [1,2,2,1], nums2 = [2,2]
//	Output: [2]
//
// Example 2:
//
//	Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//	Output: [9,4]
//
// Note:
// Each element in the result must be unique.
// The result can be in any order.
func intersection(nums1 []int, nums2 []int) []int {
	m1, m2 := make(map[int]bool), make(map[int]bool)
	for _, val := range nums1 {
		m1[val] = true
	}
	for _, val := range nums2 {
		m2[val] = true
	}
	var short, long map[int]bool
	if len(m1) > len(m2) {
		short = m2
		long = m1
	} else {
		short = m1
		long = m2
	}
	result := make([]int, 0)
	for key := range short {
		if _, exist := long[key]; exist {
			result = append(result, key)
		}
	}
	return result
}
