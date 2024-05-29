package slidingwindow

/*
Words Concatenation (hard) #
Given a string and a list of words, find all the starting indices of substrings in the given string that are a concatenation of all the given words exactly once without any overlapping of words. It is given that all words are of the same length.

Example 1:

Input: String="catfoxcat", Words=["cat", "fox"]
Output: [0, 3]
Explanation: The two substring containing both the words are "catfox" & "foxcat".
Example 2:

Input: String="catcatfoxfox", Words=["cat", "fox"]
Output: [3]
Explanation: The only substring containing both the words is "catfox".
*/

func WordsConcatenation(input string, words []string) []int {
	var wordsMap = make(map[string]int)
	var wordLen = len(words[0])
	var windowStart int
	var matchedCnt int
	var result []int
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]]++
	}

	for windowEnd := wordLen*len(words) - 1; windowEnd < len(input); windowEnd++ {
		var wordsSeen = make(map[string]int)
		for i := windowStart; i < windowEnd; i += wordLen {
			var substr = string(input[i : i+wordLen])
			wordsSeen[substr]++
			_, subStrExistsInWord := wordsMap[substr]
			if !subStrExistsInWord || wordsSeen[substr] > wordsMap[substr] {
				break
			}

			if wordsSeen[substr] == wordsMap[substr] {
				matchedCnt++
			}
			if matchedCnt == len(wordsMap) {
				result = append(result, windowStart)
			}
		}
		matchedCnt = 0
		windowStart++
	}
	return result
}
