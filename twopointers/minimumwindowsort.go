package twopointers

import (
	"dsa-patterns/helpers"
	"math"
)

/*
Minimum Window Sort (medium) #
Given an array, find the length of the smallest subarray in it which when sorted will sort the whole array.

Example 1:

Input: [1, 2, 5, 3, 7, 10, 9, 12]
Output: 5
Explanation: We need to sort only the subarray [5, 3, 7, 10, 9] to make the whole array sorted
Example 2:

Input: [1, 3, 2, 0, -1, 7, 10]
Output: 5
Explanation: We need to sort only the subarray [1, 3, 2, 0, -1] to make the whole array sorted
Example 3:

Input: [1, 2, 3]
Output: 0
Explanation: The array is already sorted
Example 4:

Input: [3, 2, 1]
Output: 3
Explanation: The whole array needs to be sorted.
*/

func MinWindowSort(arr []int) int {
	var low int
	var high = len(arr) - 1
	var subArrayMin = math.MaxInt
	var subArrayMax = math.MinInt

	for low < len(arr)-1 && arr[low] <= arr[low+1] {
		low++
	}

	if low == len(arr)-1 {
		return 0
	}

	for high > 0 && arr[high] >= arr[high-1] {
		high--
	}

	for i := low; i <= high; i++ {
		subArrayMin = helpers.Min(subArrayMin, arr[i])
		subArrayMax = helpers.Max(subArrayMax, arr[i])
	}

	for low > 0 && arr[low-1] > subArrayMin {
		low--
	}

	for high < len(arr)-1 && subArrayMax > arr[high+1] {
		high++
	}

	return high - low + 1
}
