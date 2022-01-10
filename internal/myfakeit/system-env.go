package myfakeit

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type AgentSystemEnv struct {
	Guid    string             `json:"guid" xml:"guid" fake:"{uuid}"`
	Caption string             `json:"caption" xml:"caption" fake:"{os-caption}"`
	Version string             `json:"version" xml:"version" fake:"{os-version}"`
	Qfe     []QfeInfo          `json:"qfe" xml:"qfe" fakesize:"5"`
	Volume  []VolumeInfo       `json:"volume" xml:"volume" fakesize:"2"`
	Meta    *map[string]string `json:"meta" xml:"meta" fake:"skip"`
}

type OSInfo struct {
	Caption string `json:"caption" xml:"caption" fake:"{os-caption}"`
	Version string `json:"version" xml:"version" fake:"{os-version}"`
}

type QfeInfo struct {
	HotfixId    string `json:"hotfixid" xml:"hotfixid" fake:"{qfe-hotfixid}"`
	InstalledOn string `json:"installedon" xml:"installedon" fake:"{qfe-installedon}"`
}

type VolumeInfo struct {
	Drive string `json:"drive" xml:"drive" fake:"{randomstring:[C:,D:]}"`
	Total string `json:"total" xml:"total" fake:"{number:2000000000000,4000000000000}"`
	Free  string `json:"free" xml:"free" fake:"{number:2000000000,4000000000}"`
	Type  string `json:"type" xml:"type" fake:"{randomstring:[FAT32,NTFS]}"`
}

// SystemEnv will generate a struct of system information
func SystemEnv() *AgentSystemEnv { return systemEnv(globalFaker.Rand) }

func systemEnv(r *rand.Rand) *AgentSystemEnv {
	var s AgentSystemEnv
	if err := gofakeit.Struct(&s); err != nil {
		panic(err)
	}

	// remove dash from UUID
	s.Guid = strings.Replace(s.Guid, "-", "", -1)

	s.Meta = operatingSystemMeta(r)

	return &s
}

// OperatingSystemCaption will generate a random operating system string
// func OperatingSystemCaption() string { return operatingSystemCaption(globalFaker.Rand) }

func operatingSystemCaption(r *rand.Rand) string {
	return getRandValue(r, []string{"OperatingSystem", "caption"})
}

// OperatingSystemVersion will generate a random operating system version string
// func OperatingSystemVersion() string { return operatingSystemVersion(globalFaker.Rand) }

func operatingSystemVersion(r *rand.Rand) string {
	return getRandValue(r, []string{"OperatingSystem", "version"})
}

// OperatingSystemQfe will generate a random operating system qfe
// func OperatingSystemQfe() string { return operatingSystemQfe(globalFaker.Rand) }

func operatingSystemQfe(r *rand.Rand) string {
	return "KB" + gofakeit.DigitN(7)
}

// OperatingSystemQfeInstalledOn will generate a random operating system qfe
// func OperatingSystemQfeInstalledOn() string { return operatingSystemQfeInstalledOn(globalFaker.Rand) }

func operatingSystemQfeInstalledOn(r *rand.Rand) string {
	// Declaring layout constant
	const layout = "02/01/2006"

	startdate, err := time.Parse(layout, "01/01/2021")
	if err != nil {
		panic(err)
	}
	enddate, _ := time.Parse(layout, "31/12/2021")
	if err != nil {
		panic(err)
	}
	installon := gofakeit.DateRange(startdate, enddate)
	return installon.Format(layout)
}

// OperatingSystemMeta will generate a random operating system meta
// func OperatingSystemMeta() *map[string]string { return operatingSystemMeta(globalFaker.Rand) }

func operatingSystemMeta(r *rand.Rand) *map[string]string {
	var dict = make(map[string]string)
	dict["cpuCaption"] = "Intel64 Family 6 Model 167 Stepping 1"
	dict["cpuName"] = "11th Gen Intel(R) Core(TM) i7-11700 @ 2.50GHz"
	dict["cpuArchitecture"] = strconv.Itoa(gofakeit.RandomInt([]int{0, 5, 9, 12}))
	dict["osArchitecture"] = gofakeit.RandomString([]string{"64-bit", "32-bit"})
	dict["osLanguage"] = gofakeit.RandomString([]string{"1033", "1028", "1041"})
	dict["totalMemory"] = strconv.Itoa(gofakeit.Number(30000000000, 40000000000))
	dict["freeMemory"] = strconv.Itoa(gofakeit.Number(10000000, 20000000))
	return &dict
}

func addSystemLookup() {
	gofakeit.AddFuncLookup("system", gofakeit.Info{
		Display:     "System",
		Category:    "system",
		Description: "Random set of system information",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return systemEnv(r), nil
		},
	})

	gofakeit.AddFuncLookup("os-caption", gofakeit.Info{
		Display:     "OS Caption",
		Category:    "os-caption",
		Description: "Random OS Caption",
		Example:     "Microsoft Windows 10 Enterprise",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return operatingSystemCaption(r), nil
		},
	})

	gofakeit.AddFuncLookup("os-version", gofakeit.Info{
		Display:     "OS Version",
		Category:    "os-version",
		Description: "Random OS Version",
		Example:     "10.0.19042",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return operatingSystemVersion(r), nil
		},
	})

	gofakeit.AddFuncLookup("qfe-hotfixid", gofakeit.Info{
		Display:     "Hotfix ID",
		Category:    "qfe-hotfixid",
		Description: "Random hotfix ID",
		Example:     "KB5004331",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return operatingSystemQfe(r), nil
		},
	})

	gofakeit.AddFuncLookup("qfe-installedon", gofakeit.Info{
		Display:     "Hotfix Install On",
		Category:    "qfe-installedon",
		Description: "Random hotfix Install On",
		Example:     "8/10/2021",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return operatingSystemQfeInstalledOn(r), nil
		},
	})
}
