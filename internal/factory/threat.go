package factory

import (
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type ThreatCollection struct {
	Agents []myfakeit.Threat `json:"agents" xml:"agents"`
}

type ThreatTelemetryAssociations myfakeit.Threat

func CollectThreat(size int) *ThreatCollection {
	var threats ThreatCollection

	for i := 0; i < size; i++ {
		agent := myfakeit.ThreatInfo(0)
		threats.Agents = append(threats.Agents, *agent)
	}
	
	return &threats
}

func (threats *ThreatCollection) EncodeCollectionFlat() *AgentTelemetry {
	var associatedThreats ThreatTelemetryAssociations
	for _, agent := range threats.Agents {
		if agent.AppExecBlocked != nil {
			associatedThreats.AppExecBlocked = append(associatedThreats.AppExecBlocked, agent.AppExecBlocked...)
		}
	}

	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-threat",
		ServerGuid:    strings.Replace(strings.Replace(gofakeit.UUID(), "-", "", -1), "-", "", -1),
		Associations:  associatedThreats,
	}

	return &telemetry
}

func (threats *ThreatCollection) EncodeAgentCollectionLookup() *AgentTelemetry {
	return threats.EncodeCollectionFlat()
}
