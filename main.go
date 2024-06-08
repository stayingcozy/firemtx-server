package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {

	// Production vs Dev
	production := false
	serverIP := "YOUR_SERVER_IP"

	// Init
	var app *firebase.App
	var err error
	var ctx context.Context

	// Firebase Init
	if production {
		projectID := "YOUR_FIREBASE_PROJECT_ID"

		// Use the application default credentials
		ctx := context.Background()
		conf := &firebase.Config{ProjectID: projectID}
		app, err = firebase.NewApp(ctx, conf)

	} else {
		path_to_serviceAccountKey := os.getEnv("SERVICE_KEY_PATH")

		// Use a service account
		ctx = context.Background()
		sa := option.WithCredentialsFile(path_to_serviceAccountKey)
		app, err = firebase.NewApp(ctx, nil, sa)
	}

	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// Run MediaMTX
	go runMediaMTX()            // allow firebase communication to continue
	time.Sleep(1 * time.Second) // allow mediamtx to startup before post

	// Initialize MediaMTX API
	server_url := "http://127.0.0.1:9997"
	cmd := "./brokenpipe/brokenpipe"
	initAPI(server_url)

	for {
		// Listen to streams/ in firebase
		streamMap, streamName, err := streamsListen(ctx, client)
		if err != nil {
			continue
		}

		fmt.Println("Exited streamsListen. Going to readMap and patch...")

		// Read streamMap and set new user, password, and serverIP
		user, pass, streamIPStream := readMap(streamMap)

		if streamIPStream != serverIP {
			fmt.Println("Stream IP does not match server IP. Moving on...")
			continue
		}

		// Replace with patch of patch[runOnDisconnect], patch[readUser, readPass]
		patchAPI(server_url, streamName, cmd, user, pass)

		fmt.Println("Patched. Moving on...")

		// Update stream status to "readyToReceive"
		updateStreamAsReadyToReceive(ctx, client, streamName)

		fmt.Println("ReadyToReceive. Moving on...")

		// if not "broken" stay, if broken move on
		// streamsListenBroken(ctx, client, streamName)
		streamsListenBrokenMods(ctx, client, streamName)

		fmt.Println("Exited streamsListenBroken. Looping back to streamsListen...")
	}
}
