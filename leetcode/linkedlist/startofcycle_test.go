package linkedlist

import (
	"reflect"
	"testing"
)

func Test_startOfCycle(t *testing.T) {
	list1 := createLinkedList([]int{1, 2, 3, 4, 5})
	list2 := createLinkedList([]int{6, 7, 8})
	list2.Next.Next.Next = list1.Next

	listcycle1 := createCyclicLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3)
	listcycle2 := createLinkedList([]int{9, 10})
	listcycle2.Next.Next = listcycle1.Next.Next.Next

	type args struct {
		vals []*ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "check merged list without cycle",
			args: args{
				vals: []*ListNode{list1, list2},
			},
			want: 2,
		},
		{
			name: "check start of cycle",
			args: args{
				vals: []*ListNode{listcycle1, listcycle2},
			},
			want: 4,
		},
		/* {
			name: "check start of cycle diff length",
			args: args{
				vals: []*ListNode{createCyclicLinkedList([]int{ 3, 4, 5, 6, 7, 8}, 1), createCyclicLinkedList([]int{9, 10, 11, 4, 5, 6, 7, 8}, 3)},
			},
			want: createCyclicLinkedList([]int{4, 5, 6, 7, 8}, 0),
		}, */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startOfCycle(tt.args.vals); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("startOfCycle() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}
