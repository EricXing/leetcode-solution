package problems

import "sort"

// Alice has a hand of cards, given as an array of integers.
// Now she wants to rearrange the cards into groups so that
// each group is size W, and consists of W consecutive cards.
// Return true if and only if she can.
//
// Example 1:
//	Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
//	Output: true
//	Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8].
//
// Example 2:
//	Input: hand = [1,2,3,4,5], W = 4
//	Output: false
//	Explanation: Alice's hand can't be rearranged into groups of 4.
//
// Note:
//
//	1 <= hand.length <= 10000
//	0 <= hand[i] <= 10^9
//	1 <= W <= hand.length
func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	}
	countMap := make(map[int]int)
	for _, val := range hand {
		count, exist := countMap[val]
		if exist {
			countMap[val] = count + 1
		} else {
			countMap[val] = 1
		}
	}
	sort.Sort(sort.IntSlice(hand))
	for _, val := range hand {
		count := countMap[val]
		if count > 0 {
			countMap[val] = count - 1
			for i := 1; i < W; i++ {
				c := countMap[val+i]
				if c > 0 {
					countMap[val+i] = c - 1
				} else {
					return false
				}
			}
		}
	}
	return true
}
