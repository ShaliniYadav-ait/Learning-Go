package slice

import (
	"reflect"
	"testing"
)

func Test_runningSumOfArray(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum []int
	}{
		{
			name: "test",
			args: args{
				input: []int{1, 2, 3, 4, 5, 6},
			},
			wantSum: []int{1, 3, 6, 10, 15, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := runningSumOfArray(tt.args.input); !reflect.DeepEqual(gotSum, tt.wantSum) {
				t.Errorf("runningSumOfArray() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
/* 
func Test_simple(t *testing.T) {
fmt.Println("Chala")
	t.Run("test1", func(t *testing.T) {
		result := runningSumOfArray([]int{1, 2, 3, 4, 5, 6,7})
		if !reflect.DeepEqual(result, []int{1, 3, 6, 10, 15, 21,29}) {
			fmt.Println("Shalini")
			t.Fail()
			
		}
	})
} */
