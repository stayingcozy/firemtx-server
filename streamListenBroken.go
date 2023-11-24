package main

import (
	"context"

	"cloud.google.com/go/firestore"
)

func streamsListenBroken(ctx context.Context, client *firestore.Client, streamName string) error {
	var err error
	var doc *firestore.DocumentSnapshot

	iter := client.Collection("streams").Where("status", "==", "broken").Documents(ctx)
	for {
		doc, err = iter.Next()
		if err != nil {
			return err
		}
		streamDocName := doc.Ref.ID

		if streamDocName == streamName {
			return nil
		}
	}
}