package linkedlist

import "testing"

func Test_checkCycleInList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkCycleInList(tt.args.head); got != tt.want {
				t.Errorf("checkCycleInList() = %v, want %v", got, tt.want)
			}
		})
	}
}
