package mergeintervals

import (
	"dsa-patterns/helpers"
	"dsa-patterns/priorityqueue"
	"sort"
)

/*

Maximum CPU Load (hard) #
We are given a list of Jobs. Each job has a Start time, an End time, and a CPU load when it is running. Our goal is to find the maximum CPU load at any time if all the jobs are running on the same machine.

Example 1:

Jobs: [[1,4,3], [2,5,4], [7,9,6]]
Output: 7
Explanation: Since [1,4,3] and [2,5,4] overlap, their maximum CPU load (3+4=7) will be when both the
jobs are running at the same time i.e., during the time interval (2,4).
Example 2:

Jobs: [[6,7,10], [2,4,11], [8,12,15]]
Output: 15
Explanation: None of the jobs overlap, therefore we will take the maximum load of any job which is 15.
Example 3:

Jobs: [[1,4,2], [2,4,1], [3,6,5]]
Output: 8
Explanation: Maximum CPU load will be 8 as all jobs overlap during the time interval [3,4]. */

type CPULoad struct {
	Interval
	Load int
}

func newCPULoad(jobStart, jobEnd, load int) CPULoad {
	return CPULoad{Interval: NewInterval(jobStart, jobEnd), Load: load}
}

func newPriorityQueueForCPUJob() *priorityqueue.PriorityQueue[CPULoad] {
	pq := priorityqueue.PriorityQueue[CPULoad]{
		Values: []priorityqueue.PQValue[CPULoad]{},
	}
	return &pq
}

func NewCPULoadList(jobs [][]int) ([]CPULoad, bool) {
	var cpuJobList []CPULoad
	for i := 0; i < len(jobs); i++ {
		if len(jobs[i]) != 3 {
			return nil, false
		}
		cpuJobList = append(cpuJobList, newCPULoad(jobs[i][0], jobs[i][1], jobs[i][2]))

	}
	return cpuJobList, true
}

func MaxCPULoad(jobList []CPULoad) int {
	if len(jobList) == 0 {
		return 0
	}
	var maxCPULoad int
	var currCPULoad int
	sort.SliceStable(jobList, func(i, j int) bool {
		if jobList[i].Start == jobList[j].Start {
			return jobList[i].End < jobList[j].End
		}
		return jobList[i].Start < jobList[i].Start
	})
	var priorityQ = newPriorityQueueForCPUJob()
	for i := 0; i < len(jobList); i++ {
		previousJob, _, jobFound := priorityQ.Peek()
		for previousJob.End <= jobList[i].Start && jobFound {
			currCPULoad -= previousJob.Load
			_, _, _ = priorityQ.ExtractMin()
			previousJob, _, jobFound = priorityQ.Peek()
		}
		priorityQ.Insert(jobList[i], jobList[i].End)
		currCPULoad += jobList[i].Load
		maxCPULoad = helpers.Max(maxCPULoad, jobList[i].Load)
	}
	return maxCPULoad
}
