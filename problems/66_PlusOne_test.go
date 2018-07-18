package problems

import "testing"

func TestPlusOne1(t *testing.T) {
	result := PlusOne([]int{1, 2, 3})
	if len(result) == 3 && result[0] == 1 && result[1] == 2 && result[2] == 4 {
		t.Log("Plus One: case 1 succeeded.")
	} else {
		t.Error("Plus One: case 1 failed.")
	}
}

func TestPlusOne2(t *testing.T) {
	result := PlusOne([]int{4, 3, 2, 1})
	if len(result) == 4 && result[0] == 4 && result[1] == 3 && result[2] == 2 && result[3] == 2 {
		t.Log("Plus One: case 2 succeeded.")
	} else {
		t.Error("Plus One: case 2 failed.")
	}
}

func TestPlusOne3(t *testing.T) {
	result := PlusOne([]int{ 9, 9})
	if len(result) == 3 && result[0] == 1 && result[1] == 0 && result[2] == 0 {
		t.Log("Plus One: case 3 succeeded.")
	} else {
		t.Error("Plus One: case 3 failed.")
	}
}
