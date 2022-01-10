package factory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
	"github.com/stretchr/testify/assert"
)

func TestNewServerCollection(t *testing.T) {
	gofakeit.Seed(11)

	server := NewServerCollection()

	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", server.Server.Guid)
	assert.Equal(t, "1.0.1000", server.Server.Version)
	assert.Equal(t, []string{"PTHYM3Q2DIFHYM0E490JPTGN71CBMJIX", "28UYLHB7PV5V3YDCXJGPA9091PA58EU5"}, server.Server.AC)
	assert.Equal(t, 4, len(server.Server.EnabledFeature))
	assert.Equal(t, "forward-syslog", server.Server.EnabledFeature[0])
	assert.Equal(t, 5, len(*server.Server.Meta))
}

func TestEncodeServerCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	server := NewServerCollection()

	telemetry := server.EncodeCollectionFlat(false, false, false)
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "server-telemetry", telemetry.TelemetryType)

	associatedServer := telemetry.Associations.(myfakeit.Server)
	assert.Equal(t, telemetry.ServerGuid, associatedServer.Guid)
	assert.Equal(t, "1.0.1000", associatedServer.Version)
	assert.Equal(t, []string{"PTHYM3Q2DIFHYM0E490JPTGN71CBMJIX", "28UYLHB7PV5V3YDCXJGPA9091PA58EU5"}, associatedServer.AC)
	assert.Equal(t, 4, len(associatedServer.EnabledFeature))
	assert.Equal(t, "forward-syslog", associatedServer.EnabledFeature[0])
	assert.Equal(t, 5, len(*associatedServer.Meta))
}

// identical to EncodeCollectionFlat
func TestEncodeServerCollectionLookup(t *testing.T) {
	gofakeit.Seed(11)

	server := NewServerCollection()

	telemetry := server.EncodeCollectionFlat(false, false, false)
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "server-telemetry", telemetry.TelemetryType)

	associatedServer := telemetry.Associations.(myfakeit.Server)
	assert.Equal(t, telemetry.ServerGuid, associatedServer.Guid)
	assert.Equal(t, "1.0.1000", associatedServer.Version)
	assert.Equal(t, []string{"PTHYM3Q2DIFHYM0E490JPTGN71CBMJIX", "28UYLHB7PV5V3YDCXJGPA9091PA58EU5"}, associatedServer.AC)
	assert.Equal(t, 4, len(associatedServer.EnabledFeature))
	assert.Equal(t, "forward-syslog", associatedServer.EnabledFeature[0])
	assert.Equal(t, 5, len(*associatedServer.Meta))
}
