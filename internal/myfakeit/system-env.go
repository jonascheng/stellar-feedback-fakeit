package myfakeit

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

type SystemInfo struct {
	Guid   string               `json:"guid" xml:"guid"`
	OSInfo *OperatingSystemInfo `json:"os" xml:"os"`
}

type OperatingSystemInfo struct {
	Caption string `json:"caption" xml:"caption"`
	Version string `json:"version" xml:"version"`
}

// System will generate a struct of system information
func System() *SystemInfo { return system(globalFaker.Rand) }

func system(r *rand.Rand) *SystemInfo {
	guid := gofakeit.UUID()
	os := operatingSystem(r)

	return &SystemInfo{
		Guid:   guid,
		OSInfo: os,
	}
}

// OperatingSystem will generate a struct of operating system information
func OperatingSystem() *OperatingSystemInfo { return operatingSystem(globalFaker.Rand) }

func operatingSystem(r *rand.Rand) *OperatingSystemInfo {
	caption := operatingSystemCaption(r)

	return &OperatingSystemInfo{
		Caption: caption,
	}
}

// OperatingSystemCaption will generate a random operating system string
func OperatingSystemCaption() string { return operatingSystemCaption(globalFaker.Rand) }

func operatingSystemCaption(r *rand.Rand) string {
	return getRandValue(r, []string{"OperatingSystem", "caption"})
}

func addSystemLookup() {
	gofakeit.AddFuncLookup("system", gofakeit.Info{
		Display:     "System",
		Category:    "system",
		Description: "Random set of system information",
		Example:     "TODO",
		Output:      "TODO",
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			return system(r), nil
		},
	})
}
