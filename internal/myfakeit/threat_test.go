package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestThreatInfo(t *testing.T) {
	gofakeit.Seed(11)

	threats := ThreatInfo(true)

	// AppExecBlockedEventInfo
	{
		assert.Greater(t, len(threats.AppExecBlocked), 0)
		threat := threats.AppExecBlocked[0]
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
