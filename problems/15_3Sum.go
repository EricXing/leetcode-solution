package problems

import "sort"

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.
//
// Note:
//
// The solution set must not contain duplicate triplets.
//
// Example:
//
//	Given array nums = [-1, 0, 1, 2, -1, -4],
//
//	A solution set is:
//	[
// 	 [-1, 0, 1],
// 	 [-1, -1, 2]
//	]
func ThreeSum(nums []int) [][]int {
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
			t := -(nums[i] + nums[j])
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

type IntSliceSet struct {
	intSlice [][3]int
	sliceMap map[[3]int]bool
}

func NewIntSliceSet() *IntSliceSet {
	return &IntSliceSet{intSlice: [][3]int{}, sliceMap: map[[3]int]bool{}}
}

func (set *IntSliceSet) Add(arr [3]int) {
	_, ok := set.sliceMap[arr]
	if !ok {
		set.intSlice = append(set.intSlice, arr)
		set.sliceMap[arr] = true
	}
}

func (set *IntSliceSet) GetValues() [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(set.intSlice); i++ {
		result = append(result, set.intSlice[i][:])
	}
	return result
}

type SliceDict struct {
	value map[int]bool
	idx   map[int]int
}

func NewSliceDict() *SliceDict {
	return &SliceDict{value: make(map[int]bool), idx: make(map[int]int)}
}

func (dict *SliceDict) Put(value, idx int) {
	dict.value[value] = true
	dict.idx[value] += 1
}

func (dict *SliceDict) GetValue(value int) bool {
	_, ok := dict.value[value]
	return ok
}

func (dict *SliceDict) GetIdx(value int) int {
	if v, ok := dict.idx[value]; ok {
		return v
	} else {
		return 0
	}
}
