package main

import (
	. "github.com/yacen/sword2offer/util"
)

// 输入一个链表，按链表值从尾到头的顺序返回一个ArrayList。

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	PrintListFromTailToHead(head)
}
