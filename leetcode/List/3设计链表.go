package main

import "fmt"

//在链表类中实现这些功能：
//	get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//	addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//	addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//	addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
//	deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

// new 返回的是pointer    make返回的是本身
type SingleNode struct {
	Val  int
	Next *SingleNode
}

type MyLinkedList struct {
	dummyHead *SingleNode //虚拟头节点
	Size      int         // 链表大小
}

// make 返回普通对象  new返回指针 并赋零值
func initMyList() MyLinkedList {
	newNode := &SingleNode{
		-999,
		nil,
	}
	return MyLinkedList{
		dummyHead: newNode,
		Size:      0,
	}
}

//	get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。

func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size { //索引无效
		return -1
	}
	// cur从真正的头节点开始走
	cur := this.dummyHead.Next   //设置当前节点为真实头节点
	for i := 0; i < index; i++ { //遍历到索引
		cur = cur.Next
	}
	return cur.Val //返回节点值
}

// addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &SingleNode{Val: val}
	newNode.Next = this.dummyHead.Next
	this.dummyHead.Next = newNode
	this.Size++
}

// addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) addAtTail(val int) {
	newNode := &SingleNode{Val: val}
	cur := this.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	this.Size++
}

//	addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。
//	如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。
//	如果index小于0，则在头部插入节点。

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > this.Size { //索引大于索引长度直接返回
		return
	}
	newNode := &SingleNode{Val: val}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	this.Size++
}

//	deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

func (this *MyLinkedList) deleteAtIndex(index int) {
	if index < 0 || index >= this.Size { //索引无效
		return
	}
	cur := this.dummyHead //cur.Next是要被删除的节点
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	this.Size--
}

//打印链表

func (this *MyLinkedList) PrintList() {
	cur := this.dummyHead
	for cur.Next != nil {
		fmt.Println(cur.Next.Val)
		cur = cur.Next
	}
}

func main() {
	initMyList()
}
