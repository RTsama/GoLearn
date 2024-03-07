package main

//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
//如果两个链表没有交点，返回 null  整个链式结构中不存在环。。 交点不是数值相等，而是指针相等。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	//遍历两个链表 遇到i相同跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// 双指针法
func getIntersectionNodeP(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB //
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}

	}
	return l1
}
