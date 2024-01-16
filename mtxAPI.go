package main

import(
	"fmt"
	"time"
)

func initAPI(server_url string) {
	post_resource := "/v3/config/paths/add/"
	setting := "readUser"
	valbod := []byte(`{"readUser": "", "readPass": ""}`)
	post(server_url, post_resource+setting, valbod)
}

func patchAPI(server_url string, streamName string, cmd string, user string, pass string) {

	// Set new command
	cmd_id := cmd + " " + streamName

	// patch path
	patch_res := "/v3/config/paths/patch/"
	setting := "readUser"
	valbod := []byte(fmt.Sprintf(`{"readUser": "%s", "readPass": "%s"}`, user, pass))
	patch(server_url, patch_res+setting, valbod)

	// patch global
	patch_resource := "/v3/config/global/patch"
	valbodg := []byte(fmt.Sprintf(`{"runOnDisconnect": "%s"}`, cmd_id))
	patch(server_url, patch_resource, valbodg)

	time.Sleep(200 * time.Millisecond)
}