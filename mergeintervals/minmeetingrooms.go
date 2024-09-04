package mergeintervals

import (
	"dsa-patterns/helpers"
	"dsa-patterns/priorityqueue"
)

/*
Minimum Meeting Rooms (hard) #
Given a list of intervals representing the start and end time of ‘N’ meetings, find the minimum number of rooms required to hold all the meetings.

Example 1:
Meetings: [[1,4], [2,5], [7,9]]
Output: 2
Explanation: Since [1,4] and [2,5] overlap, we need two rooms to hold these two meetings. [7,9] can
occur in any of the two rooms later.

Example 2:
Meetings: [[6,7], [2,4], [8,12]]
Output: 1
Explanation: None of the meetings overlap, therefore we only need one room to hold all meetings.

Example 3:
Meetings: [[1,4], [2,3], [3,6]]
Output:2
Explanation: Since [1,4] overlaps with the other two meetings [2,3] and [3,6], we need two rooms to
hold all the meetings.

Example 4:
Meetings: [[4,5], [2,3], [2,4], [3,5]]
Output: 2
Explanation: We will need one room for [2,3] and [3,5], and another room for [2,4] and [4,5].
*/

func NewPriorityQueueInterval() *priorityqueue.PriorityQueue[Interval] {
	pq := priorityqueue.PriorityQueue[Interval]{
		Values: []priorityqueue.PQValue[Interval]{},
	}
	return &pq
}

func MinMeetingRooms(meetings Intervals) int {
	if len(meetings) == 0 {
		return 0
	}
	meetings.SortByStartTime()
	var roomCnt int
	var priorityQ = NewPriorityQueueInterval()
	for i := 0; i < len(meetings); i++ {
		prevMeeting, _, meetingFound := priorityQ.Peek()
		for prevMeeting.End <= meetings[i].Start && meetingFound {
			_, _, _ = priorityQ.ExtractMin()
			prevMeeting, _, meetingFound = priorityQ.Peek()
		}
		priorityQ.Insert(meetings[i], meetings[i].End)
		roomCnt = helpers.Max(roomCnt, len(priorityQ.Values))
	}
	return roomCnt
}

func MinMeetingRoomsAlternate(meetings Intervals) int {
	meetings.SortByStartTime()
	if len(meetings) < 1 {
		return 0
	}
	var accommodated = []Interval{meetings[0]}
	for i := 1; i < len(meetings); i++ {
		isAllocated := false
		for j := 0; j < len(accommodated) && !isAllocated; j++ {
			if !isOverlappingAppointment(meetings[i], accommodated[j]) {
				accommodated[j] = meetings[i]
				isAllocated = true
			}
		}
		if !isAllocated {
			accommodated = append(accommodated, meetings[i])
		}
	}
	return len(accommodated)

}
