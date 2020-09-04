package assignment

import "testing"

func Test_findGCD(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "GCD 30 & 15",
			args: args{
				num1: 30,
				num2: 15,
			},
			want: 15,
		},
		{
			name: "GCD 15 & 30",
			args: args{
				num1: 15,
				num2: 30,
			},
			want: 15,
		},
		{
			name: "GCD 25 & 5",
			args: args{
				num1: 25,
				num2: 5,
			},
			want: 5,
		},
		{
			name: "GCD 30 & 9",
			args: args{
				num1: 30,
				num2: 9,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGCD(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("findGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
