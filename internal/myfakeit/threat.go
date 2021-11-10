package myfakeit

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Threat struct {
	AppExecBlocked []AppExecBlockedEvent `json:"appExecBlocked,omitempty" xml:"appExecBlocked" fake:"skip"`
}

type AppExecBlockedEvent struct {
	Guid          string `json:"guid" xml:"guid" fake:"skip"`
	TimeslotBegin int64  `json:"timeslotBegin" xml:"timeslotBegin" fake:"skip"`
	TimeslotEnd   int64  `json:"timeslotEnd" xml:"timeslotEnd" fake:"skip"`
	File          string `json:"flie" xml:"flie" fake:"{threat-filename}"`
	Hash          string `json:"hash" xml:"hash" fake:"skip"`
	Type          string `json:"type" xml:"type" fake:"{randomstring:[Virus]}"`
	Name          string `json:"name" xml:"name" fake:"{randomstring:[PE_TEST_VIRUS]}"`
	Count         int    `json:"count" xml:"count" fake:"{number:1,1000}"`
}

// ThreatInfo will generate a struct of threat information
func ThreatInfo(fakeall bool) *Threat {
	return threatInfo(globalFaker.Rand, fakeall)
}

func threatInfo(r *rand.Rand, fakeall bool) *Threat {
	var s Threat

	// AppExecBlockedEvents
	if gofakeit.Bool() || fakeall { // fakeit or not?
		for nEvents := gofakeit.Number(1, 10); nEvents > 0; nEvents-- {
			uuid := gofakeit.UUID()
			// remove dash from UUID
			uuid = strings.Replace(uuid, "-", "", -1)

			event := appExecBlockedEventInfo(r, uuid)
			s.AppExecBlocked = append(s.AppExecBlocked, *event)
		}
	}

	return &s
}

// AppExecBlockedEventInfo will generate a struct of system information
func AppExecBlockedEventInfo(uuid string) *AppExecBlockedEvent {
	return appExecBlockedEventInfo(globalFaker.Rand, uuid)
}

func appExecBlockedEventInfo(r *rand.Rand, uuid string) *AppExecBlockedEvent {
	var s AppExecBlockedEvent

	if err := gofakeit.Struct(&s); err != nil {
		panic(err)
	}

	s.Guid = uuid

	// obtain time on the hour
	subHour := gofakeit.Number(-23, -1)
	fakeNow := time.Now().Add(time.Hour * time.Duration(subHour)).UTC()
	timestamp := fakeNow.Unix() - int64(fakeNow.Second()) - int64(60*fakeNow.Minute())
	s.TimeslotBegin = timestamp
	s.TimeslotEnd = timestamp + 86400

	// obtain file hash
	h := sha256.New()
	if _, err := h.Write([]byte(s.File)); err != nil {
		panic(err)
	}
	s.Hash = hex.EncodeToString(h.Sum(nil))

	return &s
}

func threatFilename(r *rand.Rand) string {
	base := gofakeit.RandomString([]string{"C:", "C:\\Program Files (x86)", "C:\\Program Files", "C:\\Users", "C:\\Windows\\System32", "C:\\Windows\\SysWOW64"})
	folder := getRandValue(r, []string{"Application", "caption"})
	filename := folder
	return fmt.Sprintf("%s\\%s\\%s.exe", base, folder, filename)
}

func addTreatLookup() {
	gofakeit.AddFuncLookup("threat-filename", gofakeit.Info{
		Display:     "Threat Filename",
		Category:    "threat-filename",
		Description: "Random set of threat filename",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return threatFilename(r), nil
		},
	})
}
