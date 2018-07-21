package problems

import "testing"

func TestAddBinary(t *testing.T) {
	if AddBinary("11", "1") == "100" {
		t.Log("Add Binary: case 1 succeeded.")
	} else {
		t.Error("Add Binary: case 1 failed.")
	}
	if AddBinary("1010", "1011") == "10101" {
		t.Log("Add Binary: case 1 succeeded.")
	} else {
		t.Error("Add Binary: case 1 failed.")
	}
}
