package methodexample

import "testing"

type Total struct {
	data int
}

func (t *Total) Get() int {
	return t.data
}

func (t *Total) Add(n int) {
	t.data += n
}

// BAD STYLE (mixing value and pointer receivers):
func (t Total) Sub(n int) { // there is a bug in this method: we need to use pointer receiver
	t.data -= n
}

func TestTotal(t *testing.T) {
	total := Total{}

	total.Add(10) // !!! This is same as (&total).Add(10)
	total.Add(20)
	if total.Get() != 30 {
		t.Fatalf("want: %d, got: %d", 30, total.Get())
	}

	total.Sub(30)
	if total.Get() != 0 { // test will fail, if the `Sub()` method is defined on value receiver
		t.Fatalf("want: %d, got: %d", 0, total.Get())
	}
}
