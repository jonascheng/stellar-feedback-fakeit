package builder

import (
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentSoftwareEnvCollection struct {
	Agents []myfakeit.AgentSoftwareEnv `json:"agents" xml:"agents"`
}

type AgentTelemetry struct {
	Timestamp     time.Time   `json:"timestamp" xml:"timestamp"`
	TelemetryType string      `json:"telemetryType" xml:"telemetryType"`
	ServerGuid    string      `json:"serverGuid" xml:"serverGuid"`
	Lookup        interface{} `json:"lookup,omitempty" xml:"lookup,omitempty"`
	Associations  interface{} `json:"associations" xml:"associations"`
}

type AgentTelemetrySoftwareLookup struct {
	SoftwareMap map[string]AgentSoftwareApplication `json:"softwareMap" xml:"softwareMap"`
}

type AgentSoftwareApplication struct {
	Caption         string `json:"caption" xml:"caption"`
	Version         string `json:"version" xml:"version"`
	InstallLocation string `json:"installLocation" xml:"installLocation"`
}

type AgentTelemetrySoftwareAssociations struct {
	Agents []myfakeit.AgentSoftwareEnv `json:"agents" xml:"agents"`
}

type AgentSoftwareEnv struct {
	Guid     string `json:"guid" xml:"guid"`
	Software string `json:"software" xml:"software"`
}

func CollectAgentSoftwareEnv(size int) *AgentSoftwareEnvCollection {
	var agents AgentSoftwareEnvCollection

	for i := 0; i < size; i++ {
		agent := myfakeit.SoftwareEnv()
		agents.Agents = append(agents.Agents, *agent)
	}
	return &agents
}

func EncodeAgentCollectionFlat(agents *AgentSoftwareEnvCollection) *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-software-environment",
		ServerGuid:    strings.Replace(gofakeit.UUID(), "-", "", -1),
		Associations:  *agents,
	}
	return &telemetry
}
