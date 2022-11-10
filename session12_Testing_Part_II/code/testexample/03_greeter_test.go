package testspart1

import "fmt"

func ExampleGreeting() {
	message := Greeting("Sergio")
	fmt.Printf("Got: %s", message)

	// Output:
	// Got: Greeting, Sergio !
}
