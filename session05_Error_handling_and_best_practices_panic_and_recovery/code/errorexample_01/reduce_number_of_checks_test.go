package errorexample_01

import (
	"bufio"
	"strings"
	"testing"
)

// Analyze bufio.Scanner type. Definition:

// type Scanner struct {
// 	r            io.Reader // The reader provided by the client.
// 	split        SplitFunc // The function to split the tokens.
// 	maxTokenSize int       // Maximum size of a token; modified by tests.
// 	token        []byte    // Last token returned by split.
// 	buf          []byte    // Buffer used as argument to split.
// 	start        int       // First non-processed byte in buf.
// 	end          int       // End of data in buf.
// 	err          error     // Sticky error.    						<-------------- take a look at this!
// 	empties      int       // Count of successive empty tokens.
// 	scanCalled   bool      // Scan has been called; buffer is in use.
// 	done         bool      // Scan has finished.
// }

func TestReduceNumberOfErrorChecks(t *testing.T) {
	// Analyze bufio.NewScanner type

	input := strings.NewReader("some\nmultiline\ntest\ninput")

	scanner := bufio.NewScanner(input)

	// perform some operations:
	for scanner.Scan() {
		token := scanner.Text()
		// process token....
		t.Log(token)
	}

	// Check for error only once after:
	if err := scanner.Err(); err != nil {
		// process the error
		t.Log("there was an error :(")
	}
}

// NOTE:

// Sure, there is a nil check for an error, but it appears and executes only once.
// 	⁃	With the real API, the client's code therefore feels more natural: loop until done, then worry about errors.
//	⁃	In this case error handling does not obscure the flow of control.
