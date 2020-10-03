package linkedlist

import (
	"reflect"
	"testing"
)

func Test_nextLargerNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "3 element",
			args: args{
				head: createLinkedList([]int{2, 1, 5}),
			},
			want: []int{5, 5, 0},
		},
		{
			name: "increasing",
			args: args{
				head: createLinkedList([]int{1,3,5,6,7}),
			},
			want: []int{3,5,6,7, 0},
		},
		{
			name: "decreasing",
			args: args{
				head: createLinkedList([]int{9,8,7,6,5}),
			},
			want: []int{0,0,0,0,0},
		},
		{
			name: "single element",
			args: args{
				head: createLinkedList([]int{9}),
			},
			want: []int{0},
		},
		{
			name: "empty",
			args: args{
				head: createLinkedList([]int{}),
			},
			want: []int{},
		},
		{
			name: "mixed",
			args: args{
				head: createLinkedList([]int{1,7,5,1,9,2,5,1}),
			},
			want: []int{7,9,9,9,0,5,0,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextLargerNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
