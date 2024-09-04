package mergeintervals

import "sort"

/*
Problem Statement #
Given an array of intervals representing ‘N’ appointments, find out if a person can attend all the appointments.

Example 1:

Appointments: [[1,4], [2,5], [7,9]]
Output: false
Explanation: Since [1,4] and [2,5] overlap, a person cannot attend both of these appointments.
Example 2:

Appointments: [[6,7], [2,4], [8,12]]
Output: true
Explanation: None of the appointments overlap, therefore a person can attend all of them.
Example 3:

Appointments: [[4,5], [2,3], [3,6]]
Output: false
Explanation: Since [4,5] and [3,6] overlap, a person cannot attend both of these appointments.
*/

func CanAttendAllAppointments(appointments []Interval) bool {
	sort.SliceStable(appointments, func(i, j int) bool {
		return appointments[i].Start < appointments[j].Start
	})
	if len(appointments) == 0 {
		return false
	}

	for i := 0; i < len(appointments)-1; i++ {
		if isOverlappingAppointment(appointments[i], appointments[i+1]) {
			return false
		}
	}
	return true
}

func isOverlappingAppointment(apptmtA, apptmtB Interval) bool {
	if apptmtA.End > apptmtB.Start && apptmtA.Start < apptmtB.End {
		return true
	}
	return false
}
