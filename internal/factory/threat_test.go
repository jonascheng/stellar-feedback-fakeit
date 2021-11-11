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
		assert.Equal(t, "9e665f526a3d467a8854b5b1505cfb85", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Program Files (x86)\\Roboticsware FA-Server6\\71904b9e-b832-4ae9-a762-ebc1087c92b0.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Greater(t, threat.Count, 0)
	}

	// FileScanBlockedEventInfo
	{
		threat := threats.Agents[1].FileScanBlocked[0]
		assert.Equal(t, "9e665f526a3d467a8854b5b1505cfb85", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\VT5\\8b9a7e03-7a5a-408c-aa52-62898f84c1cc.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Equal(t, "C:\\Program Files\\TXOne\\StellarEnforce\\private\\quarantine\\dcee4333-ad57-4f8d-8dfc-c98c27ae6c69", threat.Quarantine)
		assert.Greater(t, threat.Count, 0)

	}

	// SuspiciousExecBlockedEventInfo
	{
		threat := threats.Agents[1].SuspiciousExecBlocked[0]
		assert.Equal(t, "9e665f526a3d467a8854b5b1505cfb85", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Program Files (x86)\\ASDA_Soft_V5\\34cfa02d-09e6-46d9-8f2e-4b800da503a4.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Greater(t, threat.Count, 0)
	}

	// OBADBlockedEvent
	{
		threat := threats.Agents[1].OBADBlocked[0]
		assert.Equal(t, "9e665f526a3d467a8854b5b1505cfb85", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\GX8 Design Studio\\bcc0ae04-be1c-42f1-87cc-1b4391df7273.exe", threat.File)
		assert.Equal(t, "Nat", threat.User)
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
		threat := threats.Agents[1].NonWhitelistingBlocked[0]
		assert.Equal(t, "9e665f526a3d467a8854b5b1505cfb85", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\OPC Server\\88a3f15d-1628-46bb-a609-6d118d6e9f8a.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Loy", threat.User)
		assert.Greater(t, threat.Count, 0)
	}

	// ADCBlockedEvent
	{
		threat := threats.Agents[1].ADCBlocked[0]
		assert.Equal(t, "9e665f526a3d467a8854b5b1505cfb85", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Users\\Roboticsware FA-Panel6 (32bit)\\0169344d-8c42-424b-b780-1573093ab49b.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.GreaterOrEqual(t, len(threat.Impacted), 3)
		assert.NotEmpty(t, threat.Impacted[0])
		assert.Equal(t, "Prevention", threat.Mode)
		assert.Greater(t, threat.Count, 0)
	}
}

func TestEncodeThreatCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	threats := NewThreatCollection(5)
	assert.Equal(t, 5, len(threats.Agents))

	telemetry := threats.EncodeCollectionFlat()
	assert.Equal(t, "3e199767913046deab1678448b32f2f7", telemetry.ServerGuid)
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
		assert.Equal(t, "C:\\Program Files\\LRXSW 3.45\\02cc3379-6a30-4aff-8b68-0eece23fc937.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Equal(t, "C:\\Program Files\\TXOne\\StellarProtect\\private\\quarantine\\36b2bdba-8956-472b-891d-702db6ce1b64", threat.Quarantine)
		assert.Greater(t, threat.Count, 0)

	}

	// SuspiciousExecBlockedEventInfo
	{
		threat := associatedThreats.SuspiciousExecBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Users\\GX8 Design Studio\\b76f4ccc-1fe7-498a-a33b-6dd43b272b52.exe", threat.File)
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
		assert.Equal(t, "C:\\Roboticsware FA-Panel6 (32bit)\\8f86b660-9f02-4a76-8e22-cf85566bc4e5.exe", threat.File)
		assert.Equal(t, "Lizzie", threat.User)
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
		threat := associatedThreats.NonWhitelistingBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\GX8 Design Studio\\aaa5c386-64fc-4228-9f1e-e6c471a48945.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Mona", threat.User)
		assert.Greater(t, threat.Count, 0)
	}

	// ADCBlockedEvent
	{
		threat := associatedThreats.ADCBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Users\\Roboticsware FA-Server6\\4edd0942-16a0-4737-ac06-bd87c5605512.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.GreaterOrEqual(t, len(threat.Impacted), 3)
		assert.NotEmpty(t, threat.Impacted[0])
		assert.Equal(t, "Prevention", threat.Mode)
		assert.Greater(t, threat.Count, 0)
	}
}
