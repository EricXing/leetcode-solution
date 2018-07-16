package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	if "fl" !=  LongestCommonPrefix([]string{"flower","flow","flight"}){
		t.Error("Longest Common Prefix: case 1 failed")
	}
	if "" !=  LongestCommonPrefix([]string{"dog","racecar","car"}){
		t.Error("Longest Common Prefix: case 2 failed")
	}
	if "" != LongestCommonPrefix([]string{}){
		t.Error("Longest Common Prefix: case 3 failed")
	}
	if "a" != LongestCommonPrefix([]string{"a"}){
		t.Error("Longest Common Prefix: case 4 failed")
	}
	if "" != LongestCommonPrefix([]string{"c","acc","ccc"}){
		t.Error("Longest Common Prefix: case 5 failed")
	}
}
