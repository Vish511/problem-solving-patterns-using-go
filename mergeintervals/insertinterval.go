package mergeintervals

import "dsa-patterns/helpers"

/*
Problem Statement #
Given a list of non-overlapping intervals sorted by their start time, insert a given interval at the correct position and merge all necessary intervals to produce a list that has only mutually exclusive intervals.

Example 1:

Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,6]
Output: [[1,3], [4,7], [8,12]]
Explanation: After insertion, since [4,6] overlaps with [5,7], we merged them into one [4,7].
Example 2:

Input: Intervals=[[1,3], [5,7], [8,12]], New Interval=[4,10]
Output: [[1,3], [4,12]]
Explanation: After insertion, since [4,10] overlaps with [5,7] & [8,12], we merged them into [4,12].
Example 3:

Input: Intervals=[[2,3],[5,7]], New Interval=[1,4]
Output: [[1,4], [5,7]]
Explanation: After insertion, since [1,4] overlaps with [2,3], we merged them into one [1,4].
*/

func InsertInterval(intervals []Interval, interval Interval) []Interval {
	var result []Interval

	var i int
	for i < len(intervals) && intervals[i].End < interval.Start {
		result = append(result, intervals[i])
		i++
	}

	for i < len(intervals) && isOverlapping(intervals[i], interval) {
		interval = Interval{Start: helpers.Min(interval.Start, intervals[i].Start), End: helpers.Max(interval.End, intervals[i].End)}
		i++
	}
	result = append(result, interval)

	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}
