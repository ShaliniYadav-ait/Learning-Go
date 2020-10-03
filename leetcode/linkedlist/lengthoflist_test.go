package linkedlist

import "testing"

func Test_lengthOfList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantcycle bool
	}{
		{
			name: "test length of list",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want:      5,
			wantcycle: false,
		},
		{
			name: "test length of cyclic list",
			args: args{
				head: createCyclicLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3),
			},
			want:      8,
			wantcycle: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, gotcycle := lengthOfList(tt.args.head); got != tt.want || gotcycle != tt.wantcycle {
				t.Errorf("lengthOfList() = %v, want %v , %v, wantcycle %v", got, tt.want, gotcycle, tt.wantcycle)
			}
		})
	}
}

func createCyclicLinkedList(vals []int, loopIndex int) *ListNode {

	dummyHead := &ListNode{}
	curr := dummyHead
	var loopNode *ListNode

	for i := 0; i < len(vals); i++ {
		node := &ListNode{
			Val: vals[i],
		}

		if i == loopIndex {
			loopNode = node
		}

		curr.Next = node
		curr = curr.Next
	}

	curr.Next = loopNode

	return dummyHead.Next
}
