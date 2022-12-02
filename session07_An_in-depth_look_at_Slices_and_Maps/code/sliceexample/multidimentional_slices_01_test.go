package sliceexample

import (
	"fmt"
	"strings"
)

func ExampleMultidimensionalUsingMake() {
	ss := make([][]uint8, 2) // ss variable is []([]uint8)
	fmt.Printf("ss:    %T %v %d/%d\n", ss, ss, len(ss), cap(ss))
	for i, s := range ss { // s is []uint8
		s = make([]uint8, 10, 20) // initialize internal slices!
		fmt.Printf("ss[%d]: %T %v %d/%d\n", i, s, s, len(s), cap(s))
	}
	// OUTPUT:
	// ss:    [][]uint8 [[] []] 2/2
	// ss[0]: []uint8 [0 0 0 0 0 0 0 0 0 0] 10/20
	// ss[1]: []uint8 [0 0 0 0 0 0 0 0 0 0] 10/20
}

func ExampleMultidimensionalSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[2][0] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Output:
	// X _ X
	// O _ _
	// X _ O
}
