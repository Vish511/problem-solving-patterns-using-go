package cyclicsort

/*
Problem Statement #
We are given an unsorted array containing numbers taken from the range 1 to ‘n’. The array can have duplicates, which means some numbers will be missing. Find all those missing numbers.

Example 1:

Input: [2, 3, 1, 8, 2, 3, 5, 1]
Output: 4, 6, 7
Explanation: The array should have all numbers from 1 to 8, due to duplicates 4, 6, and 7 are missing.
Example 2:

Input: [2, 4, 1, 2]
Output: 3
Example 3:

Input: [2, 3, 2, 1]
Output: 4
*/

func FindAllMissingNums(arr []int) []int {
	var i int
	var missingNums []int
	for i < len(arr) {
		if arr[i] != i+1 && arr[i] != arr[arr[i]-1] {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		} else {
			i++
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			missingNums = append(missingNums, i+1)
		}
	}
	return missingNums
}
