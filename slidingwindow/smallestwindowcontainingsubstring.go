package slidingwindow

import (
	"dsa-patterns/helpers"
	"math"
)

/*
Smallest Window containing Substring (hard) #
Given a string and a pattern, find the smallest substring in the given string which has all the characters of the given pattern.

Example 1:

Input: String="aabdec", Pattern="abc"
Output: "abdec"
Explanation: The smallest substring having all characters of the pattern is "abdec"
Example 2:

Input: String="abdabca", Pattern="abc"
Output: "abc"
Explanation: The smallest substring having all characters of the pattern is "abc".
Example 3:

Input: String="adcad", Pattern="abc"
Output: ""
Explanation: No substring in the given string has all characters of the pattern.
*/

func SmallestWindowContainingSubstr(input, pattern string) string {
	var smallestSubStrStart int
	var patternCharMap = make(map[string]int)
	var windowStart, windowEnd int
	var smallestSubStrLen int = math.MaxInt
	for i := 0; i < len(pattern); i++ {
		patternCharMap[string(pattern[i])]++
	}
	var remainingMatches = len(patternCharMap)
	for windowEnd = 0; windowEnd < len(input); windowEnd++ {
		var rightChar = string(input[windowEnd])

		if _, ok := patternCharMap[rightChar]; ok {
			patternCharMap[rightChar]--
			if patternCharMap[rightChar] == 0 {
				remainingMatches--
			}
		}

		for remainingMatches == 0 {
			if windowEnd-windowStart+1 < smallestSubStrLen {
				smallestSubStrLen = windowEnd - windowStart + 1
				smallestSubStrStart = windowStart
			}
			smallestSubStrLen = helpers.Min(smallestSubStrLen, windowEnd-windowStart+1)
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

	if smallestSubStrLen <= len(input) {
		return input[smallestSubStrStart : smallestSubStrStart+smallestSubStrLen]
	}
	return ""

}
