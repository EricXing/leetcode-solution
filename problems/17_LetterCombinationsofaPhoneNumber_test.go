package problems

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	if reflect.DeepEqual(LetterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}) {
		t.Log("17. Letter Combinations of a Phone Number: case 1 succeeded.")
	} else {
		t.Error("17. Letter Combinations of a Phone Number: case 1 failed.")
	}
}
