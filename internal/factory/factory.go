package factory

import "time"

type AgentTelemetry struct {
	Timestamp     time.Time   `json:"timestamp" xml:"timestamp"`
	TelemetryType string      `json:"telemetryType" xml:"telemetryType"`
	ServerGuid    string      `json:"serverGuid" xml:"serverGuid"`
	Lookup        interface{} `json:"lookup,omitempty" xml:"lookup,omitempty"`
	Associations  interface{} `json:"associations" xml:"associations"`
}

type AgentTelemetryEncoder interface {
	EncodeAgentCollectionFlat() *AgentTelemetry
	EncodeAgentCollectionLookup() *AgentTelemetry
}

func NewAgentSoftwareEnvCollection(size int) *AgentSoftwareEnvCollection {
	return CollectAgentSoftwareEnv(size)
}
