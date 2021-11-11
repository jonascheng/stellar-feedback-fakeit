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
		assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\ASDA_Soft_V5.exe", threat.File)
		assert.Equal(t, "032a4e38b86d929ca23b3755cff582df406d4cbf758e3c91715e774dc270ede2", threat.Hash)
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
		assert.Equal(t, "C:\\Program Files\\GX8 Design Studio\\GX8 Design Studio.exe", threat.File)
		assert.Equal(t, "40e9004035d707a5646ae117dd7939751728170e76ddca2a72beda729fb9ceea", threat.Hash)
		assert.Equal(t, "Virus", threat.Type)
		assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
		assert.Equal(t, "C:\\Program Files\\TXOne\\StellarProtect\\private\\quarantine\\145935bc-71ed-42b6-be44-aaa8ace0c392", threat.Quarantine)
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
	assert.Equal(t, "C:\\Vijeo Designer Runtime\\Vijeo Designer Runtime.exe", threat.File)
	assert.Equal(t, "226e61623c687060b589b8003eb1338dc12fb2b7414d4a6f7ccbe77380a7692c", threat.Hash)
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
	assert.Equal(t, "C:\\Vijeo Designer Runtime\\Vijeo Designer Runtime.exe", threat.File)
	assert.Equal(t, "226e61623c687060b589b8003eb1338dc12fb2b7414d4a6f7ccbe77380a7692c", threat.Hash)
	assert.Equal(t, "Virus", threat.Type)
	assert.Equal(t, "PE_TEST_VIRUS", threat.Name)
	assert.Equal(t, "C:\\Program Files\\TXOne\\StellarEnforce\\private\\quarantine\\55e3c7a7-6490-43e0-aa0b-6a665862c32e", threat.Quarantine)
	assert.Greater(t, threat.Count, 0)
}
