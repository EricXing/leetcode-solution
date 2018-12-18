package problems

import "testing"

func TestIsNStraightHand(t *testing.T) {
	if isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3) {
		t.Log("846. Hand of Straights: case 1 succeeded.")
	} else {
		t.Error("846. Hand of Straights: case 1 failed.")
	}
	if !isNStraightHand([]int{1, 2, 3, 4, 5}, 4) {
		t.Log("846. Hand of Straights: case 2 succeeded.")
	} else {
		t.Error("846. Hand of Straights: case 2 failed.")
	}
}
