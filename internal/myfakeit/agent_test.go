package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestAgentInfo(t *testing.T) {
	gofakeit.Seed(11)

	agent := AgentInfo()

	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, "1.0.1000", agent.Version)
	assert.Equal(t, 0, agent.Type)
	assert.Equal(t, 15888, agent.TimeGap)
}
