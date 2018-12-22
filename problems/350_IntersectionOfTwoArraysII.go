package problems

// Given two arrays, write a function to compute their intersection.
//
// Example 1:
//
//	Input: nums1 = [1,2,2,1], nums2 = [2,2]
//	Output: [2,2]
//
// Example 2:
//
//	Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//	Output: [4,9]
//
// Note:
// Each element in the result should appear as many times as it shows in both arrays.
// The result can be in any order.
//
// Follow up:
// What if the given array is already sorted? How would you optimize your algorithm?
// What if nums1's size is small compared to nums2's size? Which algorithm is better?
// What if elements of nums2 are stored on disk, and the memory is limited such that
// you cannot load all elements into the memory at once?
func intersect(nums1 []int, nums2 []int) []int {
	m1, m2 := make(map[int]int), make(map[int]int)
	for _, val := range nums1 {
		if num, ok := m1[val]; ok {
			m1[val] = num + 1
		} else {
			m1[val] = 1
		}
	}
	for _, val := range nums2 {
		if num, ok := m2[val]; ok {
			m2[val] = num + 1
		} else {
			m2[val] = 1
		}
	}
	var short, long map[int]int
	if len(m1) > len(m2) {
		short = m2
		long = m1
	} else {
		short = m1
		long = m2
	}
	result := make([]int, 0)
	for key := range short {
		if longCount, exist := long[key]; exist {
			var count int
			if longCount < short[key] {
				count = longCount
			} else {
				count = short[key]
			}
			for i := 0; i < count; i++ {
				result = append(result, key)
			}
		}
	}
	return result
}
