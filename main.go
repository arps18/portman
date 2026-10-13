package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: portman <port>")
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Incorrect port number")
		os.Exit(1)
	}

	if port < 1 || port > 65535 {
		fmt.Println("Port number out of bound")
		os.Exit(1)
	}

	cmd := exec.Command("lsof", "-nP", "-iTCP:"+strconv.Itoa(port), "-sTCP:LISTEN")
	output, err := cmd.Output()

	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			if exitErr.ExitCode() == 1 && len(exitErr.Stderr) == 0 {
				fmt.Println("No TCP listeners found on port", port)
			} else {
				fmt.Println(err)
				fmt.Print(string(exitErr.Stderr))
				os.Exit(1)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	fmt.Print(string(output))
}
