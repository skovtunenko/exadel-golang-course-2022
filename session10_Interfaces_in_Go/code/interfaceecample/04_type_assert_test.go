package interfaceecample

import (
	"testing"
	"time"
)

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

func TestTypeAssert(t *testing.T) {
	foo := map[string]interface{}{
		"Matt": 42,
	}

	timeMap(foo)

	t.Log(foo)
}
