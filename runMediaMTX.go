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

		// ffmpeg command, h264 video codec
		cmd := exec.Command(
			"./mediamtx",
		)

		err := cmd.Run()

		// if err == nil {
		// 	// Command succeeded, break out of the loop
		// 	break
		// }

		fmt.Printf("Attempt %d failed: %v\n", attempt, err)

		if attempt >= maxAttempts {
			return fmt.Errorf("max attempts reached, unable to start mediamtx")
		}

		// Increment the attempt counter and wait for a moment before retrying
		attempt++
		time.Sleep(1 * time.Second)
	}

	// return nil

}
