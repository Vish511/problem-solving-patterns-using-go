package twopointers

import (
	"dsa-patterns/helpers"
)

/*
Problem Statement #
Given an array arr of unsorted numbers and a target sum, count all triplets in it such that arr[i] + arr[j] + arr[k] < target where i, j, and k are three different indices. Write a function to return the count of such triplets.

Example 1:

Input: [-1, 0, 2, 3], target=3
Output: 2
Explanation: There are two triplets whose sum is less than the target: [-1, 0, 3], [-1, 0, 2]
Example 2:

Input: [-1, 4, 2, 1, 3], target=5
Output: 4
Explanation: There are four triplets whose sum is less than the target:
   [-1, 1, 4], [-1, 1, 3], [-1, 1, 2], [-1, 2, 3]
*/

func TripletsWithSmallerSumThanTarget(arr []int, target int) int {
	arr = helpers.MergeSort(arr)
	var count = 0
	for i := 0; i < len(arr); i++ {
		var left int = i + 1
		var right = len(arr) - 1
		for left < right {
			var currSum = arr[i] + arr[left] + arr[right]
			if currSum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}
	return count
}
