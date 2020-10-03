package linkedlist

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "merge sorted",
			args: args{
				l1: createLinkedList([]int{1, 4, 6}),
				l2: createLinkedList([]int{2, 3, 5}),
			},
			want: createLinkedList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name: "merge only one",
			args: args{
				l1: createLinkedList([]int{1, 4, 6}),
				l2: createLinkedList([]int{}),
			},
			want: createLinkedList([]int{1, 4, 6}),
		},
		{
			name: "merge one after another",
			args: args{
				l1: createLinkedList([]int{1, 4, 6}),
				l2: createLinkedList([]int{7, 8, 9}),
			},
			want: createLinkedList([]int{1, 4, 6, 7, 8, 9}),
		},
		{
			name: "merge single element",
			args: args{
				l1: createLinkedList([]int{1, 4, 6}),
				l2: createLinkedList([]int{2}),
			},
			want: createLinkedList([]int{1, 2, 4, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
