package main

import (
	"fmt"
	"os"
)

func writeToFile(filename, content string) error {
	// Write the content to the file
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Content written to %s\n", filename)
	return nil
}