package problems

import "testing"

func TestIsMatch(t *testing.T) {
	if !IsMatch("aa", "a") {
		t.Log("10. Regular Expression Matching: case 1 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 1 failed.")
	}
	if IsMatch("aa", "a*") {
		t.Log("10. Regular Expression Matching: case 2 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 2 failed.")
	}
	if IsMatch("ab", ".*") {
		t.Log("10. Regular Expression Matching: case 3 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 3 failed.")
	}
	if IsMatch("aab", "c*a*b") {
		t.Log("10. Regular Expression Matching: case 4 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 4 failed.")
	}
	if !IsMatch("mississippi", "mis*is*p*.") {
		t.Log("10. Regular Expression Matching: case 5 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 5 failed.")
	}
	if !IsMatch("abcd", "d*") {
		t.Log("10. Regular Expression Matching: case 6 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 6 failed.")
	}
	if !IsMatch("aaa", "ab*a") {
		t.Log("10. Regular Expression Matching: case 7 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 7 failed.")
	}
	if !IsMatch("ab", ".*c") {
		t.Log("10. Regular Expression Matching: case 8 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 8 failed.")
	}
	if IsMatch("aaa", "a*a") {
		t.Log("10. Regular Expression Matching: case 9 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 9 failed.")
	}
}

func TestIsMatch1(t *testing.T) {
	if !IsMatch1("aa", "a") {
		t.Log("10. Regular Expression Matching: case 1 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 1 failed.")
	}
	if IsMatch1("aa", "a*") {
		t.Log("10. Regular Expression Matching: case 2 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 2 failed.")
	}
	if IsMatch1("ab", ".*") {
		t.Log("10. Regular Expression Matching: case 3 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 3 failed.")
	}
	if IsMatch1("aab", "c*a*b") {
		t.Log("10. Regular Expression Matching: case 4 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 4 failed.")
	}
	if !IsMatch1("mississippi", "mis*is*p*.") {
		t.Log("10. Regular Expression Matching: case 5 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 5 failed.")
	}
	if !IsMatch1("abcd", "d*") {
		t.Log("10. Regular Expression Matching: case 6 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 6 failed.")
	}
	if !IsMatch1("aaa", "ab*a") {
		t.Log("10. Regular Expression Matching: case 7 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 7 failed.")
	}
	if !IsMatch1("ab", ".*c") {
		t.Log("10. Regular Expression Matching: case 8 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 8 failed.")
	}
	if IsMatch1("aaa", "a*a") {
		t.Log("10. Regular Expression Matching: case 9 succeeded.")
	} else {
		t.Error("10. Regular Expression Matching case 9 failed.")
	}
}
