package mergeintervals

import (
	"dsa-patterns/helpers"
)

/*
Problem Statement #
Given two lists of intervals, find the intersection of these two lists. Each list consists of disjoint intervals sorted on their start time.

Example 1:

Input: arr1=[[1, 3], [5, 6], [7, 9]], arr2=[[2, 3], [5, 7]]
Output: [2, 3], [5, 6], [7, 7]
Explanation: The output list contains the common intervals between the two lists.
Example 2:

Input: arr1=[[1, 3], [5, 7], [9, 12]], arr2=[[5, 10]]
Output: [5, 7], [9, 10]
Explanation: The output list contains the common intervals between the two lists.
*/

func IntersectInterval(arr1, arr2 []Interval) []Interval {
	var i, j int
	var result []Interval
	var mergeInterval Interval
	for i < len(arr1) && j < len(arr2) {
		if isOverlapping(arr1[i], arr2[j]) {
			mergeInterval = Interval{Start: helpers.Max(arr1[i].Start, arr2[j].Start), End: helpers.Min(arr1[i].End, arr2[j].End)}
			result = append(result, mergeInterval)
		}
		if arr1[i].End <= arr2[j].End {
			i++
		} else {
			j++
		}

	}
	return result
}
