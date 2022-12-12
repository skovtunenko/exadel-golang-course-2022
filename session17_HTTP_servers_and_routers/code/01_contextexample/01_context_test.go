package contextexample

import (
	"context"
	"testing"
	"time"
)

func TestContextWithCancel(t *testing.T) {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		t.Log(n)
		if n == 5 {
			break
		}
	}
}

func TestContextWithDeadline(t *testing.T) {
	const shortDuration = 1 * time.Millisecond

	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Log("overslept")
	case <-ctx.Done():
		t.Log(ctx.Err())
	}
}

func TestContextWithTimeout(t *testing.T) {
	const shortDuration = 1 * time.Millisecond

	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Log("overslept")
	case <-ctx.Done():
		t.Log(ctx.Err()) // prints "context deadline exceeded"
	}
}

func TestWithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			t.Log("found value:", v)
			return
		}
		t.Log("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}
