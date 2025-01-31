package cyclicsort

import (
	"dsa-patterns/helpers"
	"math"
)

/*
Find the First K Missing Positive Numbers (hard) #
Given an unsorted array containing numbers and a number ‘k’, find the first ‘k’ missing positive numbers in the array.

Example 1:

Input: [3, -1, 4, 5, 5], k=3
Output: [1, 2, 6]
Explanation: The smallest missing positive numbers are 1, 2 and 6.
Example 2:

Input: [2, 3, 4], k=3
Output: [1, 5, 6]
Explanation: The smallest missing positive numbers are 1, 5 and 6.
Example 3:

Input: [-2, -3, 4], k=2
Output: [1, 2]
Explanation: The smallest missing positive numbers are 1 and 2.
*/

func FindFirstKPosMissingNums(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return nil
	}
	var missingNums []int

	var i int
	for i < len(arr) {
		if arr[i] < len(arr) && arr[i] > 0 && arr[i] != arr[arr[i]-1] {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		} else {
			i++
		}
	}

	var extraNums = make(map[int]struct{})
	for i := 0; i < len(arr) && len(missingNums) < k; i++ {
		if arr[i] != i+1 {
			missingNums = append(missingNums, i+1)
			extraNums[arr[i]] = struct{}{}
		}
	}

	i = 1
	for len(missingNums) < k {
		var candidateNum = i + len(arr)
		if _, ok := extraNums[candidateNum]; !ok {
			missingNums = append(missingNums, candidateNum)
		}
		i++
	}

	return missingNums

}

func FindFirstKPosMissingNumsAlt(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return nil
	}
	var smallestNum int = math.MaxInt
	var largestNum int = math.MinInt
	var missingNums []int
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			largestNum = helpers.Max(largestNum, arr[i])
			smallestNum = helpers.Min(smallestNum, arr[i])
		}
	}

	var i int
	for i < len(arr) {
		if arr[i] != i+smallestNum && arr[i] > 0 && (arr[i]-smallestNum < len(arr)) {
			if arr[i] != arr[arr[i]-smallestNum] {
				arr[i], arr[arr[i]-smallestNum] = arr[arr[i]-smallestNum], arr[i]
			} else {
				i++
			}
		} else {
			i++
		}
	}
	var s = 1
	for s < smallestNum && len(missingNums) < k {
		missingNums = append(missingNums, s)
		s++
	}

	for i := 0; i < len(arr) && len(missingNums) < k; i++ {
		if arr[i] != i+smallestNum {
			missingNums = append(missingNums, i+smallestNum)
		}
	}

	for i := largestNum + 1; len(missingNums) < k; i++ {
		missingNums = append(missingNums, i)
	}

	return missingNums

}
