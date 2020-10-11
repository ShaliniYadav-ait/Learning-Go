package queue

import "testing"

type Iqueue interface {
	Push(val int)
	Pop() (int, error)
	Length()int
}

func TestQueue_Push(t *testing.T) {
	type fields struct {
		vals []int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Push increases length",
			fields: fields{
				vals: []int{2, 3},
			},
			args: args{
				val: 1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var q Iqueue
			q = initialize()
			/* &Queue{
				vals: tt.fields.vals,
			} */
			q.Push(tt.args.val)
			if q.Length() != tt.want {
				t.Errorf("Push function = %v, want %v", q.Length(), tt.want)
			}

			got, err := q.Pop()
			if err != nil {
				t.Errorf("Pop function = %v, want  nil", err)
			}
			if got != tt.fields.vals[0] {
				t.Errorf("Pop function = %v, want %v", got, tt.fields.vals[0])
			}

			if q.Length() != tt.want-1 {
				t.Errorf("Push function = %v, want %v", q.Length(), tt.want-1)
			}
		})
	}
}
