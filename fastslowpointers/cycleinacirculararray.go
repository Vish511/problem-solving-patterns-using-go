package fastslowpointers

import (
	"fmt"
	"math"
)

/*
Cycle in a Circular Array (hard) #
We are given an array containing positive and negative numbers. Suppose the array contains a number ‘M’ at a particular index. Now, if ‘M’ is positive we will move forward ‘M’ indices and if ‘M’ is negative move backwards ‘M’ indices. You should assume that the array is circular which means two things:

If, while moving forward, we reach the end of the array, we will jump to the first element to continue the movement.
If, while moving backward, we reach the beginning of the array, we will jump to the last element to continue the movement.
Write a method to determine if the array has a cycle. The cycle should have more than one element and should follow one direction which means the cycle should not contain both forward and backward movements.

Example 1:

Input: [1, 2, -1, 2, 2]
Output: true
Explanation: The array has a cycle among indices: 0 -> 1 -> 3 -> 0
Example 2:

Input: [2, 2, -1, 2]
Output: true
Explanation: The array has a cycle among indices: 1 -> 3 -> 1
Example 3:

Input: [2, 1, -1, -2]
Output: false
Explanation: The array does not have any cycle.
*/

func CycleInACircularArray(arr []int) bool {
	var slow, fast int
	var direction int
	for {
		slow = move(arr, slow)
		fast = move(arr, move(arr, fast))

		if slow == fast {
			if arr[slow] > 0 {
				direction = 1
			}
			var curr = move(arr, slow)
			//if even after moving, we are ending up on the same element, then break because we have a one element cycle
			if curr == slow {
				return false
			}
			var currDirection int
			fmt.Println("Cycle")
			fmt.Println("slow", slow)
			for curr != slow {
				fmt.Println("curr", curr)
				if arr[curr] > 0 {
					currDirection = 1
				}
				if currDirection != direction {
					return false
				}
				curr = move(arr, curr)
			}
			return true
		}

		if fast == len(arr)-1 {
			return false
		}
	}
	return false

}

func move(arr []int, m int) int {
	var steps = int(math.Abs(float64(arr[m])))
	var position int = m
	//If the amount of steps(either positive or negative) is greater than length of the array, we mod it to get a value which is corresponding to the length of the array in order to avoid lot of operations
	if steps > len(arr) {
		steps = steps % len(arr)
	}
	//In the previous step, since we are getting only the absolute value, we have to add the sign(positive or negative) so that we can move forward or backward
	if arr[m] < 0 {
		steps = -steps
	}
	for steps != 0 {
		if steps > 0 {
			position++
			steps--
			if position == len(arr) {
				position = 0
			}
		} else if steps < 0 {
			position--
			steps++
			if position == -1 {
				position = len(arr) - 1
			}
		}
	}
	return position

}
