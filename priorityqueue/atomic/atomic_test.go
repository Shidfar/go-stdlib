package atomic

import (
	"reflect"
	"testing"
	"time"
)

func TestAtomicPriorityQueue(t *testing.T) {
	type args[T any] struct {
		cap    int
		less   LessFun[T]
		inputs []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	tests := []testCase[int]{
		{
			"initialTest",
			args[int]{
				cap: 10,
				less: func(a int, b int) bool {
					return a < b
				},
				inputs: []int{10, 9, 2, 7, 8, 1, 6, 5, 3, 4},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			pq := NewAtomicPriorityQueue(tt.args.cap, tt.args.less)

			// act
			for _, it := range tt.args.inputs {
				pq.Push(it)
			}
			var actual []int
			for pq.Len() > 0 {
				actual = append(actual, pq.Pop())
			}

			// assert
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("priority queue atomic () = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestAtomicPriorityWithTime(t *testing.T) {
	type args[T any] struct {
		cap    int
		less   LessFun[T]
		inputs []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	type store struct {
		name string
		ttl  time.Time
	}
	now := time.Now()
	tests := []testCase[store]{
		{
			"initialTest",
			args[store]{
				cap: 10,
				less: func(a store, b store) bool {
					return a.ttl.Before(b.ttl)
				},
				inputs: []store{
					{"time at 10", now.Add(time.Second * 10)},
					{"time at 9", now.Add(time.Second * 9)},
					{"time at 2", now.Add(time.Second * 2)},
					{"time at 7", now.Add(time.Second * 7)},
					{"time at 8", now.Add(time.Second * 8)},
					{"time at 1", now.Add(time.Second * 1)},
					{"time at 6", now.Add(time.Second * 6)},
					{"time at 5", now.Add(time.Second * 5)},
					{"time at 3", now.Add(time.Second * 3)},
					{"time at 4", now.Add(time.Second * 4)},
				},
			},
			[]store{
				{"time at 1", now.Add(time.Second * 1)},
				{"time at 2", now.Add(time.Second * 2)},
				{"time at 3", now.Add(time.Second * 3)},
				{"time at 4", now.Add(time.Second * 4)},
				{"time at 5", now.Add(time.Second * 5)},
				{"time at 6", now.Add(time.Second * 6)},
				{"time at 7", now.Add(time.Second * 7)},
				{"time at 8", now.Add(time.Second * 8)},
				{"time at 9", now.Add(time.Second * 9)},
				{"time at 10", now.Add(time.Second * 10)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			pq := NewAtomicPriorityQueue(tt.args.cap, tt.args.less)

			// act
			for _, it := range tt.args.inputs {
				pq.Push(it)
			}
			var actual []store
			for pq.Len() > 0 {
				actual = append(actual, pq.Pop())
			}

			// assert
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("priority queue atomic () = %v, want %v", actual, tt.want)
			}
		})
	}
}

func TestAtomicPriorityQueueHead(t *testing.T) {
	type args[T any] struct {
		cap    int
		less   LessFun[T]
		inputs []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}

	type store struct {
		name string
		ttl  time.Time
	}
	now := time.Now()
	tests := []testCase[store]{
		{
			"initialTest",
			args[store]{
				cap: 10,
				less: func(a store, b store) bool {
					return a.ttl.Before(b.ttl)
				},
				inputs: []store{
					{"time at 10", now.Add(time.Second * 10)},
					{"time at 9", now.Add(time.Second * 9)},
					{"time at 2", now.Add(time.Second * 2)},
					{"time at 7", now.Add(time.Second * 7)},
					{"time at 8", now.Add(time.Second * 8)},
					{"time at 1", now.Add(time.Second * 1)},
					{"time at 6", now.Add(time.Second * 6)},
					{"time at 5", now.Add(time.Second * 5)},
					{"time at 3", now.Add(time.Second * 3)},
					{"time at 4", now.Add(time.Second * 4)},
				},
			},
			store{
				"time at 1", now.Add(time.Second * 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			pq := NewAtomicPriorityQueue(tt.args.cap, tt.args.less)

			// act
			for _, it := range tt.args.inputs {
				pq.Push(it)
			}
			var actual store
			actual = pq.Head()

			// assert
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("priority queue atomic () = %v, want %v", actual, tt.want)
			}
		})
	}
}
