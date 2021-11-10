package factory

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestNewThreatCollection(t *testing.T) {
	gofakeit.Seed(11)

	threats := NewThreatCollection(5)
	assert.Equal(t, 5, len(threats.Agents))

	// AppExecBlockedEventInfo
	{
		threat := threats.Agents[1].AppExecBlocked[0]
		assert.Equal(t, "55e3c7a7649043e0aa0b6a665862c32e", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Program Files (x86)\\Vijeo Designer Runtime\\Vijeo Designer Runtime.exe", threat.File)
		assert.Equal(t, "7d8ed22559df165701e6efdfd6f1bc8e84690a13a285ada096c2c8a00553cec7", threat.Hash)
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Greater(t, threat.Count, 0)
	}

	fmt.Println(threats)
}

func TestEncodeThreatCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	threats := NewThreatCollection(5)
	assert.Equal(t, 5, len(threats.Agents))

	telemetry := threats.EncodeCollectionFlat()
	assert.Equal(t, "f1fa55b4395f475fa5307089ee07516e", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-threat", telemetry.TelemetryType)

	fmt.Println(telemetry)
	// associatedAgents := telemetry.Associations.(AgentCollection)
	// assert.Equal(t, 5, len(associatedAgents.Agents))

	// agent := agents.Agents[0]
	// assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	// assert.Equal(t, "1.0.1000", agent.Version)
	// assert.Equal(t, 0, agent.Type)
	// assert.Equal(t, 15888, agent.TimeGap)
}
