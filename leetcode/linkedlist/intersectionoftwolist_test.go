package linkedlist

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {

	list1 := createLinkedList([]int{1,2,3,4,5})
	list2 := createLinkedList([]int{6,7,8})
	list2.Next.Next.Next = list1.Next

	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test intersection",
			args: args{
				headA: list1,
				headB: list2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}
