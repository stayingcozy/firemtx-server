package main

import (
	"strings"
)

func readMap(newMap map[string]interface{}) (string,string,string) {

	// Set new user and password for authentication
	user := strings.TrimSpace(newMap["user"].(string))
	pass := strings.TrimSpace(newMap["pass"].(string))
	serverIPStream := strings.TrimSpace(newMap["server"].(string))

	return user, pass, serverIPStream
}