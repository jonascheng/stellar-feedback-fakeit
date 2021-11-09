package factory

import "time"

type AgentTelemetry struct {
	Timestamp     time.Time   `json:"timestamp" xml:"timestamp"`
	TelemetryType string      `json:"telemetryType" xml:"telemetryType"`
	ServerGuid    string      `json:"serverGuid" xml:"serverGuid"`
	Lookup        interface{} `json:"lookup,omitempty" xml:"lookup,omitempty"`
	Associations  interface{} `json:"associations" xml:"associations"`
}

type IAgentTelemetry interface {
	EncodeCollectionFlat() *AgentTelemetry
	EncodeAgentCollectionLookup() *AgentTelemetry
}

func NewServerCollection() *ServerCollection {
	return CollectServer()
}

func NewAgentCollection(size int) *AgentCollection {
	return CollectAgent(size)
}

func NewAgentSystemEnvCollection(size int) *AgentSystemEnvCollection {
	return CollectAgentSystemEnv(size)
}

func NewAgentSoftwareEnvCollection(size int) *AgentSoftwareEnvCollection {
	return CollectAgentSoftwareEnv(size)
}

func NewAgentCertCollection(size int) *AgentCertCollection {
	return CollectAgentCert(size)
}
