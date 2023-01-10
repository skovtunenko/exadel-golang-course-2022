package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/headers", dumpHeaders)

	const addr = ":3000"

	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux, // or http.DefaultServeMux,
	}

	fmt.Println("starting HTTP server on", addr)
	err := httpServer.ListenAndServe()
	// OR:
	// err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel() // super important!

	// Common code for all requests can go here...

	// non-blocking check:
	select {
	case <-ctx.Done():
		msg := fmt.Sprintf("context is canceled because of: %q", ctx.Err())
		http.Error(w, msg, http.StatusGatewayTimeout)
		return
	default:
	}

	// another way to check context:
	if err := ctx.Err(); err != nil {
		msg := fmt.Sprintf("context is canceled because of: %q", err)
		http.Error(w, msg, http.StatusGatewayTimeout)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Handle the GET request...

		fmt.Fprintf(w, "hello from %q method\n", r.Method)

	case http.MethodPost:
		// Handle the POST request...

		fmt.Fprintf(w, "hello from %q method\n", r.Method)

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func dumpHeaders(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
