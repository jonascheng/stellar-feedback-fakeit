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
	assert.Equal(t, 1000, len(agent.App))
	assert.Equal(t, "145935bc715543c7a76490c3e0aa0b6a", agent.App[0].AppGuid)
	assert.Equal(t, "BetterLesson", agent.App[0].Caption)
	assert.Equal(t, "3.2.1000", agent.App[0].Version)
	assert.Equal(t, "BetterLesson", agent.App[0].Vendor)
	assert.Equal(t, "C:\\Program Files\\GX8 Design Studio\\", agent.App[0].InstallLocation)
}

func TestEncodeAgentSoftwareEnvCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentSoftwareEnvCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeCollectionFlat()
	assert.Equal(t, "c75907d1dc394af4ad324b49646f3507", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-software-environment", telemetry.TelemetryType)

	associatedAgents := telemetry.Associations.(AgentSoftwareEnvCollection)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := associatedAgents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, 1000, len(agent.App))
	assert.Equal(t, "145935bc715543c7a76490c3e0aa0b6a", agent.App[0].AppGuid)
	assert.Equal(t, "BetterLesson", agent.App[0].Caption)
	assert.Equal(t, "3.2.1000", agent.App[0].Version)
	assert.Equal(t, "BetterLesson", agent.App[0].Vendor)
	assert.Equal(t, "C:\\Program Files\\GX8 Design Studio\\", agent.App[0].InstallLocation)
}

func TestEncodeAgentSoftwareEnvCollectionLookup(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentSoftwareEnvCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "c75907d1dc394af4ad324b49646f3507", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-software-environment", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetrySoftwareLookup)
	assert.Greater(t, len(lookup.AppMap), 0)

	app, ok := lookup.AppMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "145935bc715543c7a76490c3e0aa0b6a", app.AppGuid)
	assert.Equal(t, "BetterLesson", app.Caption)
	assert.Equal(t, "3.2.1000", app.Version)
	assert.Equal(t, "BetterLesson", app.Vendor)
	assert.Equal(t, "C:\\Program Files\\GX8 Design Studio\\", app.InstallLocation)

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

	// declare fake app
	appList := make([]myfakeit.AppInfo, 3)
	appList[0] = myfakeit.AppInfo{
		Caption: "test1", Version: "1", Vendor: "test1", InstallLocation: "C:\\Program Files\\test1\\",
	}
	appList[1] = myfakeit.AppInfo{
		Caption: "test2", Version: "2", Vendor: "test2", InstallLocation: "C:\\Program Files\\test2\\",
	}
	appList[2] = myfakeit.AppInfo{
		Caption: "test3", Version: "3", Vendor: "test3", InstallLocation: "C:\\Program Files\\test3\\",
	}

	// declare fake agents
	agents.Agents = make([]myfakeit.AgentSoftwareEnv, 3)
	agents.Agents[0] = myfakeit.AgentSoftwareEnv{
		Guid: "guid1",
		App:  appList,
	}
	agents.Agents[1] = myfakeit.AgentSoftwareEnv{
		Guid: "guid2",
		App:  appList,
	}
	agents.Agents[2] = myfakeit.AgentSoftwareEnv{
		Guid: "guid3",
		App:  appList,
	}

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-software-environment", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetrySoftwareLookup)
	assert.Equal(t, 3, len(lookup.AppMap))

	var app myfakeit.AppInfo
	var ok bool
	app, ok = lookup.AppMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "test1", app.Caption)
	assert.Equal(t, "1", app.Version)
	assert.Equal(t, "test1", app.Vendor)
	assert.Equal(t, "C:\\Program Files\\test1\\", app.InstallLocation)

	app, ok = lookup.AppMap["2"]
	assert.True(t, ok)
	assert.Equal(t, "test2", app.Caption)
	assert.Equal(t, "2", app.Version)
	assert.Equal(t, "test2", app.Vendor)
	assert.Equal(t, "C:\\Program Files\\test2\\", app.InstallLocation)

	app, ok = lookup.AppMap["3"]
	assert.True(t, ok)
	assert.Equal(t, "test3", app.Caption)
	assert.Equal(t, "3", app.Version)
	assert.Equal(t, "test3", app.Vendor)
	assert.Equal(t, "C:\\Program Files\\test3\\", app.InstallLocation)

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

	// declare fake app
	var appList []myfakeit.AppInfo
	appList = append(appList, myfakeit.AppInfo{
		Caption: "test1", Version: "1", Vendor: "test1", InstallLocation: "C:\\Program Files\\test1\\",
	})

	// declare fake agents
	agents.Agents = make([]myfakeit.AgentSoftwareEnv, 3)
	agents.Agents[0] = myfakeit.AgentSoftwareEnv{
		Guid: "guid1",
		App:  appList,
	}

	appList = append(appList, myfakeit.AppInfo{
		Caption: "test2", Version: "2", Vendor: "test2", InstallLocation: "C:\\Program Files\\test2\\",
	})
	agents.Agents[1] = myfakeit.AgentSoftwareEnv{
		Guid: "guid2",
		App:  appList,
	}

	appList = append(appList, myfakeit.AppInfo{
		Caption: "test3", Version: "3", Vendor: "test3", InstallLocation: "C:\\Program Files\\test3\\",
	})
	agents.Agents[2] = myfakeit.AgentSoftwareEnv{
		Guid: "guid3",
		App:  appList,
	}

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-software-environment", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetrySoftwareLookup)
	assert.Equal(t, 3, len(lookup.AppMap))

	var app myfakeit.AppInfo
	var ok bool
	app, ok = lookup.AppMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "test1", app.Caption)
	assert.Equal(t, "1", app.Version)
	assert.Equal(t, "test1", app.Vendor)
	assert.Equal(t, "C:\\Program Files\\test1\\", app.InstallLocation)

	app, ok = lookup.AppMap["2"]
	assert.True(t, ok)
	assert.Equal(t, "test2", app.Caption)
	assert.Equal(t, "2", app.Version)
	assert.Equal(t, "test2", app.Vendor)
	assert.Equal(t, "C:\\Program Files\\test2\\", app.InstallLocation)

	app, ok = lookup.AppMap["3"]
	assert.True(t, ok)
	assert.Equal(t, "test3", app.Caption)
	assert.Equal(t, "3", app.Version)
	assert.Equal(t, "test3", app.Vendor)
	assert.Equal(t, "C:\\Program Files\\test3\\", app.InstallLocation)

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
