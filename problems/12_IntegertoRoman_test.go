package problems

import "testing"

func TestIntToRoman(t *testing.T) {
	if IntToRoman(3) == "III" {
		t.Log("12. Integer to Roman: case 1 succeeded.")
	} else {
		t.Error("12. Integer to Roman: case 1 failed.")
	}
	if IntToRoman(4) == "IV" {
		t.Log("12. Integer to Roman: case 2 succeeded.")
	} else {
		t.Error("12. Integer to Roman: case 2 failed.")
	}
	if IntToRoman(9) == "IX" {
		t.Log("12. Integer to Roman: case 3 succeeded.")
	} else {
		t.Error("12. Integer to Roman: case 3 failed.")
	}
	if IntToRoman(58) == "LVIII" {
		t.Log("12. Integer to Roman: case 4 succeeded.")
	} else {
		t.Error("12. Integer to Roman: case 4 failed.")
	}
	if IntToRoman(1994) == "MCMXCIV" {
		t.Log("12. Integer to Roman: case 5 succeeded.")
	} else {
		t.Error("12. Integer to Roman: case 5 failed.")
	}
}
