package builder

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentSystemEnvCollection struct {
	Agents []myfakeit.AgentSystemEnv `json:"agents" xml:"agents"`
}

type AgentTelemetrySystemEnv struct {
	Timestamp     time.Time                             `json:"timestamp" xml:"timestamp"`
	TelemetryType string                                `json:"telemetryType" xml:"telemetryType"`
	ServerGuid    string                                `json:"serverGuid" xml:"serverGuid"`
	Lookup        map[string]AgentTelemetrySystemLookup `json:"lookup,omitempty" xml:"lookup,omitempty"`
	Associations  AgentSystemEnvCollection              `json:"associations" xml:"associations"`
}

type AgentTelemetrySystemLookup map[string]AgentSystemOperatingSystem

type AgentSystemOperatingSystem struct {
	Caption string `json:"caption" xml:"caption"`
	Version string `json:"version" xml:"version"`
}

type AgentTelemetrySystemAssociation []myfakeit.AgentSystemEnv

func CollectAgentSystemEnv(size int) AgentSystemEnvCollection {
	var agents AgentSystemEnvCollection

	for i := 0; i < size; i++ {
		var agent myfakeit.AgentSystemEnv
		err := gofakeit.Struct(&agent)
		if err != nil {
			panic(err)
		}
		agents.Agents = append(agents.Agents, agent)
	}
	return agents
}

func EncodeAgentSystemEnvCollectionFlat(agents AgentSystemEnvCollection) *AgentTelemetrySystemEnv {
	telemetry := AgentTelemetrySystemEnv{
		Timestamp:     time.Now(),
		TelemetryType: "agent-telemetry-system-environment",
		ServerGuid:    gofakeit.UUID(),
		Associations:  agents,
	}
	return &telemetry
}
