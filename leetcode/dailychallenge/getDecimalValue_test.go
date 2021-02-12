package dailychallenge

import "testing"

func Test_getDecimalValue(t *testing.T) {
	
	list1 := createLinkedList([]int{1,0,1})
	list2 := createLinkedList([]int{0})
	list3 := createLinkedList([]int{1})
	list4 := createLinkedList([]int{1,0,0,1,0,0,1,1,1,0,0,0,0,0,0})
	list5 := createLinkedList([]int{0,0})
list6 := createLinkedList([]int{1,0,1,0,0,1,1,1,0,1,1,0,0,0,0,0,0,0,0,1})

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name : "test 101",
			args : args{
				head : 	list1,
			},
			want : 5,
		},
		{
			name : "test 0",
			args : args{
				head : 	list2,
			},
			want : 0,
		},
		{
			name : "test 1",
			args : args{
				head : 	list3,
			},
			want : 1,
		},
		{
			name : "test 18880",
			args : args{
				head : 	list4,
			},
			want : 18880,
		},
		{
			name : "test 0,0",
			args : args{
				head : 	list5,
			},
			want : 0,
		},
		{
			name : "test list6",
			args : args{
				head : 	list6,
			},
			want : 685569,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
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