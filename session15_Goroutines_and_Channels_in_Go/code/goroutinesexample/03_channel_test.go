package goroutinesexample

import (
	"fmt"
	"testing"
)

func TestChannelBasic(t *testing.T) {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func TestChannelBuffering(t *testing.T) {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func TestChannelDirections(t *testing.T) {
	ping := func(pings chan<- string, msg string) {
		pings <- msg
	}

	pong := func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
