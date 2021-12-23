package factory

import (
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type ServerCollection struct {
	Server *myfakeit.Server `json:"server" xml:"server"`
}

func CollectServer() *ServerCollection {
	var server ServerCollection

	server.Server = myfakeit.ServerInfo()

	return &server
}

func (server *ServerCollection) EncodeCollectionFlat(counterfeitHour bool, counterfeitDay bool, counterfeitMonth bool) *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     CounterfeitTimestamp(counterfeitHour, counterfeitDay, counterfeitMonth),
		TelemetryType: "server-telemetry",
		ServerGuid:    server.Server.Guid,
		Associations:  *server.Server,
	}
	return &telemetry
}

func (server *ServerCollection) EncodeAgentCollectionLookup(counterfeitHour bool, counterfeitDay bool, counterfeitMonth bool) *AgentTelemetry {
	return server.EncodeCollectionFlat(counterfeitHour, counterfeitDay, counterfeitMonth)
}
