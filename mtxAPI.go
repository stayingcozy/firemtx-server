package main

import(
	"fmt"
	// "net/http"
	// "bytes"
	// "io"
	"time"
	"strings"
)

// func patch(url string, resource string, jsonBody []byte) {

// 	// resource := "/v3/config/global/patch"
// 	url = url + resource

// 	bodyReader := bytes.NewReader(jsonBody)

// 	req, err := http.NewRequest(http.MethodPatch, url, bodyReader)
// 	if err != nil {  
// 		fmt.Printf("client: could not create request: %s\n", err)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Printf("client: error making http request: %s\n", err)
// 	}

// 	// fmt.Printf("client: got response!\n")
// 	fmt.Printf("client: status code: %d\n", res.StatusCode)

// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Printf("client: could not read response body: %s\n", err)
// 	}
// 	fmt.Printf("client: response body: %s\n", resBody)
// }

// func post(url string, resource string, jsonBody []byte) error {

// 	url = url + resource

// 	// jsonBody := []byte(`{"readUser": "user0", "readPass": "pass0"}`)
// 	bodyReader := bytes.NewReader(jsonBody)

// 	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
// 	if err != nil {
// 		fmt.Printf("client: could not create request: %s\n", err)
// 		return err
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Printf("client: error making http request: %s\n", err)
// 		return err
// 	}

// 	// fmt.Printf("client: got response!\n")
// 	fmt.Printf("client: status code: %d\n", res.StatusCode)

// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Printf("client: could not read response body: %s\n", err)
// 		return err
// 	}
// 	fmt.Printf("client: response body: %s\n", resBody)

// 	return nil
// }

// func runOnDisconnectPatch(server_url string, cmd string) {
// 	rodBody := []byte(fmt.Sprintf(`{"runOnDisconnect": "%s"}`, cmd))
// 	resourceROD := "/v3/config/global/patch"
// 	patch(server_url, resourceROD, rodBody)
// }

// func readPost(server_url string) {
// 	readBody := []byte(`{"readUser": "user0", "readPass": "pass0"}`) 
// 	resource := "/v3/config/paths/add/readUser"

// 	// For some reason the first run give EOF, then the next one works
// 	post(server_url, resource, readBody)
// 	time.Sleep(150 * time.Millisecond) // have to give REST API some time or it'll shit itself

// 	// post(server_url, resource, readBody)
// 	// time.Sleep(150 * time.Millisecond)

// }

// func readPatch(server_url string, user string, pass string) {
// 	resourceBody := "/v3/config/paths/patch/readUser"
// 	readBody := []byte(fmt.Sprintf(`{"readUser": "%s", "readPass": "%s"}`,user,pass)) 

// 	patch(server_url, resourceBody, readBody)
// }

// func initAPI_old(server_url string, cmd string) {
// 	runOnDisconnectPatch(server_url, cmd)
// 	readPost(server_url)

// // 	user := "user0"
// // 	pass := "pass0"
// // 	readPatch(server_url, user, pass)
// }

// func patchAPI_old(server_url string, newMap map[string]interface{}, streamName string, cmd string) {

// 	// Set new user and password for authentication
// 	user := strings.TrimSpace(newMap["user"].(string))
// 	pass := strings.TrimSpace(newMap["pass"].(string))

// 	// Set new command
// 	cmd_id := cmd + " " + streamName

// 	runOnDisconnectPatch(server_url, cmd_id)
// 	time.Sleep(150 * time.Millisecond)
// 	readPatch(server_url, user, pass)
// }

func initAPI(server_url string) {
	post_resource := "/v3/config/paths/add/"
	setting := "readUser"
	valbod := []byte(`{"readUser": "", "readPass": ""}`)
	post(server_url, post_resource+setting, valbod)
}

func patchAPI(server_url string, newMap map[string]interface{}, streamName string, cmd string) {

	// Set new user and password for authentication
	user := strings.TrimSpace(newMap["user"].(string))
	pass := strings.TrimSpace(newMap["pass"].(string))

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