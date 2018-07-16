package problems

import "testing"

func TestRomanToInt(t *testing.T) {
	if RomanToInt("III") != 3 {
		t.Error("Roman to Int: case 1 failed.")
	}
	if RomanToInt("IV") != 4 {
		t.Error("Roman to Int: case 2 failed.")
	}
	if RomanToInt("IX") != 9 {
		t.Error("Roman to Int: case 3 failed.")
	}
	if RomanToInt("LVIII") != 58 {
		t.Error("Roman to Int: case 4 failed.")
	}
	if RomanToInt("MCMXCIV") != 1994 {
		t.Error("Roman to Int: case 5 failed.")
	}
}
