package testspart1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGreetingSimple(t *testing.T) {
	// 1. Setup (optional)
	// nothing here

	// 2. Execute the code under testing
	got := Greeting("Sergio")

	// 3. Check (assert) the results
	want := "Greeting, Sergio !"
	if got != want {
		t.Fatalf("Greeting() = %v, want %v", got, want)
		// t.Errorf("Greeting() = %v, want %v", got, want)
	}

	// 4. Tear-down (optional)
	// nothing here
}

// !! install testify library:
// go get github.com/stretchr/testify
// go mod tidy

func TestGreetingWithSubtest(t *testing.T) {
	// 1. Setup (optional)
	// nothing here

	t.Run("non-empty_user_name", func(t *testing.T) {
		// 2. Execute the code under testing
		got := Greeting("Sergio")

		// 3. Check (assert) the results
		want := "Greeting, Sergio !"
		require.Equal(t, want, got)
	})

	t.Run("empty_user_name", func(t *testing.T) {
		// 2. Execute the code under testing
		got := Greeting("")

		// 3. Check (assert) the results
		want := ""
		require.Equal(t, want, got)
	})

	// 4. Tear-down (optional)
	// nothing here
}
