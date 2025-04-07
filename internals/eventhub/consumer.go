package eventhub

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
)

func CreateReceiver(fullyQualifiedNamespace, eventHub, partitionID string) {

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to create Azure credential: %v", err)
	}

	consumerClient, err := azeventhubs.NewConsumerClient(fullyQualifiedNamespace, eventHub, azeventhubs.DefaultConsumerGroup, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer client: %v", err)
	}

	defer consumerClient.Close(context.TODO())

	partitionClient, err := consumerClient.NewPartitionClient(partitionID, &azeventhubs.PartitionClientOptions{
		StartPosition: azeventhubs.StartPosition{
			Earliest: to.Ptr(true),
		},
	})

	if err != nil {
		panic(err)
	}

	defer partitionClient.Close(context.TODO())

}
