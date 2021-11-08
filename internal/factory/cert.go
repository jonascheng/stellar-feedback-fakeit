package factory

import (
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentCertCollection struct {
	Agents []myfakeit.AgentCert `json:"agents" xml:"agents"`
}

type AgentTelemetryCertLookup struct {
	CertMap map[string]myfakeit.CertInfo `json:"certMap" xml:"certMap"`
}

type AgentTelemetryCertAssociations struct {
	Agents []myfakeit.AgentCert `json:"agents" xml:"agents"`
}

type AgentTelemetryCertAssociationsLookup struct {
	Agents []AgentCert `json:"agents" xml:"agents"`
}

type AgentCert struct {
	Guid string   `json:"guid" xml:"guid"`
	Cert []string `json:"cert" xml:"cert"`
}

func CollectAgentCert(size int) *AgentCertCollection {
	var agents AgentCertCollection

	for i := 0; i < size; i++ {
		agent := myfakeit.Cert()
		agents.Agents = append(agents.Agents, *agent)
	}
	return &agents
}

func (agents *AgentCertCollection) EncodeAgentCollectionFlat() *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-cert",
		ServerGuid:    strings.Replace(strings.Replace(gofakeit.UUID(), "-", "", -1), "-", "", -1),
		Associations:  *agents,
	}
	return &telemetry
}

func (agents *AgentCertCollection) EncodeAgentCollectionLookup() *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "agent-telemetry-cert",
		ServerGuid:    strings.Replace(gofakeit.UUID(), "-", "", -1),
	}

	var newAgents AgentTelemetryCertAssociationsLookup
	lookup := make(map[myfakeit.CertInfo]string)
	encode_counter := 1
	for _, agent := range agents.Agents {
		var certList []string
		var val string
		var ok bool
		for _, cert := range agent.Cert {
			if val, ok = lookup[cert]; !ok {
				// not exist
				val = strconv.Itoa(encode_counter)
				lookup[cert] = val
				encode_counter++
			}
			certList = append(certList, val)
		}

		// encode agent
		newAgent := AgentCert{
			Guid: agent.Guid,
			Cert: certList,
		}

		newAgents.Agents = append(newAgents.Agents, newAgent)
	}
	telemetry.Associations = newAgents

	if len(lookup) > 0 {
		var reversedLookup AgentTelemetryCertLookup
		reversedLookup.CertMap = make(map[string]myfakeit.CertInfo)
		for key, element := range lookup {
			reversedLookup.CertMap[element] = key
		}
		telemetry.Lookup = reversedLookup
	}

	return &telemetry
}
