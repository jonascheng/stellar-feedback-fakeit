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
	assert.Equal(t, "TES1ZAEXZ", server.Server.SKU)
	assert.Equal(t, "THYM3Q2DIFHYM0E490JPTGN71CBMJIX2", server.Server.AC)
	assert.Equal(t, 5, len(server.Server.Volume))
	assert.Equal(t, "/dev/mapper/vg_redis-lvm_redis", server.Server.Volume[0].Source)
	assert.Equal(t, "3464667796943", server.Server.Volume[0].Total)
	assert.Equal(t, "2275832852", server.Server.Volume[0].Free)
	assert.Equal(t, "ext4", server.Server.Volume[0].Type)
	assert.Equal(t, 4, len(server.Server.EnabledFeature))
	assert.Equal(t, "schedule-update", server.Server.EnabledFeature[0])
	assert.Equal(t, 5, len(server.Server.Meta))
}

func TestEncodeServerCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	server := NewServerCollection()

	telemetry := server.EncodeAgentCollectionFlat()
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "server-telemetry", telemetry.TelemetryType)

	associatedServer := telemetry.Associations.(myfakeit.Server)
	assert.Equal(t, telemetry.ServerGuid, associatedServer.Guid)
	assert.Equal(t, "1.0.1000", associatedServer.Version)
	assert.Equal(t, "TES1ZAEXZ", associatedServer.SKU)
	assert.Equal(t, "THYM3Q2DIFHYM0E490JPTGN71CBMJIX2", associatedServer.AC)
	assert.Equal(t, 5, len(associatedServer.Volume))
	assert.Equal(t, "/dev/mapper/vg_redis-lvm_redis", associatedServer.Volume[0].Source)
	assert.Equal(t, "3464667796943", associatedServer.Volume[0].Total)
	assert.Equal(t, "2275832852", associatedServer.Volume[0].Free)
	assert.Equal(t, "ext4", associatedServer.Volume[0].Type)
	assert.Equal(t, 4, len(associatedServer.EnabledFeature))
	assert.Equal(t, "schedule-update", associatedServer.EnabledFeature[0])
	assert.Equal(t, 5, len(associatedServer.Meta))
}

// identical to EncodeAgentCollectionFlat
func TestEncodeServerCollectionLookup(t *testing.T) {
	gofakeit.Seed(11)

	server := NewServerCollection()

	telemetry := server.EncodeAgentCollectionFlat()
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "server-telemetry", telemetry.TelemetryType)

	associatedServer := telemetry.Associations.(myfakeit.Server)
	assert.Equal(t, telemetry.ServerGuid, associatedServer.Guid)
	assert.Equal(t, "1.0.1000", associatedServer.Version)
	assert.Equal(t, "TES1ZAEXZ", associatedServer.SKU)
	assert.Equal(t, "THYM3Q2DIFHYM0E490JPTGN71CBMJIX2", associatedServer.AC)
	assert.Equal(t, 5, len(associatedServer.Volume))
	assert.Equal(t, "/dev/mapper/vg_redis-lvm_redis", associatedServer.Volume[0].Source)
	assert.Equal(t, "3464667796943", associatedServer.Volume[0].Total)
	assert.Equal(t, "2275832852", associatedServer.Volume[0].Free)
	assert.Equal(t, "ext4", associatedServer.Volume[0].Type)
	assert.Equal(t, 4, len(associatedServer.EnabledFeature))
	assert.Equal(t, "schedule-update", associatedServer.EnabledFeature[0])
	assert.Equal(t, 5, len(associatedServer.Meta))
}
