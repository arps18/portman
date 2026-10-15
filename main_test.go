package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

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

func TestInspectPortMissingLsof(t *testing.T) {
	t.Setenv("PATH", t.TempDir())

	output, err := inspectPort(3000)

	if err == nil {
		t.Fatal("expected an error when lsof is unavailable")
	}

	if output != "" {
		t.Fatalf("expected empty output, got %q", output)
	}

	if !strings.Contains(err.Error(), "lsof not found in PATH") {
		t.Fatalf("expected missing-lsof message, got %v", err)
	}
}

func TestCLI(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "portman")

	build := exec.Command("go", "build", "-o", binary, ".")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, output)
	}

	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{
			name:       "help",
			args:       []string{"--help"},
			wantStdout: "Usage:",
		},
		{
			name:       "short help",
			args:       []string{"-h"},
			wantStdout: "Usage:",
		},
		{
			name:       "version",
			args:       []string{"--version"},
			wantStdout: "dev\n",
		},
		{
			name:       "missing argument",
			wantCode:   1,
			wantStderr: "Usage:",
		},
		{
			name:       "extra argument",
			args:       []string{"3000", "4000"},
			wantCode:   1,
			wantStderr: "Usage:",
		},
		{
			name:       "invalid text",
			args:       []string{"hello"},
			wantCode:   1,
			wantStderr: `invalid port "hello":`,
		},
		{
			name:       "out of range",
			args:       []string{"70000"},
			wantCode:   1,
			wantStderr: "port must be between 1 and 65535",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binary, tt.args...)

			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			code := 0
			if err := cmd.Run(); err != nil {
				exitErr, ok := err.(*exec.ExitError)
				if !ok {
					t.Fatalf("could not run CLI: %v", err)
				}
				code = exitErr.ExitCode()
			}

			if code != tt.wantCode {
				t.Errorf("exit status = %d, want %d", code, tt.wantCode)
			}

			if tt.wantStdout == "" {
				if stdout.Len() != 0 {
					t.Errorf("expected empty stdout, got %q", stdout.String())
				}
			} else if !strings.Contains(stdout.String(), tt.wantStdout) {
				t.Errorf("stdout = %q, want it to contain %q",
					stdout.String(), tt.wantStdout)
			}

			if tt.wantStderr == "" {
				if stderr.Len() != 0 {
					t.Errorf("expected empty stderr, got %q", stderr.String())
				}
			} else if !strings.Contains(stderr.String(), tt.wantStderr) {
				t.Errorf("stderr = %q, want it to contain %q",
					stderr.String(), tt.wantStderr)
			}
		})
	}
}
