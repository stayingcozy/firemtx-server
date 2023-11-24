package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func updateStreamAsReadyToReceive(ctx context.Context, client *firestore.Client, docName string) error {
	_, err := client.Collection("streams").Doc(docName).Update(ctx, []firestore.Update{
		{
			Path:  "status",
			Value: "readyToReceive",
		},
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error while updating has occurred: %s", err)
	}

	return err
}
