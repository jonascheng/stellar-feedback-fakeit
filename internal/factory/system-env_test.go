package factory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
	"github.com/stretchr/testify/assert"
)

func TestCollectAgentSystemEnv(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentSystemEnvCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	agent := agents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, "Microsoft Windows 10 Pro for Workstations", agent.Caption)
	assert.Equal(t, "10.0.16299", agent.Version)
	assert.Equal(t, 5, len(agent.Qfe))
	assert.Equal(t, myfakeit.QfeInfo{HotfixId: "KB5994899", InstalledOn: "21/07/2021"}, agent.Qfe[0])
	assert.Equal(t, 2, len(agent.Volume))
	assert.Equal(t, myfakeit.VolumeInfo{Total: "2333396909501", Free: "2888082257", Type: "NTFS"}, agent.Volume[0])
	assert.Equal(t, 7, len(agent.Meta))
	assert.Equal(t, myfakeit.MetaInfo{"cpuCaption": "Intel64 Family 6 Model 167 Stepping 1"}, agent.Meta[0])
}

func TestEncodeAgentCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentSystemEnvCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionFlat()
	assert.Equal(t, "483dde0df92b43109a9b9ddd66ec91f0", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-system-environment", telemetry.TelemetryType)

	associatedAgents := telemetry.Associations.(AgentSystemEnvCollection)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := associatedAgents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, "Microsoft Windows 10 Pro for Workstations", agent.Caption)
	assert.Equal(t, "10.0.16299", agent.Version)
	assert.Equal(t, 5, len(agent.Qfe))
	assert.Equal(t, myfakeit.QfeInfo{HotfixId: "KB5994899", InstalledOn: "21/07/2021"}, agent.Qfe[0])
	assert.Equal(t, 2, len(agent.Volume))
	assert.Equal(t, myfakeit.VolumeInfo{Total: "2333396909501", Free: "2888082257", Type: "NTFS"}, agent.Volume[0])
	assert.Equal(t, 7, len(agent.Meta))
	assert.Equal(t, myfakeit.MetaInfo{"cpuCaption": "Intel64 Family 6 Model 167 Stepping 1"}, agent.Meta[0])
}

func TestEncodeAgentSystemEnvCollectionLookup(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentSystemEnvCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "483dde0df92b43109a9b9ddd66ec91f0", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-system-environment", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetrySystemLookup)
	assert.Greater(t, len(lookup.SystemMap), 0)

	os, ok := lookup.SystemMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "Microsoft Windows 10 Pro for Workstations", os.Caption)
	assert.Equal(t, "10.0.16299", os.Version)

	associatedAgents := telemetry.Associations.(AgentTelemetrySystemAssociationsLookup)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := associatedAgents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, "1", agent.System)
	assert.Equal(t, 5, len(agent.Qfe))
	assert.Equal(t, myfakeit.QfeInfo{HotfixId: "KB5994899", InstalledOn: "21/07/2021"}, agent.Qfe[0])
	assert.Equal(t, 2, len(agent.Volume))
	assert.Equal(t, myfakeit.VolumeInfo{Total: "2333396909501", Free: "2888082257", Type: "NTFS"}, agent.Volume[0])
	assert.Equal(t, 7, len(agent.Meta))
	assert.Equal(t, myfakeit.MetaInfo{"cpuCaption": "Intel64 Family 6 Model 167 Stepping 1"}, agent.Meta[0])
}
