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
		return
	}

	port, err := strconv.Atoi(os.Args[1])
	// fmt.Println(os.Args[1])
	if err != nil {
		fmt.Println("Incorrect port number")
		return
	}

	if port < 1 || port > 65535 {
		fmt.Println("Port number out of bound")
		return
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
			}
		} else {
			fmt.Println(err)
		}
		return
	}

	fmt.Print(string(output))
}
