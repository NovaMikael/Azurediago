package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NovaMikael/Azurediago/internals/eventhub"
	"github.com/NovaMikael/Azurediago/internals/logprocessing"
)

func main() {
	receiver, err := eventhub.CreateReceiver("eventsochgo.servicebus.windows.net", "firewall", "0")
	if err != nil {
		log.Fatalf("Failed to create receiver: %v", err)
	}
	defer receiver.Close(context.TODO())

	events, err := receiver.ReceiveEvents(context.Background(), 10, nil)
	if err != nil {
		log.Fatalf("Failed to receive events: %v", err)
	}

	for _, evt := range events {
		line := string(evt.Body)
		fmt.Printf("Raw event: %s\n", line)

		netRule, err := logprocessing.ParseAzFWNetworkRule(line)
		if err != nil {
			log.Printf("network rule parse error: %v", err)
		} else {
			fmt.Printf("network rule: %+v\n", netRule)
		}

		appRule, err := logprocessing.ParseAzFWApplicationRule(line)
		if err != nil {
			log.Printf("application rule parse error: %v", err)
		} else {
			fmt.Printf("app rule: %+v\n", appRule)
		}
	}
}
