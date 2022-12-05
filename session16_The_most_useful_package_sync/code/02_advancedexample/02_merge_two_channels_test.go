package advancedexample

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func mergeTwo(in1 <-chan int, in2 <-chan int) <-chan int {
	out := make(chan int) // step 1

	go func() { // step 2
		defer close(out)

		for in1 != nil || in2 != nil {
			select {
			case val, ok := <-in1:
				if !ok {
					in1 = nil
					continue
				}
				out <- val
			case val, ok := <-in2:
				if !ok {
					in2 = nil
					continue
				}
				out <- val
			}
		}
	}()

	return out // step 3
}

func TestMergeTwoChannels(t *testing.T) {
	in1 := asChan(1, 1, 1)
	in2 := asChan(2, 2, 2)

	for val := range mergeTwo(in1, in2) {
		fmt.Println(val)
	}

	fmt.Println("Done.")
}

func Test_merge(t *testing.T) {
	type args struct {
		ch1 <-chan int
		ch2 <-chan int
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want []int
	}{
		{
			name: "empty chans",
			args: func(t *testing.T) args {
				t.Helper()
				return args{
					ch1: asChan(),
					ch2: asChan(),
				}
			},
			want: []int{},
		},
		{
			name: "empty first chan",
			args: func(t *testing.T) args {
				t.Helper()
				return args{
					ch1: asChan(),
					ch2: asChan(1, 2, 3),
				}
			},
			want: []int{1, 2, 3},
		},
		{
			name: "empty second chan",
			args: func(t *testing.T) args {
				t.Helper()
				return args{
					ch1: asChan(1, 2, 3),
					ch2: asChan(),
				}
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args(t)
			gotCh := mergeTwo(args.ch1, args.ch2)
			got := []int{}
			for val := range gotCh {
				got = append(got, val)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
