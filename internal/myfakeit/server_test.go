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
	assert.Equal(t, "TES1ZAEXZ", server.SKU)
	assert.Equal(t, "THYM3Q2DIFHYM0E490JPTGN71CBMJIX2", server.AC)
	assert.Equal(t, 5, len(server.Volume))
	assert.Equal(t, "/dev/mapper/vg_redis-lvm_redis", server.Volume[0].Source)
	assert.Equal(t, "3464667796943", server.Volume[0].Total)
	assert.Equal(t, "2275832852", server.Volume[0].Free)
	assert.Equal(t, "ext4", server.Volume[0].Type)
	assert.Equal(t, 4, len(server.EnabledFeature))
	assert.Equal(t, "schedule-update", server.EnabledFeature[0])
	assert.Equal(t, 5, len(server.Meta))
}
