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

// Declaring layout constant
const layout = "Jan 2, 2006 at 3:04pm (PST)"

type AppExecBlockedEvent struct {
	Guid    string    `json:"guid" xml:"guid" fake:"skip"`
	Created time.Time `json:"created" xml:"created" fake:"skip"`
	File    string    `json:"flie" xml:"flie" fake:"{threat-filename}"`
	Hash    string    `json:"hash" xml:"hash" fake:"skip"`
	Type    string    `json:"type" xml:"type" fake:"{randomstring:[Virus]}"`
	Name    string    `json:"name" xml:"name" fake:"{randomstring:[PE_TEST_VIRUS]}"`
}

// SystemEnv will generate a struct of system information
func AppExecBlockedEvents(min int, max int) []AppExecBlockedEvent {
	return appExecBlockedEvents(globalFaker.Rand, min, max)
}

func appExecBlockedEvents(r *rand.Rand, min int, max int) []AppExecBlockedEvent {
	var parsedCreated time.Time
	var err error
	if parsedCreated, err = time.Parse(layout, "Nov 1, 2021 at 0:00am (PST)"); err != nil {
		panic(err)
	}

	var list []AppExecBlockedEvent

	size := gofakeit.Number(min, max)

	uuid := gofakeit.UUID()
	// remove dash from UUID
	uuid = strings.Replace(uuid, "-", "", -1)

	for i := 0; i < size; i++ {
		h := sha256.New()
		addHour := gofakeit.Number(1, 3)
		created := parsedCreated.Add(time.Hour * time.Duration(addHour)).UTC()

		var s AppExecBlockedEvent
		if err := gofakeit.Struct(&s); err != nil {
			panic(err)
		}
		s.Guid = uuid
		s.Created = created

		if _, err := h.Write([]byte(s.File)); err != nil {
			panic(err)
		}
		s.Hash = hex.EncodeToString(h.Sum(nil))

		list = append(list, s)
	}

	return list
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
