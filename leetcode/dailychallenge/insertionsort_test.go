package dailychallenge

import (
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	
	list1 := createLinkedList([]int{4,2,1,3})
	list2 := createLinkedList([]int{-1,5,3,4,0})
	list3 := createLinkedList([]int{1,8,4,7,9,3})
	list4 := createLinkedList([]int{})
	list5 := createLinkedList([]int{4,19,14,5,-3,1,8,5,11,15})

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name : "test list 1",
			args : args{
				head : list1,
			},
			want  : createLinkedList([]int{1,2,3,4}),
		},
		{
			name : "test list 2",
			args : args{
				head : list2,
			},
			want  : createLinkedList([]int{-1,0,3,4,5}),
		},
		{
			name : "test list 3",
			args : args{
				head : list3,
			},
			want  : createLinkedList([]int{1,3,4,7,8,9}), 
		},
		{
			name : "test empty list",
			args : args{
				head : list4,
			},
			want  : createLinkedList([]int{}), 
		},
		{
			name : "test  list5",
			args : args{
				head : list5,
			},
			want  : createLinkedList([]int{-3,1,4,5,5,8,11,14,15,19}),  //[4,19,14,5,-3,1,8,5,11,15]
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
