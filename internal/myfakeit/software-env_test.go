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
	assert.Equal(t, 100, len(agent.App))
	assert.Equal(t, "VT5", agent.App[0].Caption)
	assert.Equal(t, "2.2.1000", agent.App[0].Version)
	assert.Equal(t, "VT5", agent.App[0].Vendor)
	assert.Equal(t, "C:\\Program Files\\Cappex\\", agent.App[0].InstallLocation)
}
