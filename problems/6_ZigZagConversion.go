package problems

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
//	P   A   H   N
//	A P L S I I G
//	Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
//	string convert(string s, int numRows);
//
// Example 1:
//
//	Input: s = "PAYPALISHIRING", numRows = 3
//	Output: "PAHNAPLSIIGYIR"
//
// Example 2:
//
//	Input: s = "PAYPALISHIRING", numRows = 4
//	Output: "PINALSIGYAHRPI"
//	Explanation:
//
//	P     I    N
//	A   L S  I G
//	Y A   H R
//	P     I
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, MinInt(len(s), numRows))
	currRow := 0
	goingDown := false

	for _, c := range s {
		rows[currRow] += string(c)
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			currRow += 1
		} else {
			currRow -= 1
		}
	}
	result := ""
	for _, val := range rows {
		result += val
	}
	return result
}
