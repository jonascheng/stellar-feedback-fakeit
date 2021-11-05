package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/builder"
	"github.com/jszwec/csvutil"

	// A Go (golang) command line and flag parser
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	agentSystemEnv = kingpin.Flag("agent-system-env", "Random generate agent-telemetry-system-environment.").Bool()
	benchmark      = kingpin.Flag("benchmark", "Benchmark performance.").Bool()
	size           = kingpin.Flag("size", "Random size").Default("1").Int()
	debug          = kingpin.Flag("debug", "Debug output results in json format").Bool()
)

type Benchmark struct {
	Size                 int `json:"fakesize" xml:"fakesize" csv:"fakesize"`
	SizeFlat             int `json:"sizeFlat" xml:"sizeFlat" csv:"sizeFlat"`
	SizeLookup           int `json:"sizeLookup" xml:"sizeLookup" csv:"sizeLookup"`
	SizeFlatCompressed   int `json:"sizeFlatCompressed" xml:"sizeFlatCompressed" csv:"sizeFlatCompressed"`
	SizeLookupCompressed int `json:"sizeLookupCompressed" xml:"sizeLookupCompressed" csv:"sizeLookupCompressed"`
}

func getAgentSystemEnvCollection(size int) *builder.AgentSystemEnvCollection {
	agents := builder.CollectAgentSystemEnv(size)
	return agents
}

func encodeAgentSystemEnvCollectionFlat(agents *builder.AgentSystemEnvCollection) *builder.AgentTelemetrySystemEnvFlat {
	telemetry := builder.EncodeAgentSystemEnvCollectionFlat(agents)
	return telemetry
}

func encodeAgentSystemEnvCollectionLookup(agents *builder.AgentSystemEnvCollection) *builder.AgentTelemetrySystemEnvLookup {
	telemetry := builder.EncodeAgentSystemEnvCollectionLookup(agents)
	return telemetry
}

func fullAgentSystemEnvCollection(size int) Benchmark {
	var benchmark Benchmark

	agents := getAgentSystemEnvCollection(size)
	benchmark.Size = size

	telemetryFlat := encodeAgentSystemEnvCollectionFlat(agents)
	flatFilename := fmt.Sprintf("agent-telemetry-system-environment-flat-%d.json", size)
	benchmark.SizeFlat = dumpToFile(telemetryFlat, flatFilename)
	benchmark.SizeFlatCompressed = compressFile(flatFilename)

	lookupFilename := fmt.Sprintf("agent-telemetry-system-environment-lookup-%d.json", size)
	telemetryLookup := encodeAgentSystemEnvCollectionLookup(agents)
	benchmark.SizeLookup = dumpToFile(telemetryLookup, lookupFilename)
	benchmark.SizeLookupCompressed = compressFile(lookupFilename)

	return benchmark
}

func dumpToFile(v interface{}, filename string) int {
	// encode json
	jsonBytes, err := json.Marshal(v)
	checkError(err)
	// write file
	f, err := os.Create(filename)
	checkError(err)
	defer f.Close()
	sizeWriten, err := f.WriteString(string(jsonBytes))
	checkError(err)

	if *debug {
		checkError(err)
		fmt.Println(string(jsonBytes))
	}
	return sizeWriten
}

func compressFile(source string) int {
	// open source file
	fSource, err := os.Open(source)
	checkError(err)
	defer fSource.Close()

	// Now let use read the bytes of the document we opened.
	// Create a Reader to get all the bytes from the file.
	reader := bufio.NewReader(fSource)

	// Now we would use the variable Read All to get all the bytes
	// So we just used variable data which will read all the bytes
	data, err := ioutil.ReadAll(reader)
	checkError(err)

	// Now we would use the extension method
	// Now with the help of replace command we can
	// Replace json file with gz extension
	// So we would now use the file name to give
	// this command a boost
	target := strings.Replace(source, ".json", ".json.gz", -1)

	// Open target file
	fTarget, err := os.Create(target)
	checkError(err)
	defer fTarget.Close()

	// Write compresses Data
	// We would use NewWriter to basically
	// copy all the compressed data
	writer := gzip.NewWriter(fTarget)
	defer writer.Close()

	// With the help of the Writer method, we would
	// write all the bytes in the data variable
	// copied from the original file
	_, err = writer.Write(data)
	checkError(err)
	writer.Flush()
	writer.Close()

	fi, err := fTarget.Stat()
	checkError(err)

	return int(fi.Size())
}

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	if *benchmark {
		quit := make(chan bool)

		fmt.Println(">>")
		fmt.Println("Benchmark getAgentSystemEnvCollection")
		go dots(quit)

		var benchmarkResult []Benchmark
		for i := 50; i < 10000; i = i * 2 {
			benchmark := fullAgentSystemEnvCollection(i)
			benchmarkResult = append(benchmarkResult, benchmark)
		}

		// output result in csv
		csvBytes, err := csvutil.Marshal(benchmarkResult)
		quit <- true
		checkError(err)
		fmt.Println("")
		fmt.Println(string(csvBytes))
		fmt.Println("<<")

		// clean up benchmark files
		files, err := filepath.Glob("*.json")
		checkError(err)
		for _, f := range files {
			os.Remove(f)
		}
		files, err = filepath.Glob("*.gz")
		checkError(err)
		for _, f := range files {
			os.Remove(f)
		}
		return
	}

	if *agentSystemEnv {
		fmt.Println(">>")
		fmt.Printf("Generate agent system environment collection with size %v\n", *size)
		fullAgentSystemEnvCollection(*size)
		fmt.Println("<<")
	}
}

func dots(quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Second)
			fmt.Print(".")
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
