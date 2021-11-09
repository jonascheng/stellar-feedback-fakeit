package myfakeit

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

type AgentCert struct {
	Guid string     `json:"guid" xml:"guid" fake:"{uuid}"`
	Cert []CertInfo `json:"cert" xml:"cert" fakesize:"100"`
}

type CertInfo struct {
	Subject string `json:"subject" xml:"subject" fake:"{app-caption}"`
	Issuer  string `json:"issuer" xml:"issuer" fake:"{cert-issuer}"`
	Serial  string `json:"serial" xml:"serial" fake:"{number:1,10}"`
	SHA2    string `json:"sha2" xml:"sha2" fake:"skip"`
}

// SoftwareEnv will generate a struct of software information
func Cert() *AgentCert { return cert(globalFaker.Rand) }

func cert(r *rand.Rand) *AgentCert {
	var s AgentCert
	if err := gofakeit.Struct(&s); err != nil {
		panic(err)
	}

	// remove dash from UUID
	s.Guid = strings.Replace(s.Guid, "-", "", -1)

	// set SHA2
	for i, cert := range s.Cert {
		h := sha256.New()
		if _, err := h.Write([]byte(cert.Serial)); err != nil {
			panic(err)
		}
		s.Cert[i].SHA2 = hex.EncodeToString(h.Sum(nil))
	}

	return &s
}

// CertIssuer will generate a random cert issuer string
func CertIssuer() string { return certIssuer(globalFaker.Rand) }

func certIssuer(r *rand.Rand) string {
	return getRandValue(r, []string{"Cert", "issuer"})
}

func addCertLookup() {
	gofakeit.AddFuncLookup("cert", gofakeit.Info{
		Display:     "Cert",
		Category:    "cert",
		Description: "Random set of cert information",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return cert(r), nil
		},
	})

	gofakeit.AddFuncLookup("cert-issuer", gofakeit.Info{
		Display:     "Cert Issuer",
		Category:    "cert-issuer",
		Description: "Random Cert Issuer",
		Example:     "Microsoft Code Signing PCA",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return certIssuer(r), nil
		},
	})
}
