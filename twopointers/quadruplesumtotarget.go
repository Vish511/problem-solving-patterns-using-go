package twopointers

import (
	"dsa-patterns/helpers"
)

/*
Quadruple Sum to Target (medium) #
Given an array of unsorted numbers and a target number, find all unique quadruplets in it, whose sum is equal to the target number.

Example 1:

Input: [4, 1, 2, -1, 1, -3], target=1
Output: [-3, -1, 1, 4], [-3, 1, 1, 2]
Explanation: Both the quadruplets add up to the target.
Example 2:

Input: [2, 0, -1, 1, -2, 2], target=2
Output: [-2, 0, 2, 2], [-1, 0, 1, 2]
Explanation: Both the quadruplets add up to the target.
*/

func QuadrupleSumToTarget(arr []int, target int) [][]int {
	arr = helpers.MergeSort(arr)
	var quadruplets [][]int
	for i := 0; i < len(arr)-3; i++ {
		for j := i + 1; j < len(arr)-2; j++ {
			var left = j + 1
			var right = len(arr) - 1
			for left < right {
				var currSum = arr[i] + arr[j] + arr[left] + arr[right]
				if currSum == target {
					quadruplets = append(quadruplets, []int{arr[i], arr[j], arr[left], arr[right]})
					left++
					right--
				} else if currSum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return quadruplets

}
