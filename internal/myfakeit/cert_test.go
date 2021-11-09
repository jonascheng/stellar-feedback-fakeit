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
	assert.Equal(t, "VT5", agent.Cert[0].Subject)
	assert.Equal(t, "Class 1 Public Primary Certification Authority", agent.Cert[0].Issuer)
	assert.Equal(t, "MTA=", agent.Cert[0].Serial)
	assert.Equal(t, "4a44dc15364204a80fe80e9039455cc1608281820fe2b24f1e5233ade6af1dd5", agent.Cert[0].SHA2)
}
