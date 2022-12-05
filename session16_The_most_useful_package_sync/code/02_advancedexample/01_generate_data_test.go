package advancedexample

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func asChan(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, val := range nums {
			ch <- val
		}
	}()
	return ch
}

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	defer close(out)
	for _, n := range nums {
		out <- n
	}
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	fn := func(val int) int { return val * val }
	go func() {
		defer close(out)
		for val := range in {
			out <- fn(val)
		}
	}()
	return out
}

// ---------------------------------------------------------------
// THIS IS TO DEMONSTRATE THE FLOW:
// ---------------------------------------------------------------

func TestExecuteDataPipeline(t *testing.T) {
	c := asChan(2, 3)
	// OR :
	// c := gen(2, 3)

	out := sq(c)

	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

// ---------------------------------------------------------------
// THIS IS TO SHOW HOW TO TEST CHAN-RELATED FUNCTIONS:
// ---------------------------------------------------------------

func Test_asChan(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil arguments",
			args: args{nums: nil},
			want: []int{},
		},
		{
			name: "empty incomming array",
			args: args{nums: make([]int, 0)},
			want: []int{},
		},
		{
			name: "nil arguments",
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotCh := asChan(tt.args.nums...)
			got := make([]int, 0, len(tt.args.nums))
			for val := range gotCh {
				got = append(got, val)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_sq(t *testing.T) {
	type args struct {
		in func(t *testing.T) <-chan int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "closed chan",
			args: args{
				in: func(t *testing.T) <-chan int {
					ch := make(chan int)
					defer close(ch)
					return ch
				},
			},
			want: []int{},
		},
		{
			name: "chan with one value",
			args: args{
				in: func(t *testing.T) <-chan int {
					return asChan(1)
				},
			},
			want: []int{1},
		},
		{
			name: "chan with many values",
			args: args{
				in: func(t *testing.T) <-chan int {
					return asChan(1, 2, 3)
				},
			},
			want: []int{1, 4, 9},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotCh := sq(tt.args.in(t))
			got := []int{}
			for val := range gotCh {
				got = append(got, val)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
