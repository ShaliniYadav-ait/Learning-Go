package list

import "testing"

func Test_numInArray(t *testing.T) {
	type args struct {
		list []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name : "is present",
			args : args{
				list : []int{1,2,3,4,5},
				n : 3,
			},
			want : true,
		},
		{
			name : "is not present",
			args : args{
				list : []int{1,2,3,4,5},
				n : 7,
			},
			want : false,
		},
		{
			name : "empty list",
			args : args{
				list : []int{},
				n : 4,
			},
			want : false,
		},
		{
			name : "negative int present",
			args : args{
				list : []int{1,2,3,-4,5},
				n : -4,
			},
			want : true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numInArray(tt.args.list, tt.args.n); got != tt.want {
				t.Errorf("numInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
