package problems


//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3.
//Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func LengthOfLongestSubstring(s string) int {
	var maxLength = 0
	for start := 0; start < len(s); start++ {
		for end := len(s); end > start; end-- {
			substring := s[start: end]
			length := len(substring)
			letterMap := make(map[string]bool)
			for _, v := range substring{
				letterMap[string(v)] = true
			}
			if len(letterMap) == length && length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}

func LengthOfLongestSubstring1(s string) int {
	var maxLength = 0
	var letterMap = make(map[string]int)
	var start = 0
	for i := 0; i< len(s); i++ {
		c := string(s[i])
		if v, ok := letterMap[c]; ok && v >= start {
			start = v + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		letterMap[c] = i
	}
	return maxLength
}
