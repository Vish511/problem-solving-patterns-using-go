package twopointers

/*
Problem Statement #
Given a sorted array, create a new array containing squares of all the number of the input array in the sorted order.

Example 1:

Input: [-4,-2, -1, 0, 2, 3]
Output: [0, 1, 4, 4, 9]
Example 2:

Input: [-3, -1, 0, 1, 2]
Output: [0 1 1 4 9]
*/

func SquaringASortedArray(arr []int) []int {
	var result = make([]int, len(arr))

	var left int
	var right = len(arr) - 1
	var fillIdx = len(arr) - 1

	for left < right {
		var leftSq = arr[left] * arr[left]
		var rightSq = arr[right] * arr[right]
		var biggerNum int
		if leftSq >= rightSq {
			left++
			biggerNum = leftSq
		} else {
			right--
			biggerNum = rightSq
		}
		result[fillIdx] = biggerNum
		fillIdx--
	}
	return result
}
