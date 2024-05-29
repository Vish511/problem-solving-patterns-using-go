package slidingwindow

import (
	"dsa-patterns/helpers"
)

/*
Problem Statement #
Given a string, find the length of the longest substring which has no repeating characters.

Example 1:

Input: String="aabccbb"
Output: 3
Explanation: The longest substring without any repeating characters is "abc".
Example 2:

Input: String="abbbb"
Output: 2
Explanation: The longest substring without any repeating characters is "ab".
Example 3:

Input: String="abccde"
Output: 3
Explanation: Longest substrings without any repeating characters are "abc" & "cde".

"bacabcd"
*/

func NoRepeatSubstring(input string) int {
	var windowStart int
	var longestSubStrLen int
	var visitedCharsMap = make(map[string]int)

	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		var char = string(input[windowEnd])

		if v, ok := visitedCharsMap[char]; ok {
			windowStart = helpers.Max(windowStart, v+1)
		}
		visitedCharsMap[char] = windowEnd
		longestSubStrLen = helpers.Max(longestSubStrLen, windowEnd-windowStart+1)
	}
	return longestSubStrLen
}
