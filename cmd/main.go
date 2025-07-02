package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/NovaMikael/Azurediago/internals/eventhub"
	"github.com/NovaMIkael/Azurediago/internals/logprocessing"
)

func main() {
	receiver, err := eventhub.CreateReceiver("eventsochgo.servicebus.windows.net", "storageaccount", "0")
	if err != nil {
		log.Fatalf("Failed to create receiver: %v", err)
	}
	defer receiver.Close(context.TODO())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	events, err := receiver.ReceiveEvents(ctx, 10, nil)
	if err != nil {
		log.Fatalf("Failed to receive events: %v", err)
	}

	for _, evt := range events {
		fmt.Printf("Event Body: %s\n", string(evt.Body))
	}
}
