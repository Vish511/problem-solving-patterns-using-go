package slidingwindow

/*
String Anagrams (hard) #
Given a string and a pattern, find all anagrams of the pattern in the given string.

Anagram is actually a Permutation of a string. For example, “abc” has the following six anagrams:

abc
acb
bac
bca
cab
cba
Write a function to return a list of starting indices of the anagrams of the pattern in the given string.

Example 1:

Input: String="ppqp", Pattern="pq"
Output: [1, 2]
Explanation: The two anagrams of the pattern in the given string are "pq" and "qp".
Example 2:

Input: String="abbcabc", Pattern="abc"
Output: [2, 3, 4]
Explanation: The three anagrams of the pattern in the given string are "bca", "cab", and "abc".
*/

func StringAnagrams(input, pattern string) []int {
	var patternCharMap = make(map[string]int)
	for i := 0; i < len(pattern); i++ {
		patternCharMap[string(pattern[i])]++
	}
	var remainingMatches = len(patternCharMap)
	var windowStart int
	var result []int
	for windowEnd := 0; windowEnd < len(input); windowEnd++ {
		var rightChar = string(input[windowEnd])
		if _, ok := patternCharMap[rightChar]; ok {
			patternCharMap[rightChar]--
			if patternCharMap[rightChar] == 0 {
				remainingMatches--
			}
		}
		if remainingMatches == 0 {
			result = append(result, windowStart)
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
	return result

}
