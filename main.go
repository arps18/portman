package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var version = "dev"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: portman <port>")
		os.Exit(1)
	}
	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		return
	}
	if os.Args[1] == "--version" {
		printVersion()
		return
	}
	port, err := parsePort(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	output, err := inspectPort(port)

	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprint(os.Stderr, string(exitErr.Stderr))
			os.Exit(1)
		} else {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if output == "" {
		fmt.Println("No TCP listeners found on port", port)
		return
	}
	fmt.Print(output)

}

func parsePort(input string) (int, error) {
	port, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid port %q: expected a whole number between 1 and 65535", input)
	}

	if port < 1 || port > 65535 {
		return 0, fmt.Errorf("port must be between 1 and 65535")
	}
	return port, nil
}

func inspectPort(port int) (string, error) {
	cmd := exec.Command("lsof", "-nP", "-iTCP:"+strconv.Itoa(port), "-sTCP:LISTEN")
	output, err := cmd.Output()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return "", errors.New("lsof not found in PATH: Portman requires lsof to inspect TCP listeners")
		}
		exitErr, ok := err.(*exec.ExitError)
		if ok && exitErr.ExitCode() == 1 && len(exitErr.Stderr) == 0 {
			return "", nil
		}
		return "", err
	}

	return string(output), nil
}

func printHelp() {
	fmt.Print(`Portman — inspect local TCP listeners

Usage:
  portman <port>
  portman --help

Options:
  -h, --help  Show help
 --version	  Show version   

Example:
  portman 3000
`)
}

func printVersion() {
	fmt.Println(version)
}
