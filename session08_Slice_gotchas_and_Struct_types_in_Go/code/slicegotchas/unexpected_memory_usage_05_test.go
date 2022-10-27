package slicegotchas

import (
	"fmt"
	"testing"
)

func brokenGet() []byte {
	raw := make([]byte, 10000)               // obtain huge slice
	fmt.Println(len(raw), cap(raw), &raw[0]) // prints: 10000 10000 <byte_addr_x>
	return raw[:3]                           // re-slice to return just a small sub-slice
}

func TestBrokenGet(t *testing.T) {
	data := brokenGet()
	t.Log(len(data), cap(data), &data[0]) // prints: 3 10000 <byte_addr_x>
}

func fixedGet() []byte {
	raw := make([]byte, 10000)               // get huge slice
	fmt.Println(len(raw), cap(raw), &raw[0]) // prints: 10000 10000 <byte_addr_x>

	res := make([]byte, 3)
	copy(res, raw[:3]) // copy just important subslice, `raw` will be garbage collected after
	return res
}

func TestFixedGet(t *testing.T) {
	data := fixedGet()
	t.Log(len(data), cap(data), &data[0]) // prints: 3 3 <byte_addr_y>
}
