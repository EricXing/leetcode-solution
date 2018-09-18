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
	sumSlice := make([]int, 0)
	sliceMap := NewSliceSumMap()
	sortedSlice := sort.IntSlice(nums)
	sort.Stable(sortedSlice)
	for i := 0; i < len(sortedSlice); i++ {
		for j := i + 1; j < len(sortedSlice); j++ {
			sum := sortedSlice[i] + sortedSlice[j]
			sumSlice = append(sumSlice, sum)
			sliceMap.Put(sum, []int{sortedSlice[i], sortedSlice[j]})
		}
	}

	result := NewIntSliceSetWith4Ele()
	for m := 0; m < len(sumSlice); m++ {
		for n := m + 1; n < len(sumSlice); n++ {
			sum := sumSlice[m] + sumSlice[n]
			if sum == target {
				combineIntSlice([][][]int{sliceMap.getSlices(sumSlice[m]), sliceMap.getSlices(sumSlice[n])}, make([][]int, 2), 0, result)
			}
		}
	}
	return result.GetValues()
}

// Cartesian Product of int slices.
func combineIntSlice(input [][][]int, current [][]int, k int, result *IntSliceSetWith4Ele) {
	if k == len(input) {
		tempSlice := sort.IntSlice{current[0][0], current[0][1], current[1][0], current[1][1]}
		sort.Stable(tempSlice)
		result.Add([4]int{tempSlice[0], tempSlice[1], tempSlice[2], tempSlice[3]})
	} else {
		for j := 0; j < len(input[k]); j++ {
			current[k] = input[k][j]
			combineIntSlice(input, current, k+1, result)
		}
	}
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

// SliceSumMap store sum of two numbers as key, and slices of the two numbers as value.
// There could be more than one combination.
type SliceSumMap struct {
	value  map[int]bool
	slices map[int][][]int
}

func NewSliceSumMap() *SliceSumMap {
	return &SliceSumMap{value: make(map[int]bool), slices: make(map[int][][]int, 0)}
}

func (sliceMap *SliceSumMap) Put(value int, nums []int) {
	sliceMap.value[value] = true
	if _, ok := sliceMap.slices[value]; ok {
		sliceMap.slices[value] = append(sliceMap.slices[value], nums)
	} else {
		sliceMap.slices[value] = [][]int{nums}
	}
}

func (sliceMap *SliceSumMap) Exist(value int) bool {
	_, ok := sliceMap.value[value]
	return ok
}

func (sliceMap *SliceSumMap) getSlices(value int) [][]int {
	if _, ok := sliceMap.value[value]; ok {
		return sliceMap.slices[value]
	} else {
		return [][]int{}
	}
}
