package structexample

import (
	"encoding/json"
	"testing"
)

type Core struct {
	ID  string `json:"id"`
	Num int    `json:"num"`
}

type Data struct {
	Core      // `json:"core"` // <-- try to comment/uncomment struct tag
	Count int `json:"count"`
}

func TestJsonMarshaling(t *testing.T) {
	data := Data{
		Core: Core{
			ID:  "1",
			Num: 123,
		},
		Count: 777,
	}

	out, err := json.MarshalIndent(data, "  ", "  ")
	if err != nil {
		panic(err) // for simplicity
	}
	t.Logf("\n%s\n", out)
}
