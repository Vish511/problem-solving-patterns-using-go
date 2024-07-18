package singlylinkedlist

import (
	"fmt"
)

type SinglyLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

type Node struct {
	Value int
	Next  *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (sll *SinglyLinkedList) Traverse() {
	var curr = sll.Head
	for curr != nil {
		fmt.Printf("%d->", curr.Value)
		curr = curr.Next
	}
	fmt.Println("")
}

func (sll *SinglyLinkedList) Push(value int) (int, bool) {
	var newNode = NewNode(value)
	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = newNode
	} else {
		sll.Tail.Next = newNode
		sll.Tail = newNode
	}
	sll.Len++
	return sll.Len, true
}

func (sll *SinglyLinkedList) Pop() (*Node, bool) {
	if sll.Len == 0 {
		return nil, false
	}
	var poppedNode *Node
	if sll.Len == 1 {
		poppedNode = sll.Head
		sll.Head = nil
		sll.Tail = nil
	} else {
		var curr = sll.Head
		var newTail *Node
		for curr.Next != nil {
			newTail = curr
			curr = curr.Next
		}
		poppedNode = curr
		sll.Tail = newTail
		sll.Tail.Next = nil

	}
	sll.Len--
	return poppedNode, true
}

// Remove one node from the beginning
func (sll *SinglyLinkedList) Shift() (*Node, bool) {
	if sll.Len == 0 {
		return nil, false
	}
	var removedNode *Node
	var newHead = sll.Head.Next
	removedNode = sll.Head
	sll.Head = newHead
	sll.Len--
	if sll.Len == 0 {
		sll.Tail = newHead
	}
	return removedNode, true
}

func (sll *SinglyLinkedList) Unshift(value int) (int, bool) {
	if sll.Len == 0 {
		return sll.Push(value)
	}
	var newNode = NewNode(value)
	var currHead = sll.Head
	sll.Head = newNode
	sll.Head.Next = currHead
	sll.Len++
	return sll.Len, true
}

func (sll *SinglyLinkedList) Get(idx int) (*Node, bool) {
	if idx >= sll.Len || idx < 0 {
		return nil, false
	}
	var i int
	var curr = sll.Head
	for i != idx && curr.Next != nil {
		curr = curr.Next
		i++
	}
	return curr, true
}

func (sll *SinglyLinkedList) Set(idx int, value int) bool {
	if nodeAtIdx, foundNode := sll.Get(idx); foundNode {
		nodeAtIdx.Value = value
	}
	return false
}

func (sll *SinglyLinkedList) Insert(idx int, value int) (int, bool) {
	if idx < 0 || idx > sll.Len {
		return sll.Len, false
	}
	if idx == 0 {
		return sll.Unshift(value)
	}
	if idx == sll.Len {
		return sll.Push(value)
	}
	var nodeBeforeIdx, foundNode = sll.Get(idx - 1)
	if !foundNode {
		return sll.Len, false
	}
	var newNext = nodeBeforeIdx.Next
	var newNode = NewNode(value)
	newNode.Next = newNext
	nodeBeforeIdx.Next = newNode
	sll.Len++
	return sll.Len, true
}

func (sll *SinglyLinkedList) Delete(idx int) (*Node, bool) {
	if idx < 0 || idx >= sll.Len {
		return nil, false
	}
	if idx == 0 {
		return sll.Shift()
	}
	if idx == sll.Len-1 {
		return sll.Pop()
	}
	var prev, foundNode = sll.Get(idx - 1)
	if !foundNode {
		return nil, false
	}
	var nodeToBeRemoved = prev.Next
	var next = nodeToBeRemoved.Next

	prev.Next = next
	nodeToBeRemoved.Next = nil
	sll.Len--
	return nodeToBeRemoved, true
}

/*
Problem Statement #
Given the head of a Singly LinkedList, write a function to determine if the LinkedList has a cycle in it or not.
*/
func (sll *SinglyLinkedList) CycleExists() (*Node, bool) {

	var slow = sll.Head
	var fast = sll.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow, true
		}
	}
	return nil, false
}

/*
Given the head of a Singly LinkedList that contains a cycle, write a function to find the starting node of the cycle.
*/

func (sll *SinglyLinkedList) GetStartOfCycle() *Node {
	var _, cycleExists = sll.CycleExists()
	if !cycleExists {
		return nil
	}
	var temp = &Node{}
	var curr = sll.Head
	var next *Node
	for curr != nil {
		if curr.Next == temp {
			break
		}
		next = curr.Next
		curr.Next = temp
		curr = next
	}
	return curr
}

