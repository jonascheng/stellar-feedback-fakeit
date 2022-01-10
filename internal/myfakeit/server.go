package myfakeit

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

type Server struct {
	Guid           string            `json:"guid" xml:"guid" fake:"{uuid}"`
	Version        string            `json:"version" xml:"version" fake:"{randomstring:[1.0.1000,1.1.1100]}"`
	AC             []string          `json:"ac" xml:"ac" fake:"{regex:[A-Z,0-9]{32}}" fakesize:"2"`
	EnabledFeature []string          `json:"enabledFeature" xml:"enabledFeature" fake:"skip"`
	Meta           map[string]string `json:"meta" xml:"meta" fake:"skip"`
}

type ServerVolumeInfo struct {
	Source string `json:"drive" xml:"drive" fake:"{randomstring:[/dev/sda1,/dev/mapper/vg_db-lvm_db,/dev/mapper/vg_redis-lvm_redis,/dev/mapper/vg_pattern-lvm_pattern,/dev/mapper/vg_sysinfo-lvm_sysinfo]}"`
	Total  string `json:"total" xml:"total" fake:"{number:2000000000000,4000000000000}"`
	Free   string `json:"free" xml:"free" fake:"{number:2000000000,4000000000}"`
	Type   string `json:"type" xml:"type" fake:"{randomstring:[ext3,ext4]}"`
}

// ServerInfo will generate a struct of server information
func ServerInfo() *Server { return serverInfo(globalFaker.Rand) }

func serverInfo(r *rand.Rand) *Server {
	var s Server
	if err := gofakeit.Struct(&s); err != nil {
		panic(err)
	}

	// remove dash from UUID
	s.Guid = strings.Replace(s.Guid, "-", "", -1)
	s.EnabledFeature = serverEnabledFeature(r)
	s.Meta = serverOperatingSystemMeta(r)

	return &s
}

// ServerEnabledFeature will generate a random enabled feature
// func ServerEnabledFeature() []string { return serverEnabledFeature(globalFaker.Rand) }

func serverEnabledFeature(r *rand.Rand) []string {
	var list []string
	if gofakeit.Bool() {
		list = append(list, "forward-syslog")
	}
	if gofakeit.Bool() {
		list = append(list, "schedule-report")
	}
	if gofakeit.Bool() {
		list = append(list, "schedule-update")
	}
	if gofakeit.Bool() {
		list = append(list, "send-warning-notification")
	}
	if gofakeit.Bool() {
		list = append(list, "send-outbreak-notification")
	}
	if gofakeit.Bool() {
		list = append(list, "smtp-authentication")
	}
	if gofakeit.Bool() {
		list = append(list, "self-certificate")
	}
	if gofakeit.Bool() {
		list = append(list, "proxy-s1-internet")
	}
	if gofakeit.Bool() {
		list = append(list, "proxy-s1-agent")
	}
	if gofakeit.Bool() {
		list = append(list, "proxy-agent-s1")
	}
	return list

}

// ServerOperatingSystemMeta will generate a random operating system meta
// func ServerOperatingSystemMeta() map[string]string { return serverOperatingSystemMeta(globalFaker.Rand) }

func serverOperatingSystemMeta(r *rand.Rand) map[string]string {
	dict := make(map[string]string)
	dict["cpuName"] = "Intel(R) Xeon(R) CPU E5-2609 0 @ 2.40GHz"
	dict["totalMemory"] = strconv.Itoa(gofakeit.Number(30000000000, 40000000000))
	dict["freeMemory"] = strconv.Itoa(gofakeit.Number(10000000, 20000000))
	dict["totalSeatCounts"] = strconv.Itoa(gofakeit.Number(10000, 20000))
	dict["freeSeatCounts"] = strconv.Itoa(gofakeit.Number(1000, 2000))
	return dict
}
