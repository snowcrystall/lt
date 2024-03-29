package problems

/**

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
*/

type List struct {
	head *ListNode
	tail *ListNode
}

func (list *List) Append(node *ListNode) *List {
	if list == nil {
		list = &List{node, node}
	} else {
		list.tail.Next = node
		list.tail = node
		for list.tail.Next != nil {
			list.tail = list.tail.Next
		}
	}
	return list
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p *ListNode
	var v, isCarry int
	result := l1
	for l1 != nil && l2 != nil {
		v = (l1.Val + l2.Val + isCarry)
		isCarry = v / 10
		l1.Val = v % 10
		p = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 != nil {
		p.Next = l2
		l1 = l2
	}

	if isCarry == 1 {
		for l1 != nil {
			v = (l1.Val + isCarry)
			isCarry = v / 10
			l1.Val = v % 10
			p = l1
			l1 = l1.Next
		}
	}
	if isCarry == 1 {
		p.Next = &ListNode{1, nil}
	}
	return result
}
