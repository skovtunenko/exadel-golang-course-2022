package httptestexample

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func handlerUnderTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>Hello World!</body></html>")
}

func Test_01_HandlerUnderTest(t *testing.T) {
	// step 1: prepare incoming parameters
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	// step 2: invoke HTTP handler under test
	handlerUnderTest(w, req)

	// step 3: get response and analyse it
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	t.Log(resp.StatusCode)

	require.Equal(t, `text/html; charset=utf-8`, resp.Header.Get("Content-Type"))
	t.Log(resp.Header.Get("Content-Type"))

	require.Equal(t, `<html><body>Hello World!</body></html>`, string(body))
	t.Log(string(body))
}

func Test_02_HTTPtestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handlerUnderTest))
	defer ts.Close()

	// other way to configure server manually before start:
	// unstartedServer := httptest.NewUnstartedServer(http.HandlerFunc(handlerUnderTest))
	// unstartedServer.EnableHTTP2=true
	// unstartedServer.Start()
	// defer unstartedServer.Close()

	res, err := http.Get(ts.URL)
	require.NoError(t, err)

	greeting, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	res.Body.Close()

	t.Logf("%s", greeting)
}
