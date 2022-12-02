package interfaceecample

import (
	"testing"
)

func TestFailedTypeAssertion(t *testing.T) {
	var data interface{} = "great" // actually a string value

	if res, ok := data.(int); ok {
		t.Log("[is an int] value =>", res)
	} else {
		t.Log("[not an int] value =>", res) // This is a BUG!

		// Failed type assertions return the "zero value" for the target type used in the assertion statement:
		// prints: [not an int] value => 0 (not "great")

		t.Log("[not an int] value =>", data) // OK!
		// prints: [not an int] value => great (as expected)
	}
}
