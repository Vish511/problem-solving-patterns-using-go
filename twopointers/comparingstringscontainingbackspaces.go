package twopointers

/*
Comparing Strings containing Backspaces (medium) #
Given two strings containing backspaces (identified by the character ‘#’), check if the two strings are equal.

Example 1:

Input: str1="xy#z", str2="xzz#"
Output: true
Explanation: After applying backspaces the strings become "xz" and "xz" respectively.
Example 2:

Input: str1="xy#z", str2="xyz#"
Output: false
Explanation: After applying backspaces the strings become "xz" and "xy" respectively.
Example 3:

Input: str1="xp#", str2="xyz##"
Output: true
Explanation: After applying backspaces the strings become "x" and "x" respectively.
In "xyz##", the first '#' removes the character 'z' and the second '#' removes the character 'y'.
Example 4:

Input: str1="xywrrmp", str2="xywrrmu#p"
Output: true
Explanation: After applying backspaces the strings become "xywrrmp" and "xywrrmp" respectively.

*/

func CompareStringsContainingBackspaces(str1, str2 string) bool {
	var p1 = len(str1) - 1
	var p2 = len(str2) - 1
	for p1 >= 0 && p2 >= 0 {
		i1 := GetNextUndeletedCharIdx(str1, p1)
		i2 := GetNextUndeletedCharIdx(str2, p2)
		if i1 < 0 || i2 < 0 {
			return true
		}

		if i1 < 0 || i2 < 0 {
			return false
		}

		if str1[i1] != str2[i2] {
			return false

		}
		p1 = i1 - 1
		p2 = i2 - 1
	}

	if GetNextUndeletedCharIdx(str1, p1) != GetNextUndeletedCharIdx(str2, p2) {
		return false
	}

	return true
}

func GetNextUndeletedCharIdx(input string, pos int) int {
	var backspaceCnt int
	for pos >= 0 {
		if string(input[pos]) == "#" {
			backspaceCnt++
		} else if backspaceCnt > 0 {
			backspaceCnt--
		} else {
			break
		}
		pos--
	}
	return pos
}
