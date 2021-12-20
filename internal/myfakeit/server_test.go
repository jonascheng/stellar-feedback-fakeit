package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestServerInfo(t *testing.T) {
	gofakeit.Seed(11)

	server := ServerInfo()

	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", server.Guid)
	assert.Equal(t, "1.0.1000", server.Version)
	assert.Equal(t, []string{"PTHYM3Q2DIFHYM0E490JPTGN71CBMJIX", "28UYLHB7PV5V3YDCXJGPA9091PA58EU5"}, server.AC)
	assert.Equal(t, 4, len(server.EnabledFeature))
	assert.Equal(t, "forward-syslog", server.EnabledFeature[0])
	assert.Equal(t, 5, len(server.Meta))
}
