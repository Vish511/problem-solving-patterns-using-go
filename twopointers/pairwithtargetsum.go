package twopointers

/*
Problem Statement #
Given an array of sorted numbers and a target sum, find a pair in the array whose sum is equal to the given target.

Write a function to return the indices of the two numbers (i.e. the pair) such that they add up to the given target.

Example 1:

Input: [1, 2, 3, 4, 6], target=6
Output: [1, 3]
Explanation: The numbers at index 1 and 3 add up to 6: 2+4=6
Example 2:

Input: [2, 5, 9, 11], target=11
Output: [0, 2]
Explanation: The numbers at index 0 and 2 add up to 11: 2+9=11
*/

func PairWithTargetSum(arr []int, target int) [][]int {
	var result [][]int
	var left int
	var right = len(arr) - 1

	for left < right {
		var currSum = arr[left] + arr[right]
		if currSum == target {
			result = append(result, []int{arr[left], arr[right]})
			left++
			right--
		} else if currSum < target {
			left++
		} else {
			right--
		}
	}
	return result
}
