package linkedlist

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name : "Merge 2 Lists",
			args : args{
				lists : []*ListNode{createLinkedList([]int{1,4,7}), createLinkedList([]int{2,3,5})},
			},
			want : createLinkedList([]int{1,2,3,4,5,7}),
		},
		{
			name : "Merge 3 Lists",
			args : args{  
				lists : []*ListNode{createLinkedList([]int{1,4,5}), createLinkedList([]int{1,3,4}), createLinkedList([]int{2,6})},
			},
			want : createLinkedList([]int{1,1,2,3,4,4,5,6}),
		},
		{
			name : "Merge 2 Lists containing 1",
			args : args{  
				lists : []*ListNode{createLinkedList([]int{1,1}), createLinkedList([]int{1})},
			},
			want : createLinkedList([]int{1,1,1}),
		},
		{
			name : "Merge 2 Lists containing nil",
			args : args{ 
				lists : []*ListNode{createLinkedList([]int{1,2,3}), createLinkedList([]int{})},
			},
			want : createLinkedList([]int{1,2,3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
