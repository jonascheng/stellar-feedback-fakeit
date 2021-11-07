package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAgentSystemEnv(t *testing.T) {
	gofakeit.Seed(11)

	agent := SystemEnv()

	assert.Equal(t, "590c1440-9888-45b0-bd51-a817ee07c3f2", agent.Guid)
	assert.Equal(t, "Microsoft Windows 10 Pro for Workstations", agent.Caption)
	assert.Equal(t, "10.0.16299", agent.Version)
	assert.Equal(t, 5, len(agent.Qfe))
	assert.Equal(t, "KB5994899", agent.Qfe[0].HotfixId)
	assert.Equal(t, "21/07/2021", agent.Qfe[0].InstalledOn)
	assert.Equal(t, 2, len(agent.Volume))
	assert.Equal(t, "2333396909501", agent.Volume[0].Total)
	assert.Equal(t, "2888082257", agent.Volume[0].Free)
	assert.Equal(t, "NTFS", agent.Volume[0].Type)
	assert.Equal(t, 7, len(agent.Meta))
}
