package main

import (
	binarytree "dsa-patterns/binarysearchtree"
	"dsa-patterns/cyclicsort"
	"dsa-patterns/fastslowpointers"
	"dsa-patterns/mergeintervals"
	"dsa-patterns/singlylinkedlist"
	"fmt"
)

func main() {
	/* Sliding Window */
	//Avg of Subarray of size k
	// fmt.Println(slidingwindow.AvgOfSubArrayOfSizeK([]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5))
	//Max Sum of Subarray of size k
	// fmt.Println(slidingwindow.MaxSumOfSubArray([]int{2, 3, 4, 1, 5}, 2))
	//Smallest Sub array with given sum
	// fmt.Println(slidingwindow.SmallestSubArrayWithGivenSum([]int{3, 4, 1, 1, 6}, 8))
	//Longest Sub string with k distinct characters
	// fmt.Println(slidingwindow.LongestSubStringWithKDistinctChars("araaci", 1))
	//Fruits Into Baskets
	// fmt.Println(slidingwindow.FruitsIntoBaskets([]string{"A", "B", "C", "B", "B", "C"}))
	//No Repeat substring
	// fmt.Println(slidingwindow.NoRepeatSubstring("aabccbb"))
	//Longest substring with same letters after replacement
	// fmt.Println(slidingwindow.LongestSubStrWithSameLettersAfterReplacement("aabcebb", 2))
	//Longest Substring with ones after replacement
	// fmt.Println(slidingwindow.LongestSubStrWithOnesAfterReplacement([]int{0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0}, 2))
	//Permutations in a string
	// fmt.Println(slidingwindow.PermutationsInAString("bcdxabcdy", "bcdyabcdx"))
	//String Anagrams
	// fmt.Println(slidingwindow.StringAnagrams("abbcabc", "abc"))
	//Smallest Window containing substring
	// fmt.Println(slidingwindow.SmallestWindowContainingSubstr("abdabexcaabftgbch", "abc"))
	//Words Concatenation
	// fmt.Println(slidingwindow.WordsConcatenation("catfoxcat", []string{"cat", "fox"}))

	/* Pattern Two Pointers */
	//Pair with target sum
	// fmt.Println(twopointers.PairWithTargetSum([]int{1, 2, 3, 4, 6}, 6))

	//Remove duplicates
	// fmt.Println(twopointers.RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9}))

	//Squaring a sorted array
	// fmt.Println(twopointers.SquaringASortedArray([]int{-2, -1, 0, 0, 1, 2}))

	//Find Triplets with given sum
	// fmt.Println(twopointers.TripletWithGivenSum([]int{-3, 0, 1, 2, -1, 1, -2}, 0))

	//Triplet sum close to target
	// fmt.Println(twopointers.TripletSumCloseToTarget([]int{1, 0, 1, 1}, 100))

	//Triplet sum lesser than target
	// fmt.Println(twopointers.TripletsWithSmallerSumThanTarget([]int{-1, 4, 2, 1, 3}, 5))

	//Dutch National Flag
	// fmt.Println(twopointers.DutchNationalFlag([]int{2, 2, 0, 1, 2, 0}))

	//Quadruple sum to target
	// fmt.Println(twopointers.QuadrupleSumToTarget([]int{2, 0, -1, 1, -2, 2}, 2))

	//Compare strings with backspaces
	// fmt.Println(twopointers.CompareStringsContainingBackspaces("xp#", "a#xyz##"))

	//Minimum window sort
	// fmt.Println(twopointers.MinWindowSort([]int{1, 3, 2, 0, -1, 7, 10}))

	/* Singly Linked List */
	sll := singlylinkedlist.NewSinglyLinkedList()

	for i := 1; i <= 101; i++ {
		sll.Push(i)
	}

	fmt.Println(sll.IsPalindrome())
	sll.Rearrange()

	fmt.Println("cycle", fastslowpointers.CycleInACircularArray([]int{1, 1}))
	// var node, _ = sll.Get(4)
	// // sll.Tail.Next = node
	// fmt.Println(sll.CycleExists())
	// fmt.Println(sll.GetMiddleNode())
	// sll.Reverse(sll.GetMiddleNode())
	// sll.Traverse()
	//Happy number
	// fmt.Println(fastslowpointers.IsHappyNumber(29))

	/* Merge Intervals */

	// intervals := mergeintervals.NewIntervalList([][]int{{1, 4}, {2, 5}, {7, 9}})
	// fmt.Println(mergeintervals.MergeOverlappingIntervals(intervals))

	// intervals := mergeintervals.NewIntervalList([][]int{{2, 3}, {5, 7}})
	// interval := mergeintervals.NewInterval(1, 4)
	// fmt.Println(mergeintervals.InsertInterval(intervals, interval))

	// interval1 := mergeintervals.NewIntervalList([][]int{{1, 3}, {5, 7}, {9, 11}, {13, 19}})
	// interval2 := mergeintervals.NewIntervalList([][]int{{5, 10}, {12, 14}, {16, 20}})

	// fmt.Println(mergeintervals.IntersectInterval(interval1, interval2))

	// appointments := mergeintervals.NewIntervalList([][]int{{4, 5}, {2, 3}, {3, 6}})
	// fmt.Println("Can attend all appointments ", mergeintervals.CanAttendAllAppointments(appointments))

	meetings := mergeintervals.NewIntervalList([][]int{{6, 7}, {2, 4}, {8, 12}})
	fmt.Println(mergeintervals.MinMeetingRooms(meetings))

	cpuJobsList, _ := mergeintervals.NewCPULoadList([][]int{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}})
	mergeintervals.MaxCPULoad(cpuJobsList)

	empWorkingHrs := []mergeintervals.Intervals{mergeintervals.NewIntervalList([][]int{{1, 3}}), mergeintervals.NewIntervalList([][]int{{2, 4}}), mergeintervals.NewIntervalList([][]int{{3, 5}, {7, 9}})}
	fmt.Println(mergeintervals.EmployeeFreeTime(empWorkingHrs))
	fmt.Println("alternate", mergeintervals.EmployeeFreeTimeAlternate(empWorkingHrs))

	/* Pattern Cyclic Sort */
	fmt.Println("Cyclic Sort")
	// fmt.Println(cyclicsort.CyclicSort([]int{10, 2, 9, 8, 3, 7, 5, 6, 4, 1}))

	// fmt.Println(cyclicsort.FindTheMissingNum([]int{4, 0, 3, 1, 5}))
	fmt.Println(cyclicsort.FindAllMissingNums([]int{2, 3, 1, 8, 2, 3, 5, 1}))
	fmt.Println(cyclicsort.FindDuplicateNum([]int{2, 1, 3, 3, 5, 4}))
	fmt.Println(cyclicsort.FindAllDuplicateNums([]int{3, 6, 5, 6, 2, 1, 2}))
	fmt.Println(cyclicsort.FindTheCorruptPair([]int{3, 1, 2, 5, 2}))
	fmt.Println(cyclicsort.FindSmallestMissingPosNum([]int{-3, 1, 5, 4, 2}))
	fmt.Println("Find first k positive missing numbers", cyclicsort.FindFirstKPosMissingNums([]int{12, 10, -2, -5, -7, -9, 0, 13}, 14))

	//Binary Search Tree
	var bst = binarytree.NewBinarySearchTree()
	bst.Insert(10)
	bst.Insert(8)
	bst.Insert(12)
	bst.Insert(3)
	bst.Insert(9)
	bst.Insert(11)
	bst.Insert(13)
	bst.Insert(7)
	fmt.Println(bst.Find(7))

}
