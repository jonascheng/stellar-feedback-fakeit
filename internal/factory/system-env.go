package factory

import (
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentSystemEnvCollection struct {
	Agents []myfakeit.AgentSystemEnv `json:"agents" xml:"agents"`
}

type AgentTelemetrySystemLookup struct {
	SystemMap map[string]myfakeit.OSInfo `json:"systemMap" xml:"systemMap"`
}

type AgentTelemetrySystemAssociationsLookup struct {
	Agents []AgentSystemEnv `json:"agents" xml:"agents"`
}

type AgentSystemEnv struct {
	Guid   string                `json:"guid" xml:"guid"`
	System string                `json:"system" xml:"system"`
	Qfe    []myfakeit.QfeInfo    `json:"qfe" xml:"qfe"`
	Volume []myfakeit.VolumeInfo `json:"volume" xml:"volume"`
	Meta   []myfakeit.MetaInfo   `json:"meta" xml:"meta"`
}

func CollectAgentSystemEnv(size int) *AgentSystemEnvCollection {
	var agents AgentSystemEnvCollection

	for i := 0; i < size; i++ {
		agent := myfakeit.SystemEnv()
		agents.Agents = append(agents.Agents, *agent)
	}
	return &agents
}

func (agents *AgentSystemEnvCollection) EncodeCollectionFlat() *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-system-environment",
		ServerGuid:    strings.Replace(gofakeit.UUID(), "-", "", -1),
		Associations:  *agents,
	}
	return &telemetry
}

func (agents *AgentSystemEnvCollection) EncodeAgentCollectionLookup() *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-system-environment",
		ServerGuid:    strings.Replace(gofakeit.UUID(), "-", "", -1),
	}

	var newAgents AgentTelemetrySystemAssociationsLookup
	lookup := make(map[myfakeit.OSInfo]string)
	encode_counter := 1
	for _, agent := range agents.Agents {
		var val string
		var ok bool
		os := myfakeit.OSInfo{
			Caption: agent.Caption,
			Version: agent.Version,
		}
		if val, ok = lookup[os]; !ok {
			// not exist
			val = strconv.Itoa(encode_counter)
			lookup[os] = val
			encode_counter++
		}

		// encode agent
		newAgent := AgentSystemEnv{
			Guid:   agent.Guid,
			System: val,
			Qfe:    agent.Qfe,
			Volume: agent.Volume,
			Meta:   agent.Meta,
		}

		newAgents.Agents = append(newAgents.Agents, newAgent)
	}
	telemetry.Associations = newAgents

	if len(lookup) > 0 {
		var reversedLookup AgentTelemetrySystemLookup
		reversedLookup.SystemMap = make(map[string]myfakeit.OSInfo)
		for key, element := range lookup {
			reversedLookup.SystemMap[element] = key
		}
		telemetry.Lookup = reversedLookup
	}

	return &telemetry
}
