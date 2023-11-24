package main

import (
	"fmt"
	"os/exec"
	"time"
)

func runMediaMTX() error {

	maxAttempts := 2
	attempt := 1

	for {

		cmd := exec.Command("./mediamtx")

		// Run the command and check for errors
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error running command: %v\n", err)
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
