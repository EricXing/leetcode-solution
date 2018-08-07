package problems

// Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
//
// Example 1:
//
//	Input: "babad"
//	Output: "bab"
//	Note: "aba" is also a valid answer.
//
// Example 2:
//
//	Input: "cbbd"
//	Output: "bb"
func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		len1 := expandAround(s, i, i)
		len2 := expandAround(s, i, i+1)
		maxLen := MaxInt(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAround(s string, left, right int) int {
	for ; left > 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return right - left - 1
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func LongestPalindrome1(s string) string {
	if s == "" {
		return ""
	}
	longest := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			sub := s[i : j+1]
			if isPalindrome(sub) && len(sub) > len(longest) {
				longest = sub
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	sl := len(s)
	mid := sl / 2
	var left = mid - 1
	var right int
	if sl%2 == 0 {
		right = mid
	} else {
		right = mid + 1
	}
	for ; left >= 0 && right < sl; left, right = left-1, right+1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}
