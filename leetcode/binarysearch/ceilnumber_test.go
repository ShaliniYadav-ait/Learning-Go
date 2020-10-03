package binarysearch

import "testing"

func Test_findCeil(t *testing.T) {
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
			name: "Ceil not present",
			args: args{
				vals: []int{2, 3, 4, 5},
				key:  6,
			},
			want: -1,
		},
		{
			name: "ceil present in first half",
			args: args{
				vals: []int{1, 3, 4, 5},
				key:  2,
			},
			want: 1,
		},	
		{
			name: "Ceil present in second half",
			args: args{
				vals: []int{2, 3, 4, 6},
				key:  5,
			},
			want: 3,
		},
		{
			name: "last element in 2 size",
			args: args{
				vals: []int{1,3},
				key:  2,
			},
			want: 1,
		},
		{
			name: "first element in 2 size",
			args: args{
				vals: []int{1,2},
				key:  1,
			},
			want: 0,
		},
		{
			name: "key not present",
			args: args{
				vals: []int{1,3,8,10,15},
				key:  12,
			},
			want: 4,
		},
		{
			name: "key smaller than smallest",
			args: args{
				vals: []int{4,6,10},
				key:  1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCeil(tt.args.vals, tt.args.key); got != tt.want {
				t.Errorf("findCeil() = %v, want %v", got, tt.want)
			}
		})
	}
}
