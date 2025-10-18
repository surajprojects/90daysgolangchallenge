package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Execute "ls -la" (Linux/macOS) or "dir" (Windows)
	// cmd := exec.Command("ls", "-la") // Linux/macOS
	cmd := exec.Command("cmd", "/C", "dir") // Windows
	cmd.Stdout = nil                        // Will default to os.Stdout if nil
	cmd.Stderr = nil

	output, err := cmd.CombinedOutput() // Capture stdout + stderr
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	fmt.Println("Command Output:\n", string(output))
}
