package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAgentSoftwareEnv(t *testing.T) {
	gofakeit.Seed(11)

	agent := SoftwareEnv()

	assert.Equal(t, "590c1440-9888-45b0-bd51-a817ee07c3f2", agent.Guid)
	assert.Equal(t, 100, len(agent.App))
	assert.Equal(t, "LRXSW 3.45", agent.App[0].Caption)
	assert.Equal(t, "3.2.1000", agent.App[0].Version)
	assert.Equal(t, "LRXSW 3.45", agent.App[0].Vendor)
	assert.Equal(t, "C:\\Program Files\\ASDA_Soft_V5\\", agent.App[0].InstallLocation)
}
