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
		assert.Equal(t, "6b8849fe65884ae9970c8fe8da8d86e9", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Program Files\\Buerkert Communicator\\c5bdec5c-6134-40e8-9381-63b27ee1f2d3.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Greater(t, threat.Count, 0)
	}

	// FileScanBlockedEventInfo
	{
		threat := threats.Agents[1].FileScanBlocked[0]
		assert.Equal(t, "6b8849fe65884ae9970c8fe8da8d86e9", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\LRXSW 3.45\\971346af-571a-4202-a2ad-dd8402de89af.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Equal(t, "C:\\Program Files\\TXOne\\StellarEnforce\\private\\quarantine\\aee50967-b4c9-43ae-8004-5f9ebdf4f481", threat.Quarantine)
		assert.Greater(t, threat.Count, 0)

	}

	// SuspiciousExecBlockedEventInfo
	{
		threat := threats.Agents[1].SuspiciousExecBlocked[0]
		assert.Equal(t, "6b8849fe65884ae9970c8fe8da8d86e9", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\System32\\LRXSW 3.45\\14bcd0f3-8395-4a47-a753-eaa5608e7c2a.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Greater(t, threat.Count, 0)
	}

	// OBADBlockedEvent
	{
		threat := threats.Agents[1].OBADBlocked[0]
		assert.Equal(t, "6b8849fe65884ae9970c8fe8da8d86e9", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Users\\LRH SW version 4.0.0\\87a100ed-5f5e-4c57-94c1-1678a4232f4a.exe", threat.File)
		assert.Equal(t, "Lonzo", threat.User)
		assert.NotEmpty(t, threat.Parent1)
		assert.NotEmpty(t, threat.Parent2)
		assert.NotEmpty(t, threat.Parent3)
		assert.NotEmpty(t, threat.Parent4)
		assert.Equal(t, "Detection", threat.Mode)
		assert.Equal(t, "aggressive", threat.Level)
		assert.Greater(t, threat.Count, 0)
	}

	// NonWhitelistingBlockedEvent
	{
		threat := threats.Agents[1].NonWhitelistingBlocked[0]
		assert.Equal(t, "6b8849fe65884ae9970c8fe8da8d86e9", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Buerkert Communicator\\b74fe5f4-5ca3-472d-9be0-585d65acb7e9.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Levi", threat.User)
		assert.Greater(t, threat.Count, 0)
	}

	// ADCBlockedEvent
	{
		threat := threats.Agents[1].ADCBlocked[0]
		assert.Equal(t, "6b8849fe65884ae9970c8fe8da8d86e9", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\System32\\OPC Server\\bf378c78-e5c7-4cf8-95eb-086c3ec71255.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.GreaterOrEqual(t, len(threat.Impacted), 3)
		assert.NotEmpty(t, threat.Impacted[0])
		assert.Equal(t, "Detection", threat.Mode)
		assert.Greater(t, threat.Count, 0)
	}
}

func TestEncodeThreatCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	threats := NewThreatCollection(5)
	assert.Equal(t, 5, len(threats.Agents))

	telemetry := threats.EncodeCollectionFlat()
	assert.Equal(t, "1c8b35fc32cf48b3852e8ca897ec5958", telemetry.ServerGuid)
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
		assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\145935bc-711f-4af8-9e95-3b74428fef5b.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
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
		assert.Equal(t, "C:\\LRXSW 3.45\\16a03737-ec06-42f7-b00e-4789a3b67464.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Equal(t, "C:\\Program Files\\TXOne\\StellarEnforce\\private\\quarantine\\e4b7fea8-01af-49a4-aa90-0bf820e119e7", threat.Quarantine)
		assert.Greater(t, threat.Count, 0)

	}

	// SuspiciousExecBlockedEventInfo
	{
		threat := associatedThreats.SuspiciousExecBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\Roboticsware FA-Server6\\442da0f3-252d-4c9c-9d2c-2c88d85e5cdc.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Greater(t, threat.Count, 0)
	}

	// OBADBlockedEvent
	{
		threat := associatedThreats.OBADBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\Buerkert Communicator\\56168ba3-0fdd-4889-a759-45b23b6351eb.exe", threat.File)
		assert.Equal(t, "Pauline", threat.User)
		assert.NotEmpty(t, threat.Parent1)
		assert.NotEmpty(t, threat.Parent2)
		assert.NotEmpty(t, threat.Parent3)
		assert.NotEmpty(t, threat.Parent4)
		assert.Equal(t, "Prevention", threat.Mode)
		assert.Equal(t, "aggressive", threat.Level)
		assert.Greater(t, threat.Count, 0)
	}

	// NonWhitelistingBlockedEvent
	{
		threat := associatedThreats.NonWhitelistingBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\GX8 Design Studio\\b8145dc6-2d80-4f88-a734-59c6fe289734.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Lorna", threat.User)
		assert.Greater(t, threat.Count, 0)
	}

	// ADCBlockedEvent
	{
		threat := associatedThreats.ADCBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\Panel Studio V2.31\\6479fb30-0be1-473e-ac95-19e7b27bc2ad.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.GreaterOrEqual(t, len(threat.Impacted), 3)
		assert.NotEmpty(t, threat.Impacted[0])
		assert.Equal(t, "Prevention", threat.Mode)
		assert.Greater(t, threat.Count, 0)
	}
}
