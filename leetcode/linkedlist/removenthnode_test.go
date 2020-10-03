package linkedlist

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "delete 4th node",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8}),
				n:    4,
			},
			want: createLinkedList([]int{1, 2, 3, 4, 6, 7, 8}),
		},
		{
			name: "delete from single element list",
			args: args{
				head: createLinkedList([]int{3}),
				n:    1,
			},
			want: createLinkedList([]int{}),
		},
		{
			name: "delete first element from last",
			args: args{
				head: createLinkedList([]int{1, 2}),
				n:    1,
			},
			want: createLinkedList([]int{1}),
		},
		{
			name: "delete first element",
			args: args{
				head: createLinkedList([]int{1, 2, 3}),
				n:    3,
			},
			want: createLinkedList([]int{2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
