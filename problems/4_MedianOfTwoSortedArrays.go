package problems

//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
//Example 1:
//	nums1 = [1, 3]
//	nums2 = [2]
//
//	The median is 2.0
//
// Example 2:
//	nums1 = [1, 2]
//	nums2 = [3, 4]
//
//	The median is (2 + 3)/2 = 2.5
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var arr = make([]int, len(nums1) + len(nums2))
	var i = 0
	var j = 0
	//merge two array to ordered array
	for k := 0; k < len(arr); k++ {
		switch {
		case i == len(nums1):
			arr[k] = nums2[j]
			j++
		case j == len(nums2):
			arr[k] = nums1[i]
			i++
		case nums1[i] < nums2[j]:
			arr[k] = nums1[i]
			i++
		case nums2[j] <= nums1[i]:
			arr[k] = nums2[j]
			j++
		}
	}
	//find the median
	if len(arr) % 2 == 0 {
		idx := len(arr) / 2
		return float64(arr[idx - 1] + arr[idx]) / 2.0
	} else {
		idx := (len(arr) - 1) / 2
		return float64(arr[idx])
	}
}
