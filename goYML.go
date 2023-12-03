package main

import (
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RunOnDisconnect string `yaml:"runOnDisconnect"`
	RTSP            string `yaml:"rtsp"`
	WebRTC          string `yaml:"webrtc"`
	RTMP            string `yaml:"rtmp"`
	HLS             string `yaml:"hls"`
	SRT             string `yaml:"srt"`

	Paths struct {
		Proxied struct {
		ReadUser string `yaml:"readUser"`
		ReadPass string `yaml:"readPass"`
		} `yaml:"proxied"`

		AllOthers struct{} `yaml:"all_others"`
	} `yaml:"paths"`
}

func removeQuotesFromFile(filename string) error {
	// Read the file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Replace all double-quote symbols
	re := regexp.MustCompile(`"`)
	modifiedContent := re.ReplaceAllString(string(content), "")

	// Write the modified content back to the file
	err = os.WriteFile(filename, []byte(modifiedContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func readConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func writeConfig(filename string, newMap map[string]interface{}, config *Config, streamName string) error {
	
	// set value in newMap from firebase specific stream to new yml config
	config.Paths.Proxied.ReadUser = strings.TrimSpace(newMap["user"].(string))
	config.Paths.Proxied.ReadPass = strings.TrimSpace(newMap["pass"].(string))


	// set value in runOnDisconnect 
	cmds := strings.Split(config.RunOnDisconnect," ")
	config.RunOnDisconnect = cmds[0] + " " + streamName

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	errR := removeQuotesFromFile(filename)
	if errR != nil {
		return errR
	}

	return nil
}
