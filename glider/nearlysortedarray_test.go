package glider

import "testing"

func Test_findElementInNearlySortedArray(t *testing.T) {
	type args struct {
		vals []int
		key  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name : "nearly sorted",
			args : args {
				vals : []int{2,3,5,4,7},
				key : 4,
			},
			want : 3,
		},

		{
			name : "nearly sorted",
			args : args {
				vals : []int{9,3,15,29,30,31,32},
				key : 15,
			},
			want : 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findElementInNearlySortedArray(tt.args.vals, tt.args.key); got != tt.want {
				t.Errorf("findElementInNearlySortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
