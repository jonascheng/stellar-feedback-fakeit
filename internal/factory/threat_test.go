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
		assert.Equal(t, "c3f2145935bc4155a3c7a76490c3e0aa", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Program Files (x86)\\ASDA_Soft_V5\\ASDA_Soft_V5.exe", threat.File)
		assert.Equal(t, "4eef4479d5532869f9eef67104b6559b9bbbeca70d4f345753db8bb7626d4a8b", threat.Hash)
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
	assert.Equal(t, "c33d812f93154fe5bfde4c607f36dfe5", telemetry.ServerGuid)
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
