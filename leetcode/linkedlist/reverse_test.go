package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Full reversal",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: createLinkedList([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "empty",
			args: args{
				head: createLinkedList([]int{}),
			},
			want: createLinkedList([]int{}),
		},
		{
			name: "single reversal",
			args: args{
				head: createLinkedList([]int{3}),
			},
			want: createLinkedList([]int{3}),
		},
		{
			name: "two number reversal",
			args: args{
				head: createLinkedList([]int{3, 4}),
			},
			want: createLinkedList([]int{4, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
