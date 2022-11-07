package testspart1

func Greeting(name string) string {
	// log.Println("got argument:", name) // non-pure function

	return "Greeting, " + name + " !" // pure function
}
