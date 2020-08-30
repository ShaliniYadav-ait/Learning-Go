package assignment

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name : "all aplphabet",
			args : args{
				input : "Shalini",
			},
			want : "inilahS",
		},
		{
			name : "aplphabet with special symbols",
			args : args{
				input : "Shalini&Abhishek",
			},
			want : "kehsihbA&inilahS",
		},
		{
			name : "empty string",
			args : args{
				input : "",
			},
			want : "",
		},
		{
			name : "nihongo",
			args : args{
				input : "日本語",
			},
			want : "語本日",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.input); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
