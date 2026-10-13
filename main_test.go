package main

import "testing"

func TestParsePort(t *testing.T) {
	port, err := parsePort("3000")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if port != 3000 {
		t.Fatalf("expected port 3000, got %d", port)
	}
}
