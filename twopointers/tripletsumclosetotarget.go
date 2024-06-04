package twopointers

import (
	"dsa-patterns/helpers"
	"math"
)

/*
Problem Statement #
Given an array of unsorted numbers and a target number, find a triplet in the array whose sum is as close to the target number as possible, return the sum of the triplet. If there are more than one such triplet, return the sum of the triplet with the smallest sum.

Example 1:

Input: [-2, 0, 1, 2], target=2
Output: 1
Explanation: The triplet [-2, 1, 2] has the closest sum to the target.
Example 2:

Input: [-3, -1, 1, 2], target=1
Output: 0
Explanation: The triplet [-3, 1, 2] has the closest sum to the target.
Example 3:

Input: [1, 0, 1, 1], target=100
Output: 3
Explanation: The triplet [1, 1, 1] has the closest sum to the target.
*/

func TripletSumCloseToTarget(arr []int, target int) int {
	var leastDiff = math.MaxInt
	var triplet []int
	arr = helpers.MergeSort(arr)
	for i := 0; i < len(arr); i++ {
		var left int = i + 1
		var right = len(arr) - 1
		for left < right {
			var currSum = arr[i] + arr[left] + arr[right]
			var absDiff = int(math.Abs(float64(target - currSum)))
			if absDiff == 0 {
				return arr[i] + arr[left] + arr[right]
			}
			if absDiff < leastDiff {
				leastDiff = absDiff
				triplet = []int{}
				triplet = append(triplet, arr[i], arr[left], arr[right])
			}
			if currSum < target {
				left++
			} else {
				right--
			}
		}
	}
	if len(triplet) == 3 {
		return triplet[0] + triplet[1] + triplet[2]
	}
	return math.MaxInt
}
