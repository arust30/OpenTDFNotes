package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Create a new command to execute the otdfctl binary
	cmd := exec.Command("otdfctl", "-h")

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting version:", err)
		return
	}

	// Print the version number
	version := strings.TrimSpace(string(output))[len("OpenTDF version:"):]
	fmt.Println("OpenTDF version:", version)
}

