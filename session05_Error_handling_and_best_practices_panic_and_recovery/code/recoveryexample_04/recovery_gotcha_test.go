package recoveryexample_04

import (
	"fmt"
	"testing"
)

func doRecover() {
	fmt.Println("recovered =>", recover()) // prints: recovered => <nil>
}

func Test_recovery_gotcha(t *testing.T) {
	defer func() {
		doRecover() // panic is not recovered!!!

		// This is a way to fix the problem (uncomment lines):

		// if err := recover(); err != nil {
		// 	t.Log("Fixed: recovered from", err)
		// }
	}()

	panic("not good")
}
