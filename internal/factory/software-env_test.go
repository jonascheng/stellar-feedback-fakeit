package factory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
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
	assert.Greater(t, len(lookup.AppMap), 0)

	software, ok := lookup.AppMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "LRXSW 3.45", software.Caption)
	assert.Equal(t, "3.2.1000", software.Version)
	assert.Equal(t, "LRXSW 3.45", software.Vendor)
	assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\", software.InstallLocation)

	associatedAgents := telemetry.Associations.(AgentTelemetrySoftwareAssociationsLookup)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := associatedAgents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Greater(t, len(agent.App), 0)
	assert.Equal(t, "1", agent.App[0])
}

func TestEncodeAgentSoftwareEnvCollectionLookuWithSameApp(t *testing.T) {
	gofakeit.Seed(11)

	var agents AgentSoftwareEnvCollection

	// declare fake software
	app := make([]myfakeit.AppInfo, 3)
	app[0] = myfakeit.AppInfo{
		Caption: "test1", Version: "1", Vendor: "test1", InstallLocation: "C:\\Program Files\\test1\\",
	}
	app[1] = myfakeit.AppInfo{
		Caption: "test2", Version: "2", Vendor: "test2", InstallLocation: "C:\\Program Files\\test2\\",
	}
	app[2] = myfakeit.AppInfo{
		Caption: "test3", Version: "3", Vendor: "test3", InstallLocation: "C:\\Program Files\\test3\\",
	}

	// declare fake agents
	agents.Agents = make([]myfakeit.AgentSoftwareEnv, 3)
	agents.Agents[0] = myfakeit.AgentSoftwareEnv{
		Guid: "guid1",
		App:  app,
	}
	agents.Agents[1] = myfakeit.AgentSoftwareEnv{
		Guid: "guid2",
		App:  app,
	}
	agents.Agents[2] = myfakeit.AgentSoftwareEnv{
		Guid: "guid3",
		App:  app,
	}

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-software-environment", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetrySoftwareLookup)
	assert.Equal(t, 3, len(lookup.AppMap))

	var software AgentSoftwareApplication
	var ok bool
	software, ok = lookup.AppMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "test1", software.Caption)
	assert.Equal(t, "1", software.Version)
	assert.Equal(t, "test1", software.Vendor)
	assert.Equal(t, "C:\\Program Files\\test1\\", software.InstallLocation)

	software, ok = lookup.AppMap["2"]
	assert.True(t, ok)
	assert.Equal(t, "test2", software.Caption)
	assert.Equal(t, "2", software.Version)
	assert.Equal(t, "test2", software.Vendor)
	assert.Equal(t, "C:\\Program Files\\test2\\", software.InstallLocation)

	software, ok = lookup.AppMap["3"]
	assert.True(t, ok)
	assert.Equal(t, "test3", software.Caption)
	assert.Equal(t, "3", software.Version)
	assert.Equal(t, "test3", software.Vendor)
	assert.Equal(t, "C:\\Program Files\\test3\\", software.InstallLocation)

	associatedAgents := telemetry.Associations.(AgentTelemetrySoftwareAssociationsLookup)
	assert.Equal(t, 3, len(associatedAgents.Agents))

	var agent AgentSoftwareEnv
	agent = associatedAgents.Agents[0]
	assert.Equal(t, "guid1", agent.Guid)
	assert.Equal(t, len(agent.App), 3)
	assert.Equal(t, []string{"1", "2", "3"}, agent.App)

	agent = associatedAgents.Agents[1]
	assert.Equal(t, "guid2", agent.Guid)
	assert.Equal(t, len(agent.App), 3)
	assert.Equal(t, []string{"1", "2", "3"}, agent.App)

	agent = associatedAgents.Agents[2]
	assert.Equal(t, "guid3", agent.Guid)
	assert.Equal(t, len(agent.App), 3)
	assert.Equal(t, []string{"1", "2", "3"}, agent.App)
}

func TestEncodeAgentSoftwareEnvCollectionLookuWithDifferentApp(t *testing.T) {
	gofakeit.Seed(11)

	var agents AgentSoftwareEnvCollection

	// declare fake software
	var app []myfakeit.AppInfo
	app = append(app, myfakeit.AppInfo{
		Caption: "test1", Version: "1", Vendor: "test1", InstallLocation: "C:\\Program Files\\test1\\",
	})

	// declare fake agents
	agents.Agents = make([]myfakeit.AgentSoftwareEnv, 3)
	agents.Agents[0] = myfakeit.AgentSoftwareEnv{
		Guid: "guid1",
		App:  app,
	}

	app = append(app, myfakeit.AppInfo{
		Caption: "test2", Version: "2", Vendor: "test2", InstallLocation: "C:\\Program Files\\test2\\",
	})
	agents.Agents[1] = myfakeit.AgentSoftwareEnv{
		Guid: "guid2",
		App:  app,
	}

	app = append(app, myfakeit.AppInfo{
		Caption: "test3", Version: "3", Vendor: "test3", InstallLocation: "C:\\Program Files\\test3\\",
	})
	agents.Agents[2] = myfakeit.AgentSoftwareEnv{
		Guid: "guid3",
		App:  app,
	}

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-software-environment", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetrySoftwareLookup)
	assert.Equal(t, 3, len(lookup.AppMap))

	var software AgentSoftwareApplication
	var ok bool
	software, ok = lookup.AppMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "test1", software.Caption)
	assert.Equal(t, "1", software.Version)
	assert.Equal(t, "test1", software.Vendor)
	assert.Equal(t, "C:\\Program Files\\test1\\", software.InstallLocation)

	software, ok = lookup.AppMap["2"]
	assert.True(t, ok)
	assert.Equal(t, "test2", software.Caption)
	assert.Equal(t, "2", software.Version)
	assert.Equal(t, "test2", software.Vendor)
	assert.Equal(t, "C:\\Program Files\\test2\\", software.InstallLocation)

	software, ok = lookup.AppMap["3"]
	assert.True(t, ok)
	assert.Equal(t, "test3", software.Caption)
	assert.Equal(t, "3", software.Version)
	assert.Equal(t, "test3", software.Vendor)
	assert.Equal(t, "C:\\Program Files\\test3\\", software.InstallLocation)

	associatedAgents := telemetry.Associations.(AgentTelemetrySoftwareAssociationsLookup)
	assert.Equal(t, 3, len(associatedAgents.Agents))

	var agent AgentSoftwareEnv
	agent = associatedAgents.Agents[0]
	assert.Equal(t, "guid1", agent.Guid)
	assert.Equal(t, len(agent.App), 1)
	assert.Equal(t, []string{"1"}, agent.App)

	agent = associatedAgents.Agents[1]
	assert.Equal(t, "guid2", agent.Guid)
	assert.Equal(t, len(agent.App), 2)
	assert.Equal(t, []string{"1", "2"}, agent.App)

	agent = associatedAgents.Agents[2]
	assert.Equal(t, "guid3", agent.Guid)
	assert.Equal(t, len(agent.App), 3)
	assert.Equal(t, []string{"1", "2", "3"}, agent.App)
}
