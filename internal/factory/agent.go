package factory

import (
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentCollection struct {
	Agents []myfakeit.Agent `json:"agents" xml:"agents"`
}

type AgentTelemetryAssociations struct {
	Agents []myfakeit.Agent `json:"agents" xml:"agents"`
}

func CollectAgent(size int) *AgentCollection {
	var agents AgentCollection

	for i := 0; i < size; i++ {
		agent := myfakeit.AgentInfo()
		agents.Agents = append(agents.Agents, *agent)
	}
	return &agents
}

func (agents *AgentCollection) EncodeAgentCollectionFlat() *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry",
		ServerGuid:    strings.Replace(strings.Replace(gofakeit.UUID(), "-", "", -1), "-", "", -1),
		Associations:  *agents,
	}
	return &telemetry
}

func (agents *AgentCollection) EncodeAgentCollectionLookup() *AgentTelemetry {
	return agents.EncodeAgentCollectionFlat()
}
