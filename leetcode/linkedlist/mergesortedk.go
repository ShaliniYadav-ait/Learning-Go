package linkedlist

import (
	"math"
)

// https://leetcode.com/problems/merge-k-sorted-lists/

func mergeKLists(lists []*ListNode) *ListNode {

	dummyHead := &ListNode{}
	curr := dummyHead
	listMap := map[int]int{}
	min, max := math.MaxInt32, math.MinInt32

	for i := 0; i < len(lists); i++ {
		for lists[i] != nil {
			if min > lists[i].Val {
				min = lists[i].Val
			}

			if max < lists[i].Val {
				max = lists[i].Val
			}

			listMap[lists[i].Val]++
			lists[i] = lists[i].Next
		}
	}

	for i := min; i <= max; i++ {
		freq := listMap[i]
		for j := 0; j < freq; j++ {
			node := &ListNode{Val: i}
			curr.Next = node
			curr = curr.Next
		}
	}
	return dummyHead.Next
}