package httpc

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

func server(path, port string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprintf(w, "Sorry I'm a bit slow...")
	})
	http.ListenAndServe(port, nil)
}

func TestTimeout(t *testing.T) {
	go server("/timeout", ":8080")
	client := NewClient(30 * time.Millisecond)
	_, err := client.Get("http://localhost:8080/timeout")
	if !strings.Contains(err.Error(), "request canceled") {
		t.Fatalf("Expected to fail because of timeout")
	}
}

func TestNoTimeout(t *testing.T) {
	go server("/notimeout", ":9090")
	client := DefaultClient
	_, err := client.Get("http://localhost:9090/notimeout")
	if err != nil {
		t.Fatalf("Expected error to be nil")
	}
}
