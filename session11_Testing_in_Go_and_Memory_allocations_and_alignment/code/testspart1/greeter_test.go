package testspart1

import "testing"

func TestGreeting1(t *testing.T) {
	// - Setup (optional)

	// - Execute the code under testing
	gotGreeting := Greeting("someName")

	// - Check (assert) the results
	want := "Greeting, someName !"
	if want != gotGreeting {
		t.FailNow()
	}

	// - Tear-down (optional)
}
