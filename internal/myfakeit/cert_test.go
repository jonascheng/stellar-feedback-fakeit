package myfakeit

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestCert(t *testing.T) {
	gofakeit.Seed(11)

	agent := Cert()

	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, 100, len(agent.Cert))
	assert.Equal(t, "LRXSW 3.45", agent.Cert[0].Subject)
	assert.Equal(t, "DigiCert Global Root CA", agent.Cert[0].Issuer)
	assert.Equal(t, "6", agent.Cert[0].Serial)
	assert.Equal(t, "e7f6c011776e8db7cd330b54174fd76f7d0216b612387a5ffcfb81e6f0919683", agent.Cert[0].SHA2)
}
