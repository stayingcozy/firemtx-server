package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func updateStreamAsBroken(ctx context.Context, client *firestore.Client, docName string) error {
	_, err := client.Collection("streams").Doc(docName).Update(ctx, []firestore.Update{
		{
			Path:  "status",
			Value: "broken",
		},
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error while updating has occurred: %s", err)
	}

	return err
}

func readFromFile(filename string) (string, error) {
	// Read the contents of the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {

	// grab streamName
	args := os.Args
	// Check if at least one argument is provided
	if len(args) < 2 {
		log.Fatalln("Usage: updateStreamAsBroken arg1 was not provided")
		return
	}
	streamName := args[1]

	// Firebase Init
	path_to_serviceAccountKey := "brightpaw-d6fd6-firebase-adminsdk-qqfyk-248ef821b0.json"

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(path_to_serviceAccountKey)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// code run on "runOnDisconnect" in mediamtx.yml
	// pipe that it was using will be update to "broken"
	updateStreamAsBroken(ctx, client, streamName)
}
