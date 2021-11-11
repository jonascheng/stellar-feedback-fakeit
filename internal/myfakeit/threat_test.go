package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestThreatInfo(t *testing.T) {
	gofakeit.Seed(11)

	threats := ThreatInfo(1)

	// AppExecBlockedEventInfo
	{
		assert.Greater(t, len(threats.AppExecBlocked), 0)
		threat := threats.AppExecBlocked[0]
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
		assert.Greater(t, len(threats.FileScanBlocked), 0)
		threat := threats.FileScanBlocked[0]
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

	// SuspiciousExecBlockedEvent
	{
		assert.Greater(t, len(threats.SuspiciousExecBlocked), 0)
		threat := threats.SuspiciousExecBlocked[0]
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
		assert.Greater(t, len(threats.OBADBlocked), 0)
		threat := threats.OBADBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Program Files\\OPC Server\\579b5f7e-d39d-4bd6-a386-49a911a617bc.exe", threat.File)
		assert.Equal(t, "Kaylin", threat.User)
		assert.Equal(t, "C:\\Windows\\System32\\Vijeo Designer Runtime\\e1056d92-572c-4c35-b8c8-b1cf58d750fd.exe", threat.Parent1)
		assert.NotEmpty(t, threat.Parent2)
		assert.NotEmpty(t, threat.Parent3)
		assert.NotEmpty(t, threat.Parent4)
		assert.Equal(t, "Detection", threat.Mode)
		assert.Equal(t, "aggressive", threat.Level)
		assert.Greater(t, threat.Count, 0)
	}

	// NonWhitelistingBlockedEvent
	{
		assert.Greater(t, len(threats.NonWhitelistingBlocked), 0)
		threat := threats.NonWhitelistingBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Users\\LRH SW version 4.0.0\\ac7bd37a-e76a-45f2-8d04-bfa1a26ccacf.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Triston", threat.User)
		assert.Greater(t, threat.Count, 0)
	}
	// ADCBlockedEvent
}

func TestAppExecBlockedEventInfo(t *testing.T) {
	gofakeit.Seed(11)

	threat := AppExecBlockedEventInfo("guid1")
	assert.Equal(t, "guid1", threat.Guid)
	assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
	assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
	assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
	assert.Equal(t, "C:\\Vijeo Designer Runtime\\c3f21459-35bc-4155-a3c7-a76490c3e0aa.exe", threat.File)
	assert.Equal(t, 64, len(threat.Hash))
	assert.Equal(t, "Virus", threat.Type)
	assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
	assert.Greater(t, threat.Count, 0)
}

func TestFileScanBlockedEventInfo(t *testing.T) {
	gofakeit.Seed(11)

	threat := FileScanBlockedEventInfo("guid1")
	assert.Equal(t, "guid1", threat.Guid)
	assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
	assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
	assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
	assert.Equal(t, "C:\\Vijeo Designer Runtime\\c3f21459-35bc-4155-a3c7-a76490c3e0aa.exe", threat.File)
	assert.Equal(t, 64, len(threat.Hash))
	assert.Equal(t, "Virus", threat.Type)
	assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
	assert.Equal(t, "C:\\Program Files\\TXOne\\StellarEnforce\\private\\quarantine\\0b6a6658-621f-4af8-9e95-3b74428fef5b", threat.Quarantine)
	assert.Greater(t, threat.Count, 0)
}

func TestSuspiciousExecBlockedEventInfo(t *testing.T) {
	gofakeit.Seed(11)

	threat := SuspiciousExecBlockedEventInfo("guid1")
	assert.Equal(t, "guid1", threat.Guid)
	assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
	assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
	assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
	assert.Equal(t, "C:\\Vijeo Designer Runtime\\c3f21459-35bc-4155-a3c7-a76490c3e0aa.exe", threat.File)
	assert.Equal(t, 64, len(threat.Hash))
	assert.Greater(t, threat.Count, 0)
}

func TestOBADBlockedEventInfo(t *testing.T) {
	gofakeit.Seed(11)

	threat := OBADBlockedEventInfo("guid1")
	assert.Equal(t, "guid1", threat.Guid)
	assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
	assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
	assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
	assert.Equal(t, "C:\\Vijeo Designer Runtime\\c3f21459-35bc-4155-a3c7-a76490c3e0aa.exe", threat.File)
	assert.Equal(t, "Marques", threat.User)
	assert.Equal(t, "C:\\Program Files (x86)\\Vijeo Designer Runtime\\0b6a6658-62c9-422f-94c8-f9319da9d29a.exe", threat.Parent1)
	assert.NotEmpty(t, threat.Parent2)
	assert.NotEmpty(t, threat.Parent3)
	assert.NotEmpty(t, threat.Parent4)
	assert.Equal(t, "Detection", threat.Mode)
	assert.Equal(t, "aggressive", threat.Level)
	assert.Greater(t, threat.Count, 0)
}

func TestNonWhitelistingBlockedEventInfo(t *testing.T) {
	gofakeit.Seed(11)

	threat := NonWhitelistingBlockedEventInfo("guid1")
	assert.Equal(t, "guid1", threat.Guid)
	assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
	assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
	assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
	assert.Equal(t, "C:\\Vijeo Designer Runtime\\c3f21459-35bc-4155-a3c7-a76490c3e0aa.exe", threat.File)
	assert.Equal(t, "454afe8ffed5ba3eafbc8fe0edad7500600330c9da16baeb0238d938c10611f6", threat.Hash)
	assert.Equal(t, "Marques", threat.User)
	assert.Greater(t, threat.Count, 0)
}
