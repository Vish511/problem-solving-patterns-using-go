package cyclicsort

/*
Problem Statement #
We are given an array containing ‘n’ distinct numbers taken from the range 0 to ‘n’. Since the array has only ‘n’ numbers out of the total ‘n+1’ numbers, find the missing number.

Example 1:

Input: [4, 0, 3, 1]
Output: 2

Example 2:
Input: [8, 3, 5, 2, 4, 6, 0, 1]
Output: 7
*/

func FindTheMissingNum(arr []int) int {
	var i int
	for i < len(arr) {
		if arr[i] != i && arr[i] < len(arr) {
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			return i
		}
	}
	return -1
}
