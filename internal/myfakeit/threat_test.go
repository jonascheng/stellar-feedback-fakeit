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
		assert.Equal(t, "C:\\Program Files\\Roboticsware FA-Panel6 (32bit)\\5ae31383-f63e-4c49-b6b2-c503c63bf35c.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Equal(t, "C:\\Program Files\\TXOne\\StellarProtect\\private\\quarantine\\290f4535-64a7-4654-8475-dc9b7be64803", threat.Quarantine)
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
		assert.Equal(t, "C:\\GX8 Design Studio\\0c05853a-4b35-4264-9484-52b9e9bbcfe9.exe", threat.File)
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
		assert.Equal(t, "C:\\Vijeo Designer Runtime\\a5b4d18f-1b1a-4194-ae84-943b3f8abab1.exe", threat.File)
		assert.Equal(t, "Ericka", threat.User)
		assert.Equal(t, "C:\\Windows\\System32\\LRH SW version 4.0.0\\473fe34b-0039-447a-adb3-16977c25b6b7.exe", threat.Parent1)
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
		assert.Equal(t, "C:\\Windows\\System32\\Schneider Electric LinkManager\\afe118e0-7221-4a5a-88ba-e0978f10d371.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.Equal(t, "Gerson", threat.User)
		assert.Greater(t, threat.Count, 0)
	}

	// ADCBlockedEvent
	{
		assert.Greater(t, len(threats.ADCBlocked), 0)
		threat := threats.ADCBlocked[0]
		assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", threat.Guid)
		assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
		assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
		assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
		assert.Equal(t, "C:\\Windows\\SysWOW64\\Roboticsware FA-Server6\\6cd511a2-6361-49b7-98f1-ad49a0b1b5fa.exe", threat.File)
		assert.Equal(t, 64, len(threat.Hash))
		assert.GreaterOrEqual(t, len(threat.Impacted), 3)
		assert.Equal(t, "C:\\Users\\Bertram\\Desktop\\6de8d288-2436-4e29-92d4-f90af7460cf8.docx", threat.Impacted[0])
		assert.Equal(t, "Prevention", threat.Mode)
		assert.Greater(t, threat.Count, 0)
	}
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
	assert.Equal(t, 64, len(threat.Hash))
	assert.Equal(t, "Marques", threat.User)
	assert.Greater(t, threat.Count, 0)
}

func TestADCBlockedEventInfo(t *testing.T) {
	gofakeit.Seed(11)

	threat := ADCBlockedEventInfo("guid1")
	assert.Equal(t, "guid1", threat.Guid)
	assert.Greater(t, threat.TimeslotEnd-threat.TimeslotBegin, int64(0))
	assert.Equal(t, int64(1), (threat.TimeslotEnd-threat.TimeslotBegin)/int64(86400))
	assert.Equal(t, int64(0), (threat.TimeslotEnd-threat.TimeslotBegin)%int64(86400))
	assert.Equal(t, "C:\\Vijeo Designer Runtime\\c3f21459-35bc-4155-a3c7-a76490c3e0aa.exe", threat.File)
	assert.Equal(t, 64, len(threat.Hash))
	assert.GreaterOrEqual(t, len(threat.Impacted), 3)
	assert.Equal(t, "C:\\Users\\Enrique\\Desktop\\0b6a6658-621e-460b-b665-2ef9c6c737e1.docx", threat.Impacted[0])
	assert.Equal(t, "Prevention", threat.Mode)
	assert.Greater(t, threat.Count, 0)
}
