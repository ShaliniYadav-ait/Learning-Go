package linkedlist

import (
	"reflect"
	"testing"
)

func Test_sumOfTwoLinkedLists(t *testing.T) {
	type args struct {
		head1 *ListNode
		head2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name : "sum of two lists",
			args : args{
				head1 : createLinkedList([]int{3,1,4}),
				head2 : createLinkedList([]int{7,0,9}),
			},
			want : createLinkedList([]int{0,2,3,1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfTwoLinkedLists(tt.args.head1, tt.args.head2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfTwoLinkedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
