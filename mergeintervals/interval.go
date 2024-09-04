package mergeintervals

import "sort"

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func NewIntervalList(list [][]int) []Interval {
	var intervalList []Interval
	for i := 0; i < len(list); i++ {
		intervalList = append(intervalList, NewInterval(list[i][0], list[i][1]))
	}
	return intervalList
}

func NewInterval(a, b int) Interval {
	return Interval{Start: a, End: b}
}

func isOverlapping(intervalA, intervalB Interval) bool {
	if intervalA.Start <= intervalB.End && intervalA.End >= intervalB.Start {
		return true
	}
	return false
}

func sortIntervals(intervals []Interval) []Interval {
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})
	return intervals
}

func (intervals Intervals) SortByStartTime() {
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})
}
