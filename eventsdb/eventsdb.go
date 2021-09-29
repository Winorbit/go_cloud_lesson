package eventsdb

import (
	"log"
	"context"

	"errors"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"cloud.google.com/go/firestore"
)

type Event struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Location string `json:"location"`
    When   string `json:"when"`
}

var Events []Event

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "brave-arcadia-325423"
	// !! ^^^^^^^^^^^^ !!!

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func GetEvents() []Event {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var docos []Event

	iter := client.Collection("events").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var doca Event
		doc.DataTo(&doca)
		docos = append(docos, doca)
	}
  return docos
}

func GetEventbyID(id string) (Event, error) {
	var event Event
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()
	doc, err := client.Collection("events").Doc(id).Get(ctx)
	if err != nil {
        return Event{}, errors.New("not found")
	}

	doc.DataTo(&event)
	return event, nil
}

func AddEvent(event Event) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()
	newID := uuid.New().String()
	event.ID = newID

	_, err := client.Collection("events").Doc(newID).Set(ctx, event)

	if err != nil {
        log.Fatalf("Failed adding alovelace: %v", err)
	}
}

func UpdateEvent(event Event) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, err := client.Collection("events").Doc(event.ID).Set(ctx, event)

	if err != nil {
        log.Fatalf("Failed adding alovelace: %v", err)
	}
}


func DeleteEvent(id string) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, err := client.Collection("events").Doc(id).Delete(ctx)
	if err != nil {
        // Handle any errors in an appropriate way, such as returning them.
        log.Printf("An error has occurred: %s", err)
	}
}