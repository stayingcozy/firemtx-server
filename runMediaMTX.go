package main

import (
	"fmt"
	"os/exec"
	"time"
	"os"
)

func runMediaMTX() error {

	maxAttempts := 2
	attempt := 1

	for {

		cmd := exec.Command("./mediamtx")

		// Create pipes to capture command output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the command and check for errors
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error running mediamtx command: %v\n", err)
			return err
		}

		if attempt >= maxAttempts {
			return fmt.Errorf("max attempts reached, unable to start mediamtx")
		}

		// Increment the attempt counter and wait for a moment before retrying
		attempt++
		time.Sleep(100 * time.Millisecond)
	}
}
