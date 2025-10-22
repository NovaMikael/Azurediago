package firewallparser

import (
	"errors"
	"regexp"
	"strconv"
)

type AzFWNetworkRule struct {
	TimeGenerated   string `json:"time"`
	SourceIP        string `json:"SourceIp"`
	DestinationIP   string `json:"DestinationIp"`
	DestinationPort int    `json:"DestinationPort"`
	Protocol        string `json:"Protocol"`
	Action          string `json:"Action"`
}

type AzFWApplicationRule struct {
	TimeGenerated string `json:"TimeGenerated"`
	SourceIP      string `json:"SourceIP"`
	Fqdn		string `json:"Fqdn"`
	DestinationPort int `json:"DestinationPort"`
	Protocol     string `json:"Protocol"`
	Action   	string `json:"Action"`
	Rule     string `json:"Rule"`
}

{ "time": "2025-09-22T11:32:59.124368+00:00", "resourceId": "/SUBSCRIPTIONS/0A347908-8617-40DB-8011-0F828996D75F/RESOURCEGROUPS/EVENTSMEDGO/PROVIDERS/MICROSOFT.NETWORK/AZUREFIREWALLS/BSICFIREWALL", "properties": {"Protocol":"ICMP Type=8","SourceIp":"10.0.2.4","SourcePort":0,"DestinationIp":"8.8.8.8","DestinationPort":0,"Action":"Deny","Policy":"","RuleCollectionGroup":"","RuleCollection":"","Rule":"","ActionReason":"Default Action"}, "category": "AZFWNetworkRule"}

func ParseAzFWNetworkRule(logLine string) (AzFWNetworkRule, error) {


	result := AzFWNetworkRule{
		Protocol:       match[1],
		SourceIP:       match[2],
		DestinationIP:  match[4],
		DestinationPort: atoi(match[5]),
		Action:         match[6],
	}

	return result, nil
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
func ParseAzFWApplicationRule(logLine string) (AzFWApplicationRule, error) {
	re := regexp.MustCompile(`(?P<Protocol>\w+)\srequest\sfrom\s(?P<SourceIP>\d+\.\d+\.\d+\.\d+):(?P<SourcePort>\d+)\sto\s(?P<Fqdn>[\w.-]+):(?P<DestinationPort>\d+)\.\sAction:\s(?P<Action>\w+)`)
	match := re.FindStringSubmatch(logLine)
	if match == nil {
		return AzFWApplicationRule{}, errors.New("log line does not match expected format")
	}
	result := AzFWApplicationRule{
		Protocol:       match[1],
		SourceIP:       match[2],
		Fqdn:          match[3],
		DestinationPort: atoi(match[4]),
		Action:         match[5],
	}
	return result, nil
}