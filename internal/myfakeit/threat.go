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
	AppExecBlocked         []AppExecBlockedEvent         `json:"appExecBlocked,omitempty" xml:"appExecBlocked,omitempty" fake:"skip"`
	FileScanBlocked        []FileScanBlockedEvent        `json:"fileScanBlocked,omitempty" xml:"fileScanBlocked,omitempty" fake:"skip"`
	SuspiciousExecBlocked  []SuspiciousExecBlockedEvent  `json:"suspiciousExecBlocked,omitempty" xml:"suspiciousExecBlocked,omitempty" fake:"skip"`
	OBADBlocked            []OBADBlockedEvent            `json:"obadBlocked,omitempty" xml:"obadBlocked,omitempty" fake:"skip"`
	NonWhitelistingBlocked []NonWhitelistingBlockedEvent `json:"nonWhitelistingBlocked,omitempty" xml:"nonWhitelistingBlocked,omitempty" fake:"skip"`
	ADCBlocked             []ADCBlockedEvent             `json:"adcBlocked,omitempty" xml:"adcBlocked,omitempty" fake:"skip"`
}

type CommonBlockedEvent struct {
	Guid          string `json:"guid" xml:"guid" fake:"skip"`
	TimeslotBegin int64  `json:"timeslotBegin" xml:"timeslotBegin" fake:"skip"`
	TimeslotEnd   int64  `json:"timeslotEnd" xml:"timeslotEnd" fake:"skip"`
}

type AppExecBlockedEvent struct {
	CommonBlockedEvent
	File  string `json:"flie" xml:"flie" fake:"{threat-filename}"`
	Hash  string `json:"hash" xml:"hash" fake:"skip"`
	Type  string `json:"type" xml:"type" fake:"{randomstring:[Virus]}"`
	Name  string `json:"name" xml:"name" fake:"{randomstring:[PE_TEST_VIRUS]}"`
	Count int    `json:"count" xml:"count" fake:"{number:1,1000}"`
}

type FileScanBlockedEvent struct {
	CommonBlockedEvent
	File       string `json:"flie" xml:"flie" fake:"{threat-filename}"`
	Hash       string `json:"hash" xml:"hash" fake:"skip"`
	Type       string `json:"type" xml:"type" fake:"{randomstring:[Virus]}"`
	Name       string `json:"name" xml:"name" fake:"{randomstring:[PE_TEST_VIRUS]}"`
	Result     string `json:"result" xml:"result" fake:"{randomstring:[0]}"`
	Quarantine string `json:"quarantine" xml:"quarantine" fake:"{threat-quarantine}"`
	Count      int    `json:"count" xml:"count" fake:"{number:1,1000}"`
}

type SuspiciousExecBlockedEvent struct {
	Guid          string `json:"guid" xml:"guid" fake:"skip"`
	TimeslotBegin int64  `json:"timeslotBegin" xml:"timeslotBegin" fake:"skip"`
	TimeslotEnd   int64  `json:"timeslotEnd" xml:"timeslotEnd" fake:"skip"`
	File          string `json:"flie" xml:"flie" fake:"{threat-filename}"`
	Hash          string `json:"hash" xml:"hash" fake:"skip"`
	Count         int    `json:"count" xml:"count" fake:"{number:1,1000}"`
}

type OBADBlockedEvent struct {
	Guid          string `json:"guid" xml:"guid" fake:"skip"`
	TimeslotBegin int64  `json:"timeslotBegin" xml:"timeslotBegin" fake:"skip"`
	TimeslotEnd   int64  `json:"timeslotEnd" xml:"timeslotEnd" fake:"skip"`
	File          string `json:"flie" xml:"flie" fake:"{threat-filename}"`
	User          string `json:"user" xml:"user" fake:"{firstname}"`
	Parent1       string `json:"parent1" xml:"parent1" fake:"{threat-filename}"`
	Parent2       string `json:"parent2" xml:"parent2" fake:"{threat-filename}"`
	Parent3       string `json:"parent3" xml:"parent3" fake:"{threat-filename}"`
	Parent4       string `json:"parent4" xml:"parent4" fake:"{threat-filename}"`
	Mode          string `json:"mode" xml:"mode" fake:"{randomstring:[Detection,Prevention]}"`
	Level         string `json:"lvl" xml:"lvl" fake:"{randomstring:[aggressive]}"`
	Count         int    `json:"count" xml:"count" fake:"{number:1,1000}"`
}

type NonWhitelistingBlockedEvent struct {
	Guid          string `json:"guid" xml:"guid" fake:"skip"`
	TimeslotBegin int64  `json:"timeslotBegin" xml:"timeslotBegin" fake:"skip"`
	TimeslotEnd   int64  `json:"timeslotEnd" xml:"timeslotEnd" fake:"skip"`
	File          string `json:"flie" xml:"flie" fake:"{threat-filename}"`
	Hash          string `json:"hash" xml:"hash" fake:"skip"`
	User          string `json:"user" xml:"user" fake:"{firstname}"`
	Count         int    `json:"count" xml:"count" fake:"{number:1,1000}"`
}

