package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Portfolio!"))
    })

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected 200, got %v", status)
    }

    expected := "Hello, Portfolio!"
    if rr.Body.String() != expected {
        t.Errorf("expected %q, got %q", expected, rr.Body.String())
    }
}
