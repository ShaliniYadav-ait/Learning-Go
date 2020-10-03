package linkedlist

import (
	
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5,6,7,8,9}),
				m:    2,
				n:    8,
			},
			want: createLinkedList([]int{1, 8, 7, 6, 5,4,3,2,9}),
		},{
			name: "test",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
				m:    2,
				n:    4,
			},
			want: createLinkedList([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "total reverse",
			args: args{
				head: createLinkedList([]int{1, 2, 3}),
				m:    1,
				n:    3,
			},
			want: createLinkedList([]int{3, 2, 1}),
		},
		{
			name: "two elements",
			args: args{
				head: createLinkedList([]int{1, 2}),
				m:    2,
				n:    2,
			},
			want: createLinkedList([]int{1, 2}),
		},
		{
			name: "reverse from begining to middle",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
				m:    1,
				n:    3,
			},
			want: createLinkedList([]int{3, 2, 1, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !equalList(tt.want, got) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createLinkedList(list []int) *ListNode {
	var head *ListNode
	var prev *ListNode
	for i := 0; i < len(list); i++ {
		if head == nil {
			head = &ListNode{
				Val: list[i],
			}
			prev = head
			continue
		}

		curr := &ListNode{
			Val: list[i],
		}

		prev.Next = curr
		prev = curr
	}

	return head
}

func equalList(head1 *ListNode, head2 *ListNode) bool {
	for head1 != nil {
		if head1.Val != head2.Val {
			return false
		}

		head1 = head1.Next
		head2 = head2.Next
	}

	return true
}

/*func printList(head *ListNode) string {
	res := ""
	for head != nil {
		res += fmt.Sprintf("%d ", head.Val)
		head = head.Next
	}

	return res
}
*/

/* func Test_getAt(t *testing.T) {
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
			name: "return 2nd element",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: createLinkedList([]int{2, 3, 4, 5}),
		},
		{
			name: "return 1st element",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
				n:    1,
			},
			want: createLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "return last element",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
				n:    5,
			},
			want: createLinkedList([]int{5}),
		},
	}
	/* for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAt(tt.args.head, tt.args.n); !equalList(tt.want, got) {
				t.Errorf("getAt() = %v, want %v", printList(got), printList(tt.want))
			}
		})
	} */