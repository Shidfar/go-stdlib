package queue

import "testing"

func Test_queue_PushPop(t *testing.T) {
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
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New[int]()
			for _, v := range tt.args.vals {
				q.Push(v)
			}
			var i = 0
			for !q.IsEmpty() {
				if got := q.Pop(); got != tt.want[i] {
					t.Errorf("got = %v, want = %v", got, tt.want[i])
				}
				i++
			}
			if i != len(tt.args.vals) {
				t.Errorf("queue lens don't mathc up. got = %v, want = %v", i, len(tt.args.vals))
			}
		})
	}
}
