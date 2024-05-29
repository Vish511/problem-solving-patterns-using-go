package slidingwindow

import (
	"dsa-patterns/helpers"
)

/*
Problem Statement #
Given a string with lowercase letters only, if you are allowed to replace no more than ‘k’ letters with any letter, find the length of the longest substring having the same letters after replacement.

Example 1:

Input: String="aabcebb", k=2
Output: 5
Explanation: Replace the two 'c' with 'b' to have a longest repeating substring "bbbbb".
Example 2:

Input: String="abbcb", k=1
Output: 4
Explanation: Replace the 'c' with 'b' to have a longest repeating substring "bbbb".
Example 3:

Input: String="abccde", k=1
Output: 3
Explanation: Replace the 'b' or 'd' with 'c' to have the longest repeating substring "ccc".
*/

func LongestSubStrWithSameLettersAfterReplacement(input string, k int) int {
	var windowStart int
	var maxLength int
	var freqMap = make(map[string]int)
	var maxLetterRepeatCnt int
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		var char = string(input[windowEnd])
		freqMap[char]++
		maxLetterRepeatCnt = helpers.Max(maxLetterRepeatCnt, freqMap[char])

		if windowEnd-windowStart+1-maxLetterRepeatCnt > k {
			var leftChar = string(input[windowStart])
			freqMap[leftChar]--
			windowStart += 1
		}

		maxLength = helpers.Max(maxLength, windowEnd-windowStart+1)
	}
	return maxLength
}

// func getMaxRepeatCntInWindow(freqMap map[string]int) int {
// 	var max int
// 	for _, v := range freqMap {
// 		max = helpers.Max(max, v)
// 	}
// 	return max
// }
