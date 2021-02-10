package others

import "testing"

func Test_baseConversion(t *testing.T) {
	type args struct {
		s  string
		b1 int
		b2 int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name : "binary to hexa",
			args : args{
				s : "11010101",
				b1 : 2,
				b2 : 16,
			},
			want : "D5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseConversion(tt.args.s, tt.args.b1, tt.args.b2); got != tt.want {
				t.Errorf("baseConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
