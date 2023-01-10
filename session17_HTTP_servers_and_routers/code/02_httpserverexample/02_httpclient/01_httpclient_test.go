package httpclient

import (
	"io"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		t.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", body)
}
