package slidingwindow

/*
Problem Statement #
Given an array of positive numbers and a positive number ‘k’, find the maximum sum of any contiguous subarray of size ‘k’.

Example 1:

Input: [2, 1, 5, 1, 3, 2], k=3
Output: 9
Explanation: Subarray with maximum sum is [5, 1, 3].
Example 2:

Input: [2, 3, 4, 1, 5], k=2
Output: 7
Explanation: Subarray with maximum sum is [3, 4].
*/

func MaxSumOfSubArray(arr []int, k int) int {
	if len(arr) < k {
		return -1
	}
	var windowStart int
	var currSum int
	var maxSum int
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		currSum += arr[windowEnd]
		if currSum > maxSum {
			maxSum = currSum
		}
		if windowEnd >= k-1 {
			currSum -= arr[windowStart]
			windowStart++
		}
	}
	return maxSum
}
