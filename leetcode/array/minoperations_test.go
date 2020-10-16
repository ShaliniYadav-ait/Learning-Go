package array

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name : "test for 6",
			args: args{
				n : 6,
			},
			want : 9,
		},
		{
			name : "test for 3",
			args: args{
				n : 3,
			},
			want : 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.n); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
