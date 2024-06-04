package twopointers

import (
	"dsa-patterns/helpers"
)

/*
Problem Statement #
Given an array of unsorted numbers, find all unique triplets in it that add up to zero.

Example 1:

Input: [-3, 0, 1, 2, -1, 1, -2]
[-3,-2,-1,0,1,1,2]
Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
Explanation: There are four unique triplets whose sum is equal to zero.
Example 2:

Input: [-5, 2, -1, -2, 3]
Output: [[-5, 2, 3], [-2, -1, 3]]
Explanation: There are two unique triplets whose sum is equal to zero.
*/

func TripletWithGivenSum(arr []int, target int) [][]int {
	var result [][]int
	var sortedArr = helpers.MergeSort(arr)
	for i := 0; i < len(sortedArr); i++ {
		var requiredSum = target - sortedArr[i]
		var pairWithSum = PairWithTargetSum(sortedArr[i+1:], requiredSum)
		for p := 0; p < len(pairWithSum); p++ {
			var triplet = []int{sortedArr[i], pairWithSum[p][0], pairWithSum[p][1]}
			result = append(result, triplet)
		}
	}
	return result
}
