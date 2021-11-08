package factory

import (
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentSoftwareEnvCollection struct {
	Agents []myfakeit.AgentSoftwareEnv `json:"agents" xml:"agents"`
}

type AgentSoftwareEnvCollectionLookup struct {
	Agents []AgentSoftwareEnv `json:"agents" xml:"agents"`
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
	Guid     string   `json:"guid" xml:"guid"`
	Software []string `json:"software" xml:"software"`
}

func CollectAgentSoftwareEnv(size int) *AgentSoftwareEnvCollection {
	var agents AgentSoftwareEnvCollection

	for i := 0; i < size; i++ {
		agent := myfakeit.SoftwareEnv()
		agents.Agents = append(agents.Agents, *agent)
	}
	return &agents
}

func (agents *AgentSoftwareEnvCollection) EncodeAgentCollectionFlat() *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-software-environment",
		ServerGuid:    strings.Replace(strings.Replace(gofakeit.UUID(), "-", "", -1), "-", "", -1),
		Associations:  *agents,
	}
	return &telemetry
}

func (agents *AgentSoftwareEnvCollection) EncodeAgentCollectionLookup() *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-software-environment",
		ServerGuid:    strings.Replace(gofakeit.UUID(), "-", "", -1),
	}

	var newAgents AgentSoftwareEnvCollectionLookup
	lookup := make(map[AgentSoftwareApplication]string)
	encode_counter := 1
	for _, agent := range agents.Agents {
		var softwareList []string
		var val string
		var ok bool
		for _, app := range agent.App {
			software := AgentSoftwareApplication{
				Caption:         app.Caption,
				Version:         app.Version,
				InstallLocation: app.InstallLocation,
			}
			if val, ok = lookup[software]; !ok {
				// not exist
				val = strconv.Itoa(encode_counter)
				lookup[software] = val
				encode_counter++
			}
			softwareList = append(softwareList, val)
		}

		// encode agent
		newAgent := AgentSoftwareEnv{
			Guid:     agent.Guid,
			Software: softwareList,
		}

		newAgents.Agents = append(newAgents.Agents, newAgent)
	}
	telemetry.Associations = newAgents

	if len(lookup) > 0 {
		var reversedLookup AgentTelemetrySoftwareLookup
		reversedLookup.SoftwareMap = make(map[string]AgentSoftwareApplication)
		for key, element := range lookup {
			reversedLookup.SoftwareMap[element] = key
		}
		telemetry.Lookup = reversedLookup
	}

	return &telemetry
}
