package myfakeit

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Agent struct {
	Guid           string    `json:"guid" xml:"guid" fake:"{uuid}"`
	Version        string    `json:"version" xml:"version" fake:"{randomstring:[1.0.1000,1.1.1100]}"`
	Type           int       `json:"type" xml:"type" fake:"{number:0,1}"`
	LastSyncAt     time.Time `json:"timestamp" xml:"timestamp" fake:"skip"`
	TimeGap        int       `json:"timeGap" xml:"timeGap" fake:"{number:[-86400,86400]}"`
	EnabledFeature string    `json:"enabledFeature" xml:"enabledFeature" fake:"skip"`
}

// AgentInfo will generate a struct of agent information
func AgentInfo() *Agent { return agentInfo(globalFaker.Rand) }

func agentInfo(r *rand.Rand) *Agent {
	var s Agent
	if err := gofakeit.Struct(&s); err != nil {
		panic(err)
	}

	// remove dash from UUID
	s.Guid = strings.Replace(s.Guid, "-", "", -1)

	// set LastSyncAt
	s.LastSyncAt = time.Now().UTC()

	// set EnabledFeature
	s.EnabledFeature = getEnabledFeature()

	return &s
}

var once sync.Once
var base64EnabledFeature string

func getEnabledFeature() string {
	once.Do(func() {
		// Read the entire file into a byte slice
		bytes, err := ioutil.ReadFile("./agent-data/opsrv-config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		base64EnabledFeature = base64.StdEncoding.EncodeToString(bytes)
	})
	return base64EnabledFeature
}
