package factory

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
	"github.com/stretchr/testify/assert"
)

func TestNewAgentCertCollection(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentCertCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	agent := agents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, 100, len(agent.Cert))
	assert.Equal(t, "LRXSW 3.45", agent.Cert[0].Subject)
	assert.Equal(t, "DigiCert Global Root CA", agent.Cert[0].Issuer)
	assert.Equal(t, "6", agent.Cert[0].Serial)
	assert.Equal(t, "e7f6c011776e8db7cd330b54174fd76f7d0216b612387a5ffcfb81e6f0919683", agent.Cert[0].SHA2)
}

func TestEncodeAgentCertCollectionFlat(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentCertCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionFlat()
	assert.Equal(t, "687f9ed2ed6c41e7a468f9af21d70de3", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-cert", telemetry.TelemetryType)

	associatedAgents := telemetry.Associations.(AgentCertCollection)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := agents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Equal(t, 100, len(agent.Cert))
	assert.Equal(t, "LRXSW 3.45", agent.Cert[0].Subject)
	assert.Equal(t, "DigiCert Global Root CA", agent.Cert[0].Issuer)
	assert.Equal(t, "6", agent.Cert[0].Serial)
	assert.Equal(t, "e7f6c011776e8db7cd330b54174fd76f7d0216b612387a5ffcfb81e6f0919683", agent.Cert[0].SHA2)
}

func TestEncodeAgentCertCollectionLookup(t *testing.T) {
	gofakeit.Seed(11)

	agents := NewAgentCertCollection(5)
	assert.Equal(t, 5, len(agents.Agents))

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "687f9ed2ed6c41e7a468f9af21d70de3", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-cert", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetryCertLookup)
	assert.Greater(t, len(lookup.CertMap), 0)

	cert, ok := lookup.CertMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "LRXSW 3.45", cert.Subject)
	assert.Equal(t, "DigiCert Global Root CA", cert.Issuer)
	assert.Equal(t, "6", cert.Serial)
	assert.Equal(t, "e7f6c011776e8db7cd330b54174fd76f7d0216b612387a5ffcfb81e6f0919683", cert.SHA2)

	associatedAgents := telemetry.Associations.(AgentTelemetryCertAssociationsLookup)
	assert.Equal(t, 5, len(associatedAgents.Agents))

	agent := associatedAgents.Agents[0]
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", agent.Guid)
	assert.Greater(t, len(agent.Cert), 0)
	assert.Equal(t, "1", agent.Cert[0])
}

func TestEncodeAgentCertCollectionLookuWithSameCert(t *testing.T) {
	gofakeit.Seed(11)

	var agents AgentCertCollection

	// declare fake cert
	certList := make([]myfakeit.CertInfo, 3)
	certList[0] = myfakeit.CertInfo{
		Subject: "subject1", Issuer: "issuer1", Serial: "1", SHA2: "hash1",
	}
	certList[1] = myfakeit.CertInfo{
		Subject: "subject2", Issuer: "issuer2", Serial: "2", SHA2: "hash2",
	}
	certList[2] = myfakeit.CertInfo{
		Subject: "subject3", Issuer: "issuer3", Serial: "3", SHA2: "hash3",
	}

	// declare fake agents
	agents.Agents = make([]myfakeit.AgentCert, 3)
	agents.Agents[0] = myfakeit.AgentCert{
		Guid: "guid1",
		Cert: certList,
	}
	agents.Agents[1] = myfakeit.AgentCert{
		Guid: "guid2",
		Cert: certList,
	}
	agents.Agents[2] = myfakeit.AgentCert{
		Guid: "guid3",
		Cert: certList,
	}

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-cert", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetryCertLookup)
	assert.Equal(t, 3, len(lookup.CertMap))

	var cert myfakeit.CertInfo
	var ok bool
	cert, ok = lookup.CertMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "subject1", cert.Subject)
	assert.Equal(t, "issuer1", cert.Issuer)
	assert.Equal(t, "1", cert.Serial)
	assert.Equal(t, "hash1", cert.SHA2)

	cert, ok = lookup.CertMap["2"]
	assert.True(t, ok)
	assert.Equal(t, "subject2", cert.Subject)
	assert.Equal(t, "issuer2", cert.Issuer)
	assert.Equal(t, "2", cert.Serial)
	assert.Equal(t, "hash2", cert.SHA2)

	cert, ok = lookup.CertMap["3"]
	assert.True(t, ok)
	assert.Equal(t, "subject3", cert.Subject)
	assert.Equal(t, "issuer3", cert.Issuer)
	assert.Equal(t, "3", cert.Serial)
	assert.Equal(t, "hash3", cert.SHA2)

	associatedAgents := telemetry.Associations.(AgentTelemetryCertAssociationsLookup)
	assert.Equal(t, 3, len(associatedAgents.Agents))

	var agent AgentCert
	agent = associatedAgents.Agents[0]
	assert.Equal(t, "guid1", agent.Guid)
	assert.Equal(t, len(agent.Cert), 3)
	assert.Equal(t, []string{"1", "2", "3"}, agent.Cert)

	agent = associatedAgents.Agents[1]
	assert.Equal(t, "guid2", agent.Guid)
	assert.Equal(t, len(agent.Cert), 3)
	assert.Equal(t, []string{"1", "2", "3"}, agent.Cert)

	agent = associatedAgents.Agents[2]
	assert.Equal(t, "guid3", agent.Guid)
	assert.Equal(t, len(agent.Cert), 3)
	assert.Equal(t, []string{"1", "2", "3"}, agent.Cert)
}

