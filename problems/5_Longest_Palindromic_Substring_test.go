package problems

import "testing"

func TestLongestPalindrome(t *testing.T) {
	if LongestPalindrome("babad") == "aba" {
		t.Log("Longest Palindromic Substring: case 1 succeeded.")
	} else {
		t.Error("Longest Palindromic Substring: case 1 failed.")
	}
	if LongestPalindrome("cbbd") == "bb" {
		t.Log("Longest Palindromic Substring: case 2 succeeded.")
	} else {
		t.Error("Longest Palindromic Substring: case 2 failed.")
	}
	//if LongestPalindrome("bb") == "bb" {
	//	t.Log("Longest Palindromic Substring: case 3 succeeded.")
	//} else {
	//	t.Error("Longest Palindromic Substring: case 3 failed.")
	//}
}

func TestLongestPalindrome1(t *testing.T) {
	if LongestPalindrome1("babad") == "bab" {
		t.Log("Longest Palindromic Substring: case 1 succeeded.")
	} else {
		t.Error("Longest Palindromic Substring: case 1 failed.")
	}
	if LongestPalindrome1("cbbd") == "bb" {
		t.Log("Longest Palindromic Substring: case 2 succeeded.")
	} else {
		t.Error("Longest Palindromic Substring: case 2 failed.")
	}
	if LongestPalindrome1("bb") == "bb" {
		t.Log("Longest Palindromic Substring: case 3 succeeded.")
	} else {
		t.Error("Longest Palindromic Substring: case 3 failed.")
	}
}
