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

type AgentTelemetrySystemEnvFlat struct {
	Timestamp     time.Time                `json:"timestamp" xml:"timestamp"`
	TelemetryType string                   `json:"telemetryType" xml:"telemetryType"`
	ServerGuid    string                   `json:"serverGuid" xml:"serverGuid"`
	Associations  AgentSystemEnvCollection `json:"associations" xml:"associations"`
}

type AgentTelemetrySystemEnvLookup struct {
	Timestamp     time.Time                        `json:"timestamp" xml:"timestamp"`
	TelemetryType string                           `json:"telemetryType" xml:"telemetryType"`
	ServerGuid    string                           `json:"serverGuid" xml:"serverGuid"`
	Lookup        AgentTelemetrySystemLookup       `json:"lookup" xml:"lookup"`
	Associations  AgentTelemetrySystemAssociations `json:"associations" xml:"associations"`
}

type AgentTelemetrySystemLookup struct {
	SystemMap map[string]AgentSystemOperatingSystem `json:"systemMap" xml:"systemMap"`
}

type AgentSystemOperatingSystem struct {
	Caption string `json:"caption" xml:"caption"`
	Version string `json:"version" xml:"version"`
}

type AgentTelemetrySystemAssociations struct {
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

func EncodeAgentSystemEnvCollectionFlat(agents *AgentSystemEnvCollection) *AgentTelemetrySystemEnvFlat {
	telemetry := AgentTelemetrySystemEnvFlat{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-system-environment",
		ServerGuid:    strings.Replace(gofakeit.UUID(), "-", "", -1),
		Associations:  *agents,
	}
	return &telemetry
}

func EncodeAgentSystemEnvCollectionLookup(agents *AgentSystemEnvCollection) *AgentTelemetrySystemEnvLookup {
	telemetry := AgentTelemetrySystemEnvLookup{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-system-environment",
		ServerGuid:    strings.Replace(gofakeit.UUID(), "-", "", -1),
	}

	lookup := make(map[AgentSystemOperatingSystem]string)
	encode_counter := 1
	for _, agent := range agents.Agents {
		var val string
		var ok bool
		os := AgentSystemOperatingSystem{
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

		telemetry.Associations.Agents = append(telemetry.Associations.Agents, newAgent)
	}

	if len(lookup) > 0 {
		telemetry.Lookup.SystemMap = make(map[string]AgentSystemOperatingSystem)
		for key, element := range lookup {
			telemetry.Lookup.SystemMap[element] = key
		}
	}

	return &telemetry
}
