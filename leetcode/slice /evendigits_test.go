package slice

import "testing"

func Test_findEvenDigits(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				input: []int{12, 6, 4, 78, 345, 6543},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				input: []int{3},
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				input: []int{},
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				input: []int{3, 34, 12, 3455, 235344},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEvenDigits(tt.args.input); got != tt.want {
				t.Errorf("findEvenDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
