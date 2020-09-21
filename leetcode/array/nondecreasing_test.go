package array

import "testing"

func Test_findNonDecreasing(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				nums: []int{4, 2, 1},
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				nums: []int{4, 6, 2, 8},
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				nums: []int{4, 4, 2, 3},
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				nums: []int{4, 2, 4},
			},
			want: true,
		},
		{
			name: "test5",
			args: args{
				nums: []int{1},
			},
			want: true,
		},
		{
			name: "test6",
			args: args{
				nums: []int{2, 4, 6, 8},
			},
			want: true,
		},
		{
			name: "test",
			args: args{
				nums: []int{2, 4, 6, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNonDecreasing(tt.args.nums); got != tt.want {
				t.Errorf("findNonDecreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
