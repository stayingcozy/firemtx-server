package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

func streamsListen(ctx context.Context, client *firestore.Client) (map[string]interface{}, string, error) {
	var err error
	var doc *firestore.DocumentSnapshot

	iter := client.Collection("streams").Where("status", "==", "readyToStream").Documents(ctx)
	for {
		doc, err = iter.Next()
		// if err == iterator.Done {
		// 	break
		// }
		if err != nil {
			return nil, "", err
		}
		fmt.Println(doc.Data())

		streamData := doc.Data()
		streamDocName := doc.Ref.ID

		if streamData != nil {
			return streamData, streamDocName, nil
		}
	}
}
