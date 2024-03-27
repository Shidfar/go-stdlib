package stack

import "testing"

func Test_stack_Push(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"sample 1",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			[]int{8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New[int]()
			for _, v := range tt.args.vals {
				s.Push(v)
			}
			var i = 0
			for !s.IsEmpty() {
				if got := s.Pop(); got != tt.want[i] {
					t.Errorf("got = %v, want = %v", got, tt.want[i])
				}
				i++
			}
			if i != len(tt.args.vals) {
				t.Errorf("stack lens don't mathc up. got = %v, want = %v", i, len(tt.args.vals))
			}
		})
	}
}
