package main

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//进阶：你能尝试使用一趟扫描实现吗？
//双指针的经典应用，如果要删除倒数第n个节点，让fast移动n步，然后让fast和slow同时移动，直到fast指向链表末尾。删掉slow所指向的节点就可以了。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dunmmyhead := &ListNode{}
	dunmmyhead.Next = head
	cur := head
	prev := dunmmyhead //慢指针，从虚拟头节点开始走 方便删除元素
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next //用删除节点的前序节点删除节点
	return dunmmyhead.Next
}
