package factory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestCollectAgentSoftwareEnv(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentSoftwareEnvCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	agent := agents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, 100, len(agent.App))
	assert.Equal(t, "LRXSW 3.45", agent.App[0].Caption)
	assert.Equal(t, "3.2.1000", agent.App[0].Version)
	assert.Equal(t, "LRXSW 3.45", agent.App[0].Vendor)
	assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\", agent.App[0].InstallLocation)
}

func TestEncodeAgentSoftwareEnvCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentSoftwareEnvCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionFlat()
	assert.Equal(t, "0b6db5705a7e4a3fb4dd084b3f16e637", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-software-environment", telemetry.TelemetryType)

	associatedAgents := telemetry.Associations.(AgentSoftwareEnvCollection)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := associatedAgents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, 100, len(agent.App))
	assert.Equal(t, "LRXSW 3.45", agent.App[0].Caption)
	assert.Equal(t, "3.2.1000", agent.App[0].Version)
	assert.Equal(t, "LRXSW 3.45", agent.App[0].Vendor)
	assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\", agent.App[0].InstallLocation)
}

func TestEncodeAgentSoftwareEnvCollectionLookup(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentSoftwareEnvCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "0b6db5705a7e4a3fb4dd084b3f16e637", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-software-environment", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetrySoftwareLookup)
	assert.Greater(t, len(lookup.SoftwareMap), 0)

	software, ok := lookup.SoftwareMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "LRXSW 3.45", software.Caption)
	assert.Equal(t, "3.2.1000", software.Version)
	assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\", software.InstallLocation)

	associatedAgents := telemetry.Associations.(AgentTelemetrySoftwareAssociationsLookup)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := associatedAgents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Greater(t, len(agent.Software), 0)
	assert.Equal(t, "1", agent.Software[0])
}
