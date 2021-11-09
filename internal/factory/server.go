package factory

import (
	"time"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type ServerCollection struct {
	Server *myfakeit.Server `json:"server" xml:"server"`
}

type ServerTelemetryAssociations myfakeit.Server

func CollectServer() *ServerCollection {
	var server ServerCollection

	server.Server = myfakeit.ServerInfo()

	return &server
}

func (server *ServerCollection) EncodeCollectionFlat() *AgentTelemetry {
	telemetry := AgentTelemetry{
		Timestamp:     time.Now().UTC(),
		TelemetryType: "server-telemetry",
		ServerGuid:    server.Server.Guid,
		Associations:  *server.Server,
	}
	return &telemetry
}

func (server *ServerCollection) EncodeAgentCollectionLookup() *AgentTelemetry {
	return server.EncodeCollectionFlat()
}
