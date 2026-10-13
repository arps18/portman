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

func TestParsePortInvalidText(t *testing.T) {
	port, err := parsePort("hello")

	if err == nil {
		t.Fatal("expected an error for invalid text")
	}

	if port != 0 {
		t.Fatalf("expected port 0 on failure, got %d", port)
	}
}
