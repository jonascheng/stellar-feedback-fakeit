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

type AgentTelemetrySoftwareLookup struct {
	AppMap map[string]myfakeit.AppInfo `json:"appMap" xml:"appMap"`
}

type AgentTelemetrySoftwareAssociationsLookup struct {
	Agents []AgentSoftwareEnv `json:"agents" xml:"agents"`
}

type AgentSoftwareEnv struct {
	Guid string   `json:"guid" xml:"guid"`
	App  []string `json:"app" xml:"app"`
}

func CollectAgentSoftwareEnv(size int) *AgentSoftwareEnvCollection {
	var agents AgentSoftwareEnvCollection

	for i := 0; i < size; i++ {
		agent := myfakeit.SoftwareEnv()
		agents.Agents = append(agents.Agents, *agent)
	}
	return &agents
}

func (agents *AgentSoftwareEnvCollection) EncodeCollectionFlat() *AgentTelemetry {
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

	var newAgents AgentTelemetrySoftwareAssociationsLookup
	lookup := make(map[myfakeit.AppInfo]string)
	encode_counter := 1
	for _, agent := range agents.Agents {
		var appList []string
		var val string
		var ok bool
		for _, app := range agent.App {
			if val, ok = lookup[app]; !ok {
				// not exist
				val = strconv.Itoa(encode_counter)
				lookup[app] = val
				encode_counter++
			}
			appList = append(appList, val)
		}

		// encode agent
		newAgent := AgentSoftwareEnv{
			Guid: agent.Guid,
			App:  appList,
		}

		newAgents.Agents = append(newAgents.Agents, newAgent)
	}
	telemetry.Associations = newAgents

	if len(lookup) > 0 {
		var reversedLookup AgentTelemetrySoftwareLookup
		reversedLookup.AppMap = make(map[string]myfakeit.AppInfo)
		for key, element := range lookup {
			reversedLookup.AppMap[element] = key
		}
		telemetry.Lookup = reversedLookup
	}

	return &telemetry
}
