package mergeintervals

import (
	"dsa-patterns/helpers"
	"dsa-patterns/priorityqueue"
)

/*
Employee Free Time (hard) #
For ‘K’ employees, we are given a list of intervals representing the working hours of each employee. Our goal is to find out if there is a free interval that is common to all employees. You can assume that each list of employee working hours is sorted on the start time.

Example 1:
Input: Employee Working Hours=[[[1,3], [5,6]], [[2,3], [6,8]]]
Output: [3,5]
Explanation: Both the employess are free between [3,5].

Example 2:
Input: Employee Working Hours=[[[1,3], [9,12]], [[2,4]], [[6,8]]]
Output: [4,6], [8,9]
Explanation: All employess are free between [4,6] and [8,9].

Example 3:
Input: Employee Working Hours=[[[1,3]], [[2,4]], [[3,5], [7,9]]]
Output: [5,7]
Explanation: All employess are free between [5,7].
*/

func EmployeeFreeTime(empWorkingHrs []Intervals) Intervals {
	var allEmpWorkingHrs Intervals
	var empFreeTime Intervals
	for e := 0; e < len(empWorkingHrs); e++ {
		allEmpWorkingHrs = append(allEmpWorkingHrs, empWorkingHrs[e]...)
	}
	allEmpWorkingHrs.SortByStartTime()
	var interval = allEmpWorkingHrs[0]
	for i := 1; i < len(allEmpWorkingHrs); i++ {
		if !isOverlappingAppointment(interval, allEmpWorkingHrs[i]) {
			empFreeTime = append(empFreeTime, Interval{Start: interval.End, End: allEmpWorkingHrs[i].Start})
			interval = allEmpWorkingHrs[i]
		} else {
			interval = Interval{Start: helpers.Min(interval.Start, allEmpWorkingHrs[i].Start), End: helpers.Max(interval.End, allEmpWorkingHrs[i].End)}
		}
	}
	return empFreeTime
}

type EmployeeWorkingHrs struct {
	EmpID int
	Interval
	IntervalIdx int
}

func newPriorityQueueForEmp() priorityqueue.PriorityQueue[EmployeeWorkingHrs] {
	return priorityqueue.PriorityQueue[EmployeeWorkingHrs]{}
}

// The below approach uses a Priority Queue
func EmployeeFreeTimeAlternate(empWorkingHrs []Intervals) Intervals {
	var empFreeTime Intervals
	var priorityQ = newPriorityQueueForEmp()
	for i := 0; i < len(empWorkingHrs); i++ {
		priorityQ.Insert(EmployeeWorkingHrs{EmpID: i, Interval: empWorkingHrs[i][0], IntervalIdx: 0}, empWorkingHrs[i][0].Start)
	}

	prevEmpHr, _, _ := priorityQ.Peek()
	var prevInterval = prevEmpHr.Interval
	for len(priorityQ.Values) > 0 {
		currEmpHr, _, _ := priorityQ.ExtractMin()
		if !isOverlappingAppointment(prevInterval, currEmpHr.Interval) {
			empFreeTime = append(empFreeTime, Interval{Start: prevInterval.End, End: currEmpHr.Start})
			prevInterval = currEmpHr.Interval
		} else {
			if prevInterval.End < currEmpHr.End {
				prevInterval = currEmpHr.Interval
			}
		}
		var empSchedule = empWorkingHrs[currEmpHr.EmpID]
		if len(empSchedule) > currEmpHr.IntervalIdx+1 {
			empIntervalIdx := currEmpHr.IntervalIdx + 1
			priorityQ.Insert(EmployeeWorkingHrs{EmpID: currEmpHr.EmpID, Interval: empWorkingHrs[currEmpHr.EmpID][empIntervalIdx], IntervalIdx: currEmpHr.IntervalIdx + 1}, empWorkingHrs[currEmpHr.EmpID][empIntervalIdx].Start)
		}
	}
	return empFreeTime
}