/*
Problem Statement #
Given the head of a Singly LinkedList, write a method to return the middle node of the LinkedList.

If the total number of nodes in the LinkedList is even, return the second middle node.

Example 1:
Input: 1 -> 2 -> 3 -> 4 -> 5 -> null
Output: 3

Example 2:
Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null
Output: 4

Example 3:
Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> null
Output: 4
*/
func (sll *SinglyLinkedList) GetMiddleNode() *Node {
	var slow = sll.Head
	var fast = sll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow == sll.Head {
		return nil
	}
	return slow
}

/*
Palindrome LinkedList (medium) #
Given the head of a Singly LinkedList, write a method to check if the LinkedList is a palindrome or not.

Your algorithm should use constant space and the input LinkedList should be in the original form once the algorithm is finished. The algorithm should have
𝑂(𝑁) time complexity where ‘N’ is the number of nodes in the LinkedList.

Example 1:
Input: 2 -> 4 -> 6 -> 4 -> 2 -> null
Output: true

Example 2:
Input: 2 -> 4 -> 6 -> 4 -> 2 -> 2 -> null
Output: false
*/
func (sll *SinglyLinkedList) IsPalindrome() bool {
	if sll.Len < 1 {
		return false
	}
	var isPalindrome bool
	middleNode := sll.GetMiddleNode()
	newMiddleNode := sll.Reverse(middleNode)
	sll.Traverse()
	var slow = sll.Head
	var fast = newMiddleNode
	for slow != nil && fast != nil {
		if slow.Value != fast.Value {
			isPalindrome = false
			break
		}

		slow = slow.Next
		fast = fast.Next
		fmt.Println("slow", "fast", slow, fast)

		if slow == newMiddleNode || fast == nil {
			fmt.Println("Inside")
			isPalindrome = true
			break
		}
	}
	sll.Reverse(newMiddleNode)
	return isPalindrome

}

// 1->2->3->4->5   =  5->4->3->2->1
func (sll *SinglyLinkedList) Reverse(head *Node) *Node {
	var curr = head
	var prev *Node
	var next *Node
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	if head == sll.Head {
		sll.Head, sll.Tail = sll.Tail, sll.Head
		return sll.Head
	}

	var currNode = sll.Head
	var nodePreviousToHead *Node
	for currNode != head {
		nodePreviousToHead = currNode
		currNode = currNode.Next
	}
	head, sll.Tail = sll.Tail, head
	nodePreviousToHead.Next = head
	return head

}

/*
Rearrange a LinkedList (medium) #
Given the head of a Singly LinkedList, write a method to modify the LinkedList such that the nodes from the second half of the LinkedList are inserted alternately to the nodes from the first half in reverse order. So if the LinkedList has nodes 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null, your method should return 1 -> 6 -> 2 -> 5 -> 3 -> 4 -> null.

Your algorithm should not use any extra space and the input LinkedList should be modified in-place.

Example 1:

Input: 2 -> 4 -> 6 -> 8 -> 10 -> 12 -> null
Output: 2 -> 12 -> 4 -> 10 -> 6 -> 8 -> null
Example 2:

Input: 2 -> 4 -> 6 -> 8 -> 10 -> null
Output: 2 -> 10 -> 4 -> 8 -> 6 -> null
*/

func (sll *SinglyLinkedList) Rearrange() bool {
	if sll.Len == 0 {
		return false
	}
	var headSecondHalf = sll.Head
	var tailFirstHalf *Node
	for i := 0; i < sll.Len/2; i++ {
		tailFirstHalf = headSecondHalf
		headSecondHalf = headSecondHalf.Next
	}
	var newHeadSecondHalf = sll.Reverse(headSecondHalf)
	tailFirstHalf.Next = nil
	var slow = sll.Head
	var fast = newHeadSecondHalf
	var slowNext, fastNext *Node
	for slow != nil && fast != nil {
		slowNext = slow.Next
		fastNext = fast.Next
		slow.Next = fast
		fast.Next = slowNext
		sll.Tail = fast
		fast = fastNext
		slow = slowNext
	}

	if fast != nil {
		sll.Tail.Next = fast
		sll.Tail = fast
		fast = fast.Next
	}
	sll.Traverse()

	return true

}
