package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
		carry = sum / 10
	}

	return head.Next
}

func newList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
		l = l.Next
	}
	fmt.Println()
}

func main() {
	l1 := newList([]int{2, 4, 3})
	l2 := newList([]int{5, 6, 4})

	fmt.Print("L1: ")
	printList(l1)
	fmt.Print("L2: ")
	printList(l2)

	result := addTwoNumbers(l1, l2)

	fmt.Print("Resultado: ")
	printList(result)

	fmt.Println()
	l3 := newList([]int{9, 9, 9, 9, 9, 9, 9})
	l4 := newList([]int{9, 9, 9, 9})
	fmt.Print("L3: ")
	printList(l3)
	fmt.Print("L4: ")
	printList(l4)
	fmt.Print("Resultado: ")
	printList(addTwoNumbers(l3, l4))
}
