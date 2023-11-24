package main

import (
	"os"
	"fmt"

	"gopkg.in/yaml.v3"
)

func writeMediaMTXyml(mtxyml_path string, streamMap map[string]interface{}, ymlMap map[string]interface{}) error {

	// Edit ymlMap based on streamMap
	// readUser := streamMap["user"]
	// readPass := streamMap["pass"]
	// data := ymlMap["paths"]
	// pdata := data["proxied0"]
	// user := pdata["readUser"]

	if nestedValue, ok := ymlMap["paths"].(map[string]interface {}); ok {
		// key3 is a map, access its nested values
		if nestedValue1, ok := nestedValue["proxied0"].(map[string]interface {}); ok {
			fmt.Printf("Nested Value: %v\n", nestedValue1)
			if nestedKey, ok := nestedValue1["readPass"]; ok {
				fmt.Printf("Nested Value: %v\n", nestedKey)
			} else {
				fmt.Println("Nested key 'nested_key' not found.")
			}
		} else {
			fmt.Println("Nested key 'nested_key' not found.")
		}
	} else {
		fmt.Println("Key 'key3' not found or not a map.")
	}
	// ((ymlMap["paths"].(map[string]interface {}))["proxied0"].(map[string]interface {}))["readPass"]


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

	return nil

}
