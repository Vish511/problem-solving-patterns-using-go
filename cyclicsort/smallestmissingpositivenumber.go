package cyclicsort

import "fmt"

/*
Find the Smallest Missing Positive Number (medium) #
Given an unsorted array containing numbers, find the smallest missing positive number in it.

Example 1:

Input: [-3, 1, 5, 4, 2]
Output: 3
Explanation: The smallest missing positive number is '3'
Example 2:

Input: [3, -2, 0, 1, 2]
Output: 4
Example 3:

Input: [3, 2, 5, 1]
Output: 4
*/

func FindSmallestMissingPosNum(arr []int) int {
	var i int
	for i < len(arr) {
		if arr[i] != i+1 && arr[i] > 0 && arr[i] < len(arr) {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	fmt.Println(arr)
	return -1
}
