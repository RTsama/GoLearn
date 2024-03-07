package main

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: head, //使用虚拟头节点 简化操作
	}
	p := dummyHead
	for head != nil && head.Next != nil {
		p.Next = head.Next
		n := head.Next.Next
		head.Next.Next = head
		head.Next = n
		//p = list[(i+2)-1]
		p = head
		//head= list[(i+2)]
		head = n
	}
	return dummyHead.Next
}

// 递归版本
func swapPairsD(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
