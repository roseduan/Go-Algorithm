package leetcode

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表排序

// 第一种方式，取出链表数据，排序后重新构造链表
func sortList1(head *ListNode) *ListNode {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}

	sort.Ints(data)
	dummy := &ListNode{}
	p := dummy
	for _, v := range data {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return dummy.Next
}
