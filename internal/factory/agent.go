package factory

import (
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentCollection struct {
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

func (agents *AgentCollection) EncodeCollectionFlat(counterfeitHour bool, counterfeitDay bool, counterfeitMonth bool) *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     CounterfeitTimestamp(counterfeitHour, counterfeitDay, counterfeitMonth),
		TelemetryType: "agent-telemetry",
		ServerGuid:    strings.Replace(strings.Replace(gofakeit.UUID(), "-", "", -1), "-", "", -1),
		Associations:  *agents,
	}
	return &telemetry
}

func (agents *AgentCollection) EncodeAgentCollectionLookup(counterfeitHour bool, counterfeitDay bool, counterfeitMonth bool) *AgentTelemetry {
	return agents.EncodeCollectionFlat(counterfeitHour, counterfeitDay, counterfeitMonth)
}
