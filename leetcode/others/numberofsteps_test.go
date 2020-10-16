package others

import "testing"

func Test_numberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test numer 14",
			args: args{
				num: 14,
			},
			want: 6,
		},
		{
			name: "test numer 8",
			args: args{
				num: 8,
			},
			want: 4,
		},
		{
			name: "test numer 123",
			args: args{
				num: 123,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.args.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
