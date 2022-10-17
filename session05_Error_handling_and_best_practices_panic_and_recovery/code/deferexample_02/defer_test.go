package deferexample_02

import "fmt"

func ExampleDefer() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}

	// OUTPUT:
	// 3210
}

func ExampleHeaderPrinting() {
	printHeader()
	defer printFooter()

	for i := 0; i < 3; i++ {
		fmt.Printf("%d...", i+1)
	}
	fmt.Println()

	// OUTPUT:
	// <--- THIS IS HEADER --->
	// 1...2...3...
	// <--- THIS IS FOOTER --->
}

func printFooter() {
	fmt.Println("<--- THIS IS FOOTER --->")
}

func printHeader() {
	fmt.Println("<--- THIS IS HEADER --->")
}
