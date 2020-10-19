package stack

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
 		{
			name : "test for path /a/./b/../../c/",
			args : args{
				path : "/a/./b/../../c/",
			},
			want : "/c",
		},
		{
			name : "test for path /a/./b/../c/",
			args : args{
				path : "/a/./b/../c/",
			},
			want : "/a/c",
		}, 
		{
			name : "test for path /a/../../b/../c//.//",
			args : args{
				path : "/a/../../b/../c//.//",
			},
			want : "/c",
		}, 
		{
			name : "test for path /./../a",
			args : args{
				path : "/./../a",
			},
			want : "/a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
