package assignment

import "testing"

func Test_findFibonacci(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fibonacci 6",
			args: args{
				num: 6,
			},
			want: 8,
		},
		{
			name: "fibonacci 7",
			args: args{
				num: 7,
			},
			want: 13,
		},
		{
			name: "fibonacci 14",
			args: args{
				num: 14,
			},
			want: 377,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFibonacci(tt.args.num); got != tt.want {
				t.Errorf("findFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
