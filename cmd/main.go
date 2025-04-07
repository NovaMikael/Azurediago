package main

import "time"

type LogEntry struct {
	Time            time.Time `json:"time"`
	Category        string    `json:"category"`
	OperationName   string    `json:"operationName"`
	ResultType      string    `json:"resultType"`
	CorrelationID   string    `json:"correlationId"`
	CallerIPAddress string    `json:"callerIpAddress"`
	Identity        struct {
		Claim map[string]string `json:"claim"`
	} `json:"identity"`
	Properties struct {
		ID             string `json:"id"`
		ClientInfo     string `json:"clientInfo"`
		RequestURI     string `json:"requestUri"`
		HTTPStatusCode int    `json:"httpStatusCode"`
		Properties     struct {
			SKU struct {
				Family   string      `json:"Family"`
				Name     string      `json:"Name"`
				Capacity interface{} `json:"Capacity"`
			} `json:"sku"`
			TenantID                     string      `json:"tenantId"`
			NetworkACLs                  interface{} `json:"networkAcls"`
			EnabledForDeployment         int         `json:"enabledForDeployment"`
			EnabledForDiskEncryption     int         `json:"enabledForDiskEncryption"`
			EnabledForTemplateDeployment int         `json:"enabledForTemplateDeployment"`
			EnableSoftDelete             int         `json:"enableSoftDelete"`
			SoftDeleteRetentionInDays    int         `json:"softDeleteRetentionInDays"`
			EnableRbacAuthorization      int         `json:"enableRbacAuthorization"`
			EnablePurgeProtection        interface{} `json:"enablePurgeProtection"`
		} `json:"properties"`
	} `json:"properties"`
	ResourceID       string `json:"resourceId"`
	OperationVersion string `json:"operationVersion"`
	ResultSignature  string `json:"resultSignature"`
	DurationMs       string `json:"durationMs"`
}

type EventHubMessage struct {
	Records               []LogEntry `json:"records"`
	EventProcessedUtcTime time.Time  `json:"EventProcessedUtcTime"`
	PartitionID           int        `json:"PartitionId"`
	EventEnqueuedUtcTime  time.Time  `json:"EventEnqueuedUtcTime"`
}
