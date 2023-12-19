package main

import (
	"context"
	"fmt"
	
	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

// listenChanges listens to a query, returning the list of document changes.
func streamsListenBrokenMods(ctx context.Context, client *firestore.Client, streamName string) error {
	// projectID = "brightpawk-d6fd6"
	// ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	// defer cancel()

	// client, err := firestore.NewClient(ctx, projectID)
	// if err != nil {
	// 		return fmt.Errorf("firestore.NewClient: %w", err)
	// }
	// defer client.Close()

	it := client.Collection("streams").Where("status", "==", "broken").Snapshots(ctx)
	for {
			snap, err := it.Next()
			// DeadlineExceeded will be returned when ctx is cancelled.
			if status.Code(err) == codes.DeadlineExceeded {
					return nil
			}
			if err != nil {
					return fmt.Errorf("Snapshots.Next: %w", err)
			}
			if snap != nil {
					for _, change := range snap.Changes {
							switch change.Kind {
							// case firestore.DocumentAdded:
							// 		fmt.Fprintf(w, "New city: %v\n", change.Doc.Data())
							case firestore.DocumentModified:
									fmt.Printf("Modified stream: %v\n", change.Doc.Data())
									if change.Doc.Ref.ID == streamName {
										return nil
									}
							case firestore.DocumentRemoved:
									fmt.Printf("Removed stream: %v\n", change.Doc.Data())
									if change.Doc.Ref.ID == streamName {
										return nil
									}
							}
					}
			}
	}
}


func streamsListenBroken(ctx context.Context, client *firestore.Client, streamName string) error {
	var err error
	var doc *firestore.DocumentSnapshot


	iter := client.Collection("streams").Where("status", "==", "broken").Documents(ctx)
	for {
		doc, err = iter.Next()
		if err != nil {
			return err
		}
		fmt.Printf("Error on streamsListenBroken: %s \n",err)
		streamDocName := doc.Ref.ID

		if streamDocName == streamName {
			return nil
		}
	}
}