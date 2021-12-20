package myfakeit

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

type AgentSoftwareEnv struct {
	Guid string    `json:"guid" xml:"guid" fake:"{uuid}"`
	App  []AppInfo `json:"app" xml:"app" fakesize:"1000"`
}

type AppInfo struct {
	Caption         string `json:"caption" xml:"caption" fake:"{app-caption}"`
	Version         string `json:"version" xml:"version" fake:"{app-version}"`
	Vendor          string `json:"vendor" xml:"vendor" fake:"skip"`
	InstallLocation string `json:"installLocation" xml:"installLocation" fake:"{app-installlocation}"`
}

// SoftwareEnv will generate a struct of software information
func SoftwareEnv() *AgentSoftwareEnv { return sfotwareEnv(globalFaker.Rand) }

func sfotwareEnv(r *rand.Rand) *AgentSoftwareEnv {
	var s AgentSoftwareEnv
	if err := gofakeit.Struct(&s); err != nil {
		panic(err)
	}

	// remove dash from UUID
	s.Guid = strings.Replace(s.Guid, "-", "", -1)

	// set vendor to caption
	for i, app := range s.App {
		s.App[i].Vendor = app.Caption
	}

	return &s
}

// ApplicationCaption will generate a random application caption string
// func ApplicationCaption() string { return applicationCaption(globalFaker.Rand) }

func applicationCaption(r *rand.Rand) string {
	n := gofakeit.Number(0, 1)
	if n%2 == 0 {
		return getRandValue(r, []string{"Application", "caption"})
	} else {
		return gofakeit.Company()
	}
}

// ApplicationVersion will generate a random application version string
// func ApplicationVersion() string { return applicationVersion(globalFaker.Rand) }

func applicationVersion(r *rand.Rand) string {
	major := gofakeit.Number(1, 3)
	minor := gofakeit.Number(1, 3)
	build := 1000

	return fmt.Sprintf("%d.%d.%d", major, minor, build)
}

// ApplicationInstallLocation will generate a random application version string
// func ApplicationInstallLocation() string { return applicationInstallLocation(globalFaker.Rand) }

func applicationInstallLocation(r *rand.Rand) string {
	base := gofakeit.RandomString([]string{"C:\\Program Files"})
	caption := applicationCaption(r)
	return fmt.Sprintf("%s\\%s\\", base, caption)
}

func addSoftwareLookup() {
	gofakeit.AddFuncLookup("software", gofakeit.Info{
		Display:     "Software",
		Category:    "software",
		Description: "Random set of software information",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return sfotwareEnv(r), nil
		},
	})

	gofakeit.AddFuncLookup("app-caption", gofakeit.Info{
		Display:     "App Caption",
		Category:    "app-caption",
		Description: "Random App Caption",
		Example:     "Schneider Electric LinkManager",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return applicationCaption(r), nil
		},
	})

	gofakeit.AddFuncLookup("app-version", gofakeit.Info{
		Display:     "App Version",
		Category:    "app-version",
		Description: "Random App Version",
		Example:     "10.0.19042",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return applicationVersion(r), nil
		},
	})

	gofakeit.AddFuncLookup("app-installlocation", gofakeit.Info{
		Display:     "App Install Location",
		Category:    "app-installlocation",
		Description: "Random App Install Location",
		Example:     "C:\\Program Files\\Schneider Electric LinkManager\\",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return applicationInstallLocation(r), nil
		},
	})
}
