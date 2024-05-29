package slidingwindow

import (
	"dsa-patterns/helpers"
)

/*
Problem Statement #
Given a string, find the length of the longest substring in it with no more than K distinct characters.

Example 1:

Input: String="araaci", K=2
Output: 4
Explanation: The longest substring with no more than '2' distinct characters is "araa".
Example 2:

Input: String="araaci", K=1
Output: 2
Explanation: The longest substring with no more than '1' distinct characters is "aa".
Example 3:

Input: String="cbbebi", K=3
Output: 5
Explanation: The longest substrings with no more than '3' distinct characters are "cbbeb" & "bbebi"
*/

func LongestSubStringWithKDistinctChars(input string, k int) int {
	var windowStart int
	var distinctCharsMap = make(map[string]int)
	var lenOfLongestSubstr int
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		var char = string(input[windowEnd])
		distinctCharsMap[char]++

		if len(distinctCharsMap) == k {
			lenOfLongestSubstr = helpers.Max(lenOfLongestSubstr, windowEnd-windowStart+1)
		}

		for len(distinctCharsMap) > k {
			var leftChar = string(input[windowStart])
			distinctCharsMap[leftChar]--
			if distinctCharsMap[leftChar] == 0 {
				delete(distinctCharsMap, leftChar)
			}
			windowStart++
		}
	}
	return lenOfLongestSubstr
}
