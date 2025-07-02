// eventhub/receiver.go
package eventhub

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
)

func CreateReceiver(fullyQualifiedNamespace, eventHub, partitionID string) (*azeventhubs.PartitionClient, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	consumerClient, err := azeventhubs.NewConsumerClient(fullyQualifiedNamespace, eventHub, azeventhubs.DefaultConsumerGroup, cred, nil)
	if err != nil {
		return nil, err
	}

	partitionClient, err := consumerClient.NewPartitionClient(partitionID, &azeventhubs.PartitionClientOptions{
		StartPosition: azeventhubs.StartPosition{
			Earliest: to.Ptr(true),
		},
	})
	if err != nil {
		return nil, err
	}

	return partitionClient, nil
}
