package slice

import "testing"

func Test_consecutiveOne(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				input: []int{1, 0, 1, 1, 1, 0},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				input: []int{1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consecutiveOne(tt.args.input); got != tt.want {
				t.Errorf("consecutiveOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
