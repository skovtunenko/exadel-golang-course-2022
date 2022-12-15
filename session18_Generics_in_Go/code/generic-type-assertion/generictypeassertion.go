package main

import "fmt"

func main() {
	Greet("world")
	Greet(777)
}

func Greet[T any](val T) {
	//switch val.(type) { // <------ WILL NOT COMPILE

	// we need to cast to interface first:
	switch (any)(val).(type) {
	case string:
		fmt.Println("hello, string")
	case int:
		fmt.Println("hello, int")
	default:
		fmt.Println("not supported type")
	}
}
