package main

import (
	"context"
	"log"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

func main() {
	log.SetOutput(os.Stdout)
	ctx := context.Background()

	p, err := cloudevents.NewHTTP(cloudevents.WithPort(8080), cloudevents.WithPath("/"))
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}
	c, err := cloudevents.NewClient(p,
		cloudevents.WithUUIDs(),
		cloudevents.WithTimeNow(),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err.Error())
	}

	log.Println("listening on: localhost:8080")
	if err := c.StartReceiver(ctx, eventReceiver); err != nil {
		log.Fatalf("failed to start receiver: %s", err.Error())
	}

	<-ctx.Done()
}

func eventReceiver(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, error) {
	log.Printf("Got Event Context: %+v\n", event.Context)
	log.Printf("----------------------------\n")

	responseEvent := cloudevents.NewEvent()
	responseEvent.SetID(uuid.New().String())
	responseEvent.SetSource("oger-at-door")
	responseEvent.SetType("door-events")
	_ = responseEvent.SetData(cloudevents.ApplicationJSON, `{ "event": "The door opens, you are allowed to pass!" }`)

	return &responseEvent, nil
}