type ADCBlockedEvent struct {
	Guid          string   `json:"guid" xml:"guid" fake:"skip"`
	TimeslotBegin int64    `json:"timeslotBegin" xml:"timeslotBegin" fake:"skip"`
	TimeslotEnd   int64    `json:"timeslotEnd" xml:"timeslotEnd" fake:"skip"`
	File          string   `json:"flie" xml:"flie" fake:"{threat-filename}"`
	Hash          string   `json:"hash" xml:"hash" fake:"skip"`
	Impacted      []string `json:"impacted" xml:"impacted" fake:"skip"`
	Mode          string   `json:"mode" xml:"mode" fake:"{randomstring:[Detection,Prevention]}"`
	Count         int      `json:"count" xml:"count" fake:"{number:1,1000}"`
}

// ThreatInfo will generate a struct of threat information
func ThreatInfo(min int) *Threat {
	return threatInfo(globalFaker.Rand, min)
}

func threatInfo(r *rand.Rand, min int) *Threat {
	var s Threat

	uuid := gofakeit.UUID()
	// remove dash from UUID
	uuid = strings.Replace(uuid, "-", "", -1)

	// AppExecBlockedEvents
	for nEvents := gofakeit.Number(min, 10); nEvents > 0; nEvents-- {
		event := appExecBlockedEventInfo(r, uuid)
		s.AppExecBlocked = append(s.AppExecBlocked, *event)
	}

	// FileScanBlockedEvent
	for nEvents := gofakeit.Number(min, 10); nEvents > 0; nEvents-- {
		event := fileScanBlockedEventInfo(r, uuid)
		s.FileScanBlocked = append(s.FileScanBlocked, *event)
	}

	return &s
}

// AppExecBlockedEventInfo will generate a struct of threat information
func AppExecBlockedEventInfo(uuid string) *AppExecBlockedEvent {
	return appExecBlockedEventInfo(globalFaker.Rand, uuid)
}

// FileScanBlockedEventInfo will generate a struct of threat information
func FileScanBlockedEventInfo(uuid string) *FileScanBlockedEvent {
	return fileScanBlockedEventInfo(globalFaker.Rand, uuid)
}

func appExecBlockedEventInfo(r *rand.Rand, uuid string) *AppExecBlockedEvent {
	var s AppExecBlockedEvent

	if err := gofakeit.Struct(&s); err != nil {
		panic(err)
	}

	s.setCommonBlockedEvent(uuid)

	// obtain file hash
	h := sha256.New()
	if _, err := h.Write([]byte(s.File)); err != nil {
		panic(err)
	}
	s.Hash = hex.EncodeToString(h.Sum(nil))

	return &s
}

func fileScanBlockedEventInfo(r *rand.Rand, uuid string) *FileScanBlockedEvent {
	var s FileScanBlockedEvent

	if err := gofakeit.Struct(&s); err != nil {
		panic(err)
	}

	s.setCommonBlockedEvent(uuid)

	// obtain file hash
	h := sha256.New()
	if _, err := h.Write([]byte(s.File)); err != nil {
		panic(err)
	}
	s.Hash = hex.EncodeToString(h.Sum(nil))

	return &s
}

func (event *CommonBlockedEvent) setCommonBlockedEvent(uuid string) {
	event.Guid = uuid

	// obtain time on the hour
	subHour := gofakeit.Number(-23, -1)
	fakeNow := time.Now().Add(time.Hour * time.Duration(subHour)).UTC()
	timestamp := fakeNow.Unix() - int64(fakeNow.Second()) - int64(60*fakeNow.Minute())
	event.TimeslotBegin = timestamp
	event.TimeslotEnd = timestamp + 86400

}

func threatFilename(r *rand.Rand) string {
	base := gofakeit.RandomString([]string{"C:", "C:\\Program Files (x86)", "C:\\Program Files", "C:\\Users", "C:\\Windows\\System32", "C:\\Windows\\SysWOW64"})
	folder := getRandValue(r, []string{"Application", "caption"})
	filename := folder
	return fmt.Sprintf("%s\\%s\\%s.exe", base, folder, filename)
}

func quarantineFilename(r *rand.Rand) string {
	base := gofakeit.RandomString([]string{"C:\\Program Files\\TXOne\\StellarProtect\\private\\quarantine", "C:\\Program Files\\TXOne\\StellarEnforce\\private\\quarantine"})
	filename := gofakeit.UUID()
	return fmt.Sprintf("%s\\%s", base, filename)
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

	gofakeit.AddFuncLookup("threat-quarantine", gofakeit.Info{
		Display:     "Quarantine Filename",
		Category:    "threat-quarantine",
		Description: "Random set of quarantine filename",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return quarantineFilename(r), nil
		},
	})
}
