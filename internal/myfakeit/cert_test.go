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
	assert.Equal(t, "Civinomics", agent.Cert[0].Subject)
	assert.Equal(t, "DigiCert Global Root CA", agent.Cert[0].Issuer)
	assert.Equal(t, "xHkiJ", agent.Cert[0].Serial)
	assert.Equal(t, "bfff758c3318d335b0cae32492bf81f62d2f243deda6a7adb956a631089c7dbc", agent.Cert[0].SHA2)
}
