package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Run "echo Hello Tiger" command
	cmd := exec.Command("echo", "Hello Tiger 🐯")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Command Output:", string(output))
}
