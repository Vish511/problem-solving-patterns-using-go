package slidingwindow

/*
Permutation in a String (hard) #
Given a string and a pattern, find out if the string contains any permutation of the pattern.

Permutation is defined as the re-arranging of the characters of the string. For example, “abc” has the following six permutations:

abc
acb
bac
bca
cab
cba
If a string has ‘n’ distinct characters it will have
𝑛
!
n! permutations.

Example 1:

Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.
Example 2:

Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.
Example 3:

Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.
Example 4:

Input: String="aaacb", Pattern="abc"
Output: true
Explanation: The string contains "acb" which is a permutation of the given pattern.
*/

func PermutationsInAString(input, pattern string) bool {
	var patternCharMap = make(map[string]int)
	for i := 0; i < len(pattern); i++ {
		patternCharMap[string(pattern[i])]++
	}
	var remainingMatches = len(patternCharMap)
	var windowStart int
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		var rightChar = string(input[windowEnd])
		if _, ok := patternCharMap[rightChar]; ok {
			patternCharMap[rightChar]--
			if patternCharMap[rightChar] == 0 {
				remainingMatches--
			}
		}

		if remainingMatches == 0 {
			return true
		}

		if windowEnd >= len(pattern)-1 {
			var leftChar = string(input[windowStart])
			if _, ok := patternCharMap[leftChar]; ok {
				patternCharMap[leftChar]++
				if patternCharMap[leftChar] > 0 {
					remainingMatches++
				}
			}
			windowStart++
		}

	}
	return false
}
