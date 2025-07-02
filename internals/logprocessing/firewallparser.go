package firewallparser

import (
	"errors"
	"regexp"
	"strconv"
)

type AzFWNetworkRule struct {
	TimeGenerated string `json:"TimeGenerated"`
	SourceIP      string `json:"SourceIP"`
	DestinationIP string `json:"DestinationIP"`
	DestinationPort int `json:"DestinationPort"`
	Protocol     string `json:"Protocol"`
	Action   	string `json:"Action"`
	Rule     string `json:"Rule"`
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

func ParseAzFWNetworkRule(logLine string) (AzFWNetworkRule, error) {
	re := regexp.MustCompile(`(?P<Protocol>\w+)\srequest\sfrom\s(?P<SourceIP>\d+\.\d+\.\d+\.\d+):(?P<SourcePort>\d+)\sto\s(?P<DestinationIP>\d+\.\d+\.\d+\.\d+):(?P<DestinationPort>\d+)\.\sAction:\s(?P<Action>\w+)`)
	match := re.FindStringSubmatch(logLine)

	if match == nil {
		return AzFWNetworkRule{}, errors.New("log line does not match expected format")
	}

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