package problems

// Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.
//
// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).
//
// Note:
//
//	* s could be empty and contains only lowercase letters a-z.
//	* p could be empty and contains only lowercase letters a-z, and characters like . or *.
//
// Example 1:
//
//	Input:
//	s = "aa"
//	p = "a"
//	Output: false
//	Explanation: "a" does not match the entire string "aa".
//
// Example 2:
//
//	Input:
//	s = "aa"
//	p = "a*"
//	Output: true
//
// Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
//
// Example 3:
//
//	Input:
//	s = "ab"
//	p = ".*"
//	Output: true
//	Explanation: ".*" means "zero or more (*) of any character (.)".
//
// Example 4:
//
//	Input:
//	s = "aab"
//	p = "c*a*b"
//	Output: true
//	Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
//
// Example 5:
//
//	Input:
//	s = "mississippi"
//	p = "mis*is*p*."
//	Output: false
func IsMatch(s string, p string) bool {
	p = compile(p)
	j := 0

	for i := 0; i < len(p); i++ {
		// string exhausted before pattern, and there is char not followed by *
		if j == len(s) {
			if p[i] != '*' || i < len(p)-1 && p[i+1] != '*' {
				return false
			} else {
				continue
			}
		}
		for ; j < len(s); j++ {
			// current character match failed.
			if p[i] != s[j] && p[i] != '.' {
				// Two conditions match can continue
				// 1. Next char of patter is *
				// 2. There is still pattern left to match
				if i < len(p)-1 && p[i+1] == '*' {
					// Skip *
					i++
					break
				} else if j > 0 && p[i] == s[j-1] {
					break
				} else {
					return false
				}
			}
			// Both pattern and string meet the end
			if i == len(p)-1 && j == len(s)-1 {
				j++
				break
			}
			if i == len(p)-1 || p[i+1] != '*' {
				break
			}
		}
	}
	return j == len(s)

}

func IsMatch1(s string, p string) bool {
	p = compile(p)
	j := 0
	for i := 0; i < len(p); i++ {
		// string exhausted before pattern, and there is char not followed by *
		if j == len(s) {
			if p[i] != '*' || i < len(p)-1 && p[i+1] != '*' {
				return false
			} else {
				continue
			}
		}
		step := j + 1
		char := p[i]
		// next pattern is *, zero or multiple chars can be tested
		if i < len(p)-1 && p[i+1] == '*' {
			step = len(s)
			i++
		}
		for ; j < step; j++ {
			if char != '.' && char != s[j] {
				break
			}
		}
	}
	return j == len(s)
}

func compile(p string) string {
	newP := make([]byte, 0, len(p))
	mark := -1
	for i:=0;i<len(p);i++{
		// The last pattern of p is .*, so the rest can be skipped
		if p[i] == '.' && i < len(p) -1 && p[i+1] =='*' {
			newP = append(newP, '*')
			break
		}
		if p[i] !='*'{
			mark = i
		}
		if mark <len(p) && p[i] == p[mark]{
			continue
		}
		newP = append(newP, p[i])
	}
	return string(newP)
}
