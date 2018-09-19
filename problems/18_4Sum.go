package problems

import "sort"

// Given an array nums of n integers and an integer target, are there elements a, b, c, and d
// in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
//
// Note:
//
// The solution set must not contain duplicate quadruplets.
//
// Example:
//
//	Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
//
//	A solution set is:
//	[
// 	 [-1,  0, 0, 1],
// 	 [-2, -1, 1, 2],
// 	 [-2,  0, 0, 2]
//	]
func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	tempSlice := make([]int, len(nums))
	for idx, num := range nums {
		tempSlice[idx] = target - num
	}
	result := NewIntSliceSetWith4Ele()
	for i := 0; i < len(tempSlice); i++ {
		subSlice := make([]int, 0, len(nums)-1)
		subSlice = append(subSlice, nums[:i]...)
		subSlice = append(subSlice, nums[i+1:]...)
		s := threeSum(subSlice, tempSlice[i])
		for _, val := range s {
			resultSlice := append(val, nums[i])
			sort.Ints(resultSlice)
			result.Add([4]int{resultSlice[0], resultSlice[1], resultSlice[2], resultSlice[3]})
		}
	}
	return result.GetValues()
}

type IntSliceSetWith4Ele struct {
	intSlice [][4]int
	sliceMap map[[4]int]bool
}

func NewIntSliceSetWith4Ele() *IntSliceSetWith4Ele {
	return &IntSliceSetWith4Ele{intSlice: [][4]int{}, sliceMap: map[[4]int]bool{}}
}

func (set *IntSliceSetWith4Ele) Add(arr [4]int) {
	_, ok := set.sliceMap[arr]
	if !ok {
		set.intSlice = append(set.intSlice, arr)
		set.sliceMap[arr] = true
	}
}

func (set *IntSliceSetWith4Ele) GetValues() [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(set.intSlice); i++ {
		result = append(result, set.intSlice[i][:])
	}
	return result
}

func threeSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	var dict = NewSliceDict()
	for idx, value := range nums {
		dict.Put(value, idx)
	}
	result := NewIntSliceSet()
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			t := target - (nums[i] + nums[j])
			if ok := dict.GetValue(t); ok {
				indexCount := dict.GetIdx(t)
				if nums[i] == t {
					indexCount -= 1
				}
				if nums[j] == t {
					indexCount -= 1
				}
				if indexCount > 0 {
					r := sort.IntSlice{nums[i], nums[j], t}
					sort.Stable(r)
					result.Add([3]int{r[0], r[1], r[2]})
				}
			}
		}
	}
	return result.GetValues()
}
