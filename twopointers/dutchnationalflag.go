package twopointers

/*
Problem Statement #
Given an array containing 0s, 1s and 2s, sort the array in-place. You should treat numbers of the array as objects, hence, we can’t count 0s, 1s, and 2s to recreate the array.

The flag of the Netherlands consists of three colors: red, white and blue; and since our input array also consists of three different numbers that is why it is called Dutch National Flag problem.

Example 1:

Input: [1, 0, 2, 1, 0]
Output: [0 0 1 1 2]
Example 2:

Input: [2, 2, 0, 1, 2, 0]
Output: [0 0 1 2 2 2 ]
*/

func DutchNationalFlag(arr []int) []int {
	var low int
	var high = len(arr) - 1
	var i int

	for i <= high {
		if arr[i] == 0 {
			arr[low], arr[i] = arr[i], arr[low]
			low++
			i++
		} else if arr[i] == 2 {
			arr[high], arr[i] = arr[i], arr[high]
			high--
		} else {
			i++
		}
	}
	return arr
}
