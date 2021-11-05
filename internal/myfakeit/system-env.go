package myfakeit

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type SystemInfo struct {
	Guid    string       `json:"guid" xml:"guid" fake:"{uuid}"`
	Caption string       `json:"caption" xml:"caption" fake:"{os-caption}"`
	Version string       `json:"version" xml:"version" fake:"{os-version}"`
	Qfe     []QfeInfo    `json:"qfe" xml:"qfe" fakesize:"5"`
	Volume  []VolumeInfo `json:"volume" xml:"volume" fakesize:"2"`
	// Meta    []MetaInfo   `json:"meta" xml:"meta"`
}

type QfeInfo struct {
	HotfixId    string    `json:"hotfixid" xml:"hotfixid" fake:"{hotfixid}"`
	InstalledOn time.Time `json:"installedon" xml:"installedon" format:"MM/dd/yyyy"`
}

type VolumeInfo struct {
	Total string `json:"total" xml:"total" fake:"{number:2000000000000,4000000000000}"`
	Free  string `json:"free" xml:"free" fake:"{number:2000000000,4000000000}"`
	Type  string `json:"type" xml:"type" fake:"{randomstring:[FAT32,NTFS]}"`
}

// type MetaInfo struct {
// }

// System will generate a struct of system information
func SystemEnv() *SystemInfo { return systemEnv(globalFaker.Rand) }

func systemEnv(r *rand.Rand) *SystemInfo {
	var s SystemInfo
	err := gofakeit.Struct(&s)
	if err != nil {
		panic(err)
	}
	return &s
}

// OperatingSystemCaption will generate a random operating system string
func OperatingSystemCaption() string { return operatingSystemCaption(globalFaker.Rand) }

func operatingSystemCaption(r *rand.Rand) string {
	return getRandValue(r, []string{"OperatingSystem", "caption"})
}

// OperatingSystemVersion will generate a random operating system version string
func OperatingSystemVersion() string { return operatingSystemVersion(globalFaker.Rand) }

func operatingSystemVersion(r *rand.Rand) string {
	return getRandValue(r, []string{"OperatingSystem", "version"})
}

// OperatingSystemQfe will generate a random operating system qfe
func OperatingSystemQfe() string { return operatingSystemQfe(globalFaker.Rand) }

func operatingSystemQfe(r *rand.Rand) string {
	return "KB" + gofakeit.DigitN(7)
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

	gofakeit.AddFuncLookup("hotfixid", gofakeit.Info{
		Display:     "Hotfix ID",
		Category:    "hotfixid",
		Description: "Random hotfix ID",
		Example:     "KB5004331",
		Output:      "string",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return operatingSystemQfe(r), nil
		},
	})
}
