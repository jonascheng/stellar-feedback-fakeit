package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAgentSoftwareEnv(t *testing.T) {
	gofakeit.Seed(11)

	agent := SoftwareEnv()

	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, 1000, len(agent.App))
	assert.Equal(t, "145935bc715543c7a76490c3e0aa0b6a", agent.App[0].AppGuid)
	assert.Equal(t, "BetterLesson", agent.App[0].Caption)
	assert.Equal(t, "3.2.1000", agent.App[0].Version)
	assert.Equal(t, "BetterLesson", agent.App[0].Vendor)
	assert.Equal(t, "C:\\Program Files\\GX8 Design Studio\\", agent.App[0].InstallLocation)
}
