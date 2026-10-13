package main

import "testing"

func TestParsePort(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantPort int
		wantErr  bool
	}{
		{name: "typical port", input: "3000", wantPort: 3000},
		{name: "minimum port", input: "1", wantPort: 1},
		{name: "maximum port", input: "65535", wantPort: 65535},
		{name: "zero", input: "0", wantErr: true},
		{name: "negative", input: "-1", wantErr: true},
		{name: "above maximum", input: "65536", wantErr: true},
		{name: "invalid text", input: "hello", wantErr: true},
		{name: "empty input", input: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port, err := parsePort(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("parsePort(%q): error = %v, wantErr = %t",
					tt.input, err, tt.wantErr)
			}

			if port != tt.wantPort {
				t.Fatalf("parsePort(%q): port = %d, want %d",
					tt.input, port, tt.wantPort)
			}
		})
	}
}
