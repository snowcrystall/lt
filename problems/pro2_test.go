package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var list1, list2 *List
	for _, v := range []int{5, 4, 3} {
		list1 = list1.Append(&ListNode{v, nil})
	}
	for _, v := range []int{5, 3, 4, 5, 6} {
		list2 = list2.Append(&ListNode{v, nil})
	}
	l1 := list1.head
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
	l2 := list2.head
	for l2 != nil {
		fmt.Println(l2.Val)
		l2 = l2.Next
	}

	l := addTwoNumbers(list1.head, list2.head)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
