package cyclicsort

/*
Problem Statement #
We are given an unsorted array containing ‘n’ numbers taken from the range 1 to ‘n’. The array has some duplicates, find all the duplicate numbers without using any extra space.

Example 1:
Input: [3, 4, 4, 5, 5]
Output: [4, 5]

Example 2:
Input: [5, 4, 7, 2, 3, 5, 3]
Output: [3, 5]
*/

func FindAllDuplicateNums(arr []int) []int {
	var duplicates []int
	var i int
	for i < len(arr) {
		if arr[i] != i+1 {
			if arr[i] != arr[arr[i]-1] {
				arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
			} else {
				duplicates = append(duplicates, arr[i])
				i++
			}
		} else {
			i++
		}
	}
	return duplicates
}
