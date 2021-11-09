package factory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestNewAgentCollection(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	agent := agents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, "1.0.1000", agent.Version)
	assert.Equal(t, 0, agent.Type)
	assert.Equal(t, 1338363715, agent.TimeGap)
}

func TestEncodeAgentCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionFlat()
	assert.Equal(t, "56769046311d483d91c150205f6674f4", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry", telemetry.TelemetryType)

	associatedAgents := telemetry.Associations.(AgentCollection)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := agents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, "1.0.1000", agent.Version)
	assert.Equal(t, 0, agent.Type)
	assert.Equal(t, 1338363715, agent.TimeGap)
}

// identical to EncodeAgentCollectionFlat
func TestEncodeAgentCollectionLookup(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "56769046311d483d91c150205f6674f4", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry", telemetry.TelemetryType)

	associatedAgents := telemetry.Associations.(AgentCollection)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := agents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, "1.0.1000", agent.Version)
	assert.Equal(t, 0, agent.Type)
	assert.Equal(t, 1338363715, agent.TimeGap)
}
