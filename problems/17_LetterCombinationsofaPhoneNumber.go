package problems

// Given a string containing digits from 2-9 inclusive, return all possible letter c
// ombinations that the number could represent.
//
// A mapping of digit to letters (just like on the telephone buttons) is given below.
// Note that 1 does not map to any letters.
//
// Example:
//
// Input: "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// Note:
//
// Although the above answer is in lexicographical order, your answer could be in any order you want.
// https://stackoverflow.com/questions/10262215/how-to-create-cartesian-product-over-arbitrary-groups-of-numbers-in-java
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dict := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	result := make([]string, 0)
	runeSlice := make([][]rune, 0)
	for _, r := range digits {
		runeSlice = append(runeSlice, dict[r])
	}
	combineRune(runeSlice, make([]rune, len(runeSlice)), 0, &result)
	return result
}

func combineRune(input [][]rune, current []rune, k int, result *[]string) {
	if k == len(input) {
		*result = append(*result, string(current))
	} else {
		for j := 0; j < len(input[k]); j++ {
			current[k] = input[k][j]
			combineRune(input, current, k+1, result)
		}
	}
}
