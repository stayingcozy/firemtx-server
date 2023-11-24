package main

import(
	"os"

	"gopkg.in/yaml.v3"
)

func readMediaMTXyml(mtxyml_path string) (map[string]interface{}, error) {

	// Read the YAML file
	yamlFile, err := os.ReadFile(mtxyml_path)
	if err != nil {
		return nil, err
	}

	// Create a map to store the YAML content
	yamlData := make(map[string]interface{})

	// Unmarshal the YAML data into the map
	err = yaml.Unmarshal(yamlFile, yamlData)
	if err != nil {
		return nil, err
	}

	return yamlData, nil
}