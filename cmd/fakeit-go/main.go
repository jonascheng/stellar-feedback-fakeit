package main

import (
	"encoding/json"
	"fmt"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/builder"
	// A Go (golang) command line and flag parser
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	showList   = kingpin.Flag("list", "Show available speedtest.net servers.").Short('l').Bool()
	serverIds  = kingpin.Flag("server", "Select server id to speedtest.").Short('s').Ints()
	savingMode = kingpin.Flag("saving-mode", "Using less memory (≒10MB), though low accuracy (especially > 30Mbps).").Bool()
	jsonOutput = kingpin.Flag("json", "Output results in json format").Bool()
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	agents := builder.CollectAgentSystemEnv(1)
	jsonBytes, err := json.Marshal(agents)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))

	telemetryFlat := builder.EncodeAgentSystemEnvCollectionFlat(agents)
	jsonBytes, err = json.Marshal(telemetryFlat)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))

	telemetryLookup := builder.EncodeAgentSystemEnvCollectionLookup(agents)
	jsonBytes, err = json.Marshal(telemetryLookup)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
