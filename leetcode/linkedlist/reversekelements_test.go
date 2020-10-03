package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "3 reversal",
			args : args{
				head : createLinkedList([]int{1,2,3,4,5,6,7,8}),
				k : 3,
			},
			want : createLinkedList([]int{3,2,1,6,5,4,7,8}),
		},
		{
			name: "2 reversal",
			args : args{
				head : createLinkedList([]int{1,2,3,4,5,6,7,8}),
				k : 2,
			},
			want : createLinkedList([]int{2,1,4,3,6,5,8,7}),
		},
		{
			name: "3 reversal",
			args : args{
				head : createLinkedList([]int{1,2,3,4,5}),
				k : 10,
			},
			want : createLinkedList([]int{1,2,3,4,5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
