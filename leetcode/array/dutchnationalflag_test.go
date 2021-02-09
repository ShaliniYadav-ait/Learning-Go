package array

import "testing"

func Test_dutchNationalFlag(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 5 colors",
			args: args{
				nums: []int{1, 4, 3, 7, 0, 1},
				n:    2,
			},
			want: []int{1, 1, 0, 3, 7, 4},
		},
		{
			name: "test no colors",
			args: args{
				nums: []int{},
				n :  0,
			},
			want: []int{},
		},
		{
			name: "test 3 colors",
			args: args{
				nums: []int{0,1,2,0,0,1,1,2,1},
				n :  1,
			},
			want: []int{0,0,0,1,1,1,1,2,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dutchNationalFlag(tt.args.nums, tt.args.n)
			for i := 0; i < len(tt.args.nums); i++ {
				if tt.args.nums[i] == tt.want[i] {
					continue
				}
				t.Errorf("failed got %v want %v ", tt.args.nums, tt.want)
				break
			}
		})
	}
}