func TestEncodeAgentCertCollectionLookuWithDifferentCert(t *testing.T) {
	gofakeit.Seed(11)

	var agents AgentCertCollection

	// declare fake cert
	var certList []myfakeit.CertInfo
	certList = append(certList, myfakeit.CertInfo{
		Subject: "subject1", Issuer: "issuer1", Serial: "1", SHA2: "hash1"},
	)

	// declare fake agents
	agents.Agents = make([]myfakeit.AgentCert, 3)
	agents.Agents[0] = myfakeit.AgentCert{
		Guid: "guid1",
		Cert: certList,
	}

	certList = append(certList, myfakeit.CertInfo{
		Subject: "subject2", Issuer: "issuer2", Serial: "2", SHA2: "hash2"},
	)
	agents.Agents[1] = myfakeit.AgentCert{
		Guid: "guid2",
		Cert: certList,
	}

	certList = append(certList, myfakeit.CertInfo{
		Subject: "subject3", Issuer: "issuer3", Serial: "3", SHA2: "hash3"},
	)
	agents.Agents[2] = myfakeit.AgentCert{
		Guid: "guid3",
		Cert: certList,
	}

	telemetry := agents.EncodeAgentCollectionLookup()
	assert.Equal(t, "590c1440988845b0bd51a817ee07c3f2", telemetry.ServerGuid)
	assert.Equal(t, "agent-telemetry-cert", telemetry.TelemetryType)

	lookup := telemetry.Lookup.(AgentTelemetryCertLookup)
	assert.Equal(t, 3, len(lookup.CertMap))

	var cert myfakeit.CertInfo
	var ok bool
	cert, ok = lookup.CertMap["1"]
	assert.True(t, ok)
	assert.Equal(t, "subject1", cert.Subject)
	assert.Equal(t, "issuer1", cert.Issuer)
	assert.Equal(t, "1", cert.Serial)
	assert.Equal(t, "hash1", cert.SHA2)

	cert, ok = lookup.CertMap["2"]
	assert.True(t, ok)
	assert.Equal(t, "subject2", cert.Subject)
	assert.Equal(t, "issuer2", cert.Issuer)
	assert.Equal(t, "2", cert.Serial)
	assert.Equal(t, "hash2", cert.SHA2)

	cert, ok = lookup.CertMap["3"]
	assert.True(t, ok)
	assert.Equal(t, "subject3", cert.Subject)
	assert.Equal(t, "issuer3", cert.Issuer)
	assert.Equal(t, "3", cert.Serial)
	assert.Equal(t, "hash3", cert.SHA2)

	associatedAgents := telemetry.Associations.(AgentTelemetryCertAssociationsLookup)
	assert.Equal(t, 3, len(associatedAgents.Agents))

	var agent AgentCert
	agent = associatedAgents.Agents[0]
	assert.Equal(t, "guid1", agent.Guid)
	assert.Equal(t, len(agent.Cert), 1)
	assert.Equal(t, []string{"1"}, agent.Cert)

	agent = associatedAgents.Agents[1]
	assert.Equal(t, "guid2", agent.Guid)
	assert.Equal(t, len(agent.Cert), 2)
	assert.Equal(t, []string{"1", "2"}, agent.Cert)

	agent = associatedAgents.Agents[2]
	assert.Equal(t, "guid3", agent.Guid)
	assert.Equal(t, len(agent.Cert), 3)
	assert.Equal(t, []string{"1", "2", "3"}, agent.Cert)
}
