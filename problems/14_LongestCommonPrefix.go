package problems

//Write a function to find the longest common prefix string amongst an array of strings.
//
//If there is no common prefix, return an empty string "".
//
//Example 1:
//
//	Input: ["flower","flow","flight"]
//	Output: "fl"
//
// Example 2:
//
//	Input: ["dog","racecar","car"]
//	Output: ""
//	Explanation: There is no common prefix among the input strings.
//
// Note:
//
//All given inputs are in lowercase letters a-z.
func LongestCommonPrefix(strs []string) string {
	var result = ""
	if len(strs) == 0 {
		return ""
	}
	outter:
		for i := 0; i < len(strs[0]) ; i++ {
			var pre = strs[0][i]
			for j := 1; j < len(strs); j++ {
				if i >= len(strs[j]) || uint8(pre) != strs[j][i] {
					break outter
				}
			}
			result =  result + string(pre)
		}
	return result
}
