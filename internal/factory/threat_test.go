package factory

import (
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
		assert.Equal(t, "38e86d1d05a54bff879e47811a3ab4d3", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\OPC Server\\OPC Server.exe", threat.File)
		assert.Equal(t, "ad76c196e6ac105f02197f91d4d82a6f0d39f48ab796d2cea6cd0358d650674a", threat.Hash)
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Greater(t, threat.Count, 0)
	}

	// FileScanBlockedEventInfo
	{
		threat := threats.Agents[1].FileScanBlocked[0]
		assert.Equal(t, "38e86d1d05a54bff879e47811a3ab4d3", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\GX8 Design Studio\\GX8 Design Studio.exe", threat.File)
		assert.Equal(t, "58f4520d0acc3f3911f9ba0a7acdcbda3bc148648ff4dbb7bfc54deb07e6657c", threat.Hash)
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Equal(t, "C:\\Program Files\\TXOne\\StellarEnforce\\private\\quarantine\\623555c1-6a36-4dff-9ebc-f54a9e427a61", threat.Quarantine)
		assert.Greater(t, threat.Count, 0)

	}
}

func TestEncodeThreatCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	threats := NewThreatCollection(5)
	assert.Equal(t, 5, len(threats.Agents))

	telemetry := threats.EncodeCollectionFlat()
	assert.Equal(t, "b4f1fca7f807412ea9e19db681d82995", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-threat", telemetry.TelemetryType)

	associatedThreats := telemetry.Associations.(ThreatTelemetryAssociations)
	assert.Greater(t, len(associatedThreats.AppExecBlocked), 0)

	// AppExecBlockedEventInfo
	{
		threat := associatedThreats.AppExecBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\ASDA_Soft_V5.exe", threat.File)
		assert.Equal(t, "032a4e38b86d929ca23b3755cff582df406d4cbf758e3c91715e774dc270ede2", threat.Hash)
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Greater(t, threat.Count, 0)
	}

	// FileScanBlockedEventInfo
	{
		threat := associatedThreats.FileScanBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Program Files\\GX8 Design Studio\\GX8 Design Studio.exe", threat.File)
		assert.Equal(t, "40e9004035d707a5646ae117dd7939751728170e76ddca2a72beda729fb9ceea", threat.Hash)
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Equal(t, "C:\\Program Files\\TXOne\\StellarProtect\\private\\quarantine\\145935bc-71ed-42b6-be44-aaa8ace0c392", threat.Quarantine)
		assert.Greater(t, threat.Count, 0)

	}
}
