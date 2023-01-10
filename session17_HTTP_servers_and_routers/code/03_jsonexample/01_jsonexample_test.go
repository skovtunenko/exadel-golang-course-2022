package jsonexample

import (
	"encoding/json"
	"os"
	"testing"
)

func TestJSONmarshaling(t *testing.T) {
	bolB, _ := json.Marshal(true)
	t.Log(string(bolB))

	intB, _ := json.Marshal(1)
	t.Log(string(intB))

	fltB, _ := json.Marshal(2.34)
	t.Log(string(fltB))

	strB, _ := json.Marshal("gopher")
	t.Log(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	t.Log(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	t.Log(string(mapB))
}

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func TestJSONmarshal(t *testing.T) {
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	t.Log(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	t.Log(string(res2B))
}

func TestUnknownJSONstructure(t *testing.T) {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	t.Log(dat)

	num := dat["num"].(float64)
	t.Log(num)
	// Accessing nested data requires a series of conversions:
	strs := dat["strs"].([]interface{}) // !!!!!
	str1 := strs[0].(string)
	t.Log(str1)
}

func TestJSONCustomDataType(t *testing.T) {
	// We can also decode JSON into custom data types:

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	t.Log(res)
	t.Log(res.Fruits[0])
}

func TestStreamJSON(t *testing.T) {
	// We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies:

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
