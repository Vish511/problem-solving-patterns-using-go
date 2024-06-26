package slidingwindow

import (
	"dsa-patterns/helpers"
	"fmt"
)

/*
Problem Statement #
Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s, find the length of the longest contiguous subarray having all 1s.

Example 1:

Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
Output: 6
Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.
Example 2:

Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
Output: 9
Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.

*/

func LongestSubStrWithOnesAfterReplacement(arr []int, k int) int {
	var windowStart int
	var onesCount int
	var maxLen int
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		if arr[windowEnd] == 1 {
			onesCount++
		}

		if windowEnd-windowStart+1-onesCount > k {
			if arr[windowStart] == 1 {
				onesCount--
			}
			windowStart++
		}

		if windowEnd-windowStart+1-onesCount <= k {
			fmt.Println(windowEnd, windowStart, windowEnd-windowStart+1)
			maxLen = helpers.Max(maxLen, windowEnd-windowStart+1)
		}
	}
	return maxLen
}
