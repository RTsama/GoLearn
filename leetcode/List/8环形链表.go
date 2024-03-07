package main

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//为了表示给定链表中的环，使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
//如果 pos 是 -1，则在该链表中没有环。

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && slow != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next //slow剩余的路程就算入环的索引
				head = head.Next
			}
			return head
		}
	}
	return nil
}
