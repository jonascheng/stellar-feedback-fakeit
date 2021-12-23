package factory

import (
	"strings"

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

func (threats *ThreatCollection) EncodeCollectionFlat(counterfeitHour bool, counterfeitDay bool, counterfeitMonth bool) *AgentTelemetry {
	var associatedThreats ThreatTelemetryAssociations
	for _, agent := range threats.Agents {
		if agent.AppExecBlocked != nil {
			associatedThreats.AppExecBlocked = append(associatedThreats.AppExecBlocked, agent.AppExecBlocked...)
		}
		if agent.FileScanBlocked != nil {
			associatedThreats.FileScanBlocked = append(associatedThreats.FileScanBlocked, agent.FileScanBlocked...)
		}
		if agent.SuspiciousExecBlocked != nil {
			associatedThreats.SuspiciousExecBlocked = append(associatedThreats.SuspiciousExecBlocked, agent.SuspiciousExecBlocked...)
		}
		if agent.OBADBlocked != nil {
			associatedThreats.OBADBlocked = append(associatedThreats.OBADBlocked, agent.OBADBlocked...)
		}
		if agent.NonWhitelistingBlocked != nil {
			associatedThreats.NonWhitelistingBlocked = append(associatedThreats.NonWhitelistingBlocked, agent.NonWhitelistingBlocked...)
		}
		if agent.ADCBlocked != nil {
			associatedThreats.ADCBlocked = append(associatedThreats.ADCBlocked, agent.ADCBlocked...)
		}
	}

	telemetry := AgentTelemetry{
		Timestamp:     CounterfeitTimestamp(counterfeitHour, counterfeitDay, counterfeitMonth),
		TelemetryType: "agent-telemetry-threat",
		ServerGuid:    strings.Replace(strings.Replace(gofakeit.UUID(), "-", "", -1), "-", "", -1),
		Associations:  associatedThreats,
	}

	return &telemetry
}

func (threats *ThreatCollection) EncodeAgentCollectionLookup(counterfeitHour bool, counterfeitDay bool, counterfeitMonth bool) *AgentTelemetry {
	return threats.EncodeCollectionFlat(counterfeitHour, counterfeitDay, counterfeitMonth)
}
