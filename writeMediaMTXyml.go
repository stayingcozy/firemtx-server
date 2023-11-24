package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

func writeMediaMTXyml(mtxyml_path string, streamMap map[string]interface{}, ymlMap map[string]interface{}) error {

	// Edit ymlMap based on streamMap
	if paths, ok := ymlMap["paths"].(map[string]interface {}); ok {
		if proxied0, ok := paths["proxied0"].(map[string]interface {}); ok {
			if readPass, ok := proxied0["readPass"]; ok {
				// Print the original value
				fmt.Printf("Original Value: %v\n", readPass)

				// Modify the original value
				proxied0["readPass"] = streamMap["pass"]
			} else {
				fmt.Println("Nested key 'readPass' not found.")
			}
			if readPass, ok := proxied0["readUser"]; ok {
				// Print the original value
				fmt.Printf("Original Value: %v\n", readPass)

				// Modify the original value
				proxied0["readUser"] = streamMap["user"]

			} else {
				fmt.Println("Nested key 'readPass' not found.")
			}
		} else {
			fmt.Println("Nested key 'proxied0' not found.")
		}
	} else {
		fmt.Println("Key 'paths' not found or not a map.")
	}

	// Marshal the map into YAML format
	yamlBytes, err := yaml.Marshal(ymlMap)
	if err != nil {
		return err
	}

	// Write the YAML data to the file
	err = os.WriteFile(mtxyml_path, yamlBytes, 0644)
	if err != nil {
		return err
	}

	time.Sleep(1 *time.Second) // Allow some time for mediamtx to register the yml change and rerun accordingly

	return nil

}
