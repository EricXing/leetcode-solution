package problems

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var l1 = LengthOfLongestSubstring("abcabcbb")
	if l1 != 3 {
		t.Error("LengthOfLongestSubstring case 1 failed")
	}
	var l2 = LengthOfLongestSubstring("bbbbb")
	if l2 != 1 {
		t.Error("LengthOfLongestSubstring case 2 failed")
	}
	var l3 = LengthOfLongestSubstring("pwwkew")
	if l3 != 3 {
		t.Error("LengthOfLongestSubstring case 3 failed")
	}
	var l4 = LengthOfLongestSubstring("vtbiuianjbxbifhpgnhqcvaclbzvqhynazxwhxjtygujodausimufpysqzniinvinsgwsppsylbmqc")
	if l4 != 10 {
		t.Error("LengthOfLongestSubstring case 4 failed")
	}
	var l5 = LengthOfLongestSubstring("dvdf")
	if l5 != 3 {
		t.Error("LengthOfLongestSubstring case 5 failed")
	}
}

func TestLengthOfLongestSubstring1(t *testing.T) {
	var l1 = LengthOfLongestSubstring1("abcabcbb")
	if l1 != 3 {
		t.Error("LengthOfLongestSubstring case 1 failed")
	}
	var l2 = LengthOfLongestSubstring1("bbbbb")
	if l2 != 1 {
		t.Error("LengthOfLongestSubstring case 2 failed")
	}
	var l3 = LengthOfLongestSubstring1("pwwkew")
	if l3 != 3 {
		t.Error("LengthOfLongestSubstring case 3 failed")
	}
	var l4 = LengthOfLongestSubstring1("vtbiuianjbxbifhpgnhqcvaclbzvqhynazxwhxjtygujodausimufpysqzniinvinsgwsppsylbmqc")
	if l4 != 10 {
		t.Error("LengthOfLongestSubstring case 4 failed")
	}
	var l5 = LengthOfLongestSubstring1("dvdf")
	if l5 != 3 {
		t.Error("LengthOfLongestSubstring case 5 failed")
	}
}
