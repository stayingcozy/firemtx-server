package main

import (
	"bytes"
	"fmt"
	"net/http"
	"io"
)

func patch(url string, resource string, jsonBody []byte) {

	url = url + resource

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPatch, url, bodyReader)
	if err != nil {  
		fmt.Printf("client: could not create request: %s\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}

	// fmt.Printf("client: got response!\n")
	fmt.Printf("patch client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}

func post(url string, resource string, jsonBody []byte) error {

	url = url + resource

	// jsonBody := []byte(`{"readUser": "user0", "readPass": "pass0"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	// req, err := http.NewRequest(http.MethodPatch, url, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return err
	}

	// fmt.Printf("client: got response!\n")
	fmt.Printf("post client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return err
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return nil
}