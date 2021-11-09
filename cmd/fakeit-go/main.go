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
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/factory"
	"github.com/jszwec/csvutil"

	// A Go (golang) command line and flag parser
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	serverInfo       = kingpin.Flag("server-info", "Random generate server-telemetry.").Bool()
	agentInfo        = kingpin.Flag("agent-info", "Random generate agent-telemetry.").Bool()
	agentSystemEnv   = kingpin.Flag("agent-system-env", "Random generate agent-telemetry-system-environment.").Bool()
	agentSoftwareEnv = kingpin.Flag("agent-software-env", "Random generate agent-telemetry-software-environment.").Bool()
	agentCert        = kingpin.Flag("agent-cert", "Random generate agent-telemetry-cert.").Bool()
	benchmark        = kingpin.Flag("benchmark", "Benchmark performance.").Bool()
	size             = kingpin.Flag("size", "Random size").Default("1").Int()
	debug            = kingpin.Flag("debug", "Debug output results in json format").Bool()
)

type Benchmark struct {
	Size                  int     `json:"fakesize" xml:"fakesize" csv:"fakesize"`
	SizeFlat              int     `json:"sizeFlat" xml:"sizeFlat" csv:"sizeFlat"`
	SizeLookup            int     `json:"sizeLookup" xml:"sizeLookup" csv:"sizeLookup"`
	SizeFlatCompressed    int     `json:"sizeFlatCompressed" xml:"sizeFlatCompressed" csv:"sizeFlatCompressed"`
	RatioFlatCompressed   float32 `json:"ratioFlatCompressed" xml:"ratioFlatCompressed" csv:"ratioFlatCompressed"`
	SizeLookupCompressed  int     `json:"sizeLookupCompressed" xml:"sizeLookupCompressed" csv:"sizeLookupCompressed"`
	RatioLookupCompressed float32 `json:"ratioLookupCompressed" xml:"ratioLookupCompressed" csv:"ratioLookupCompressed"`
}

func encodeCollectionFlat(benchmark *Benchmark, agents factory.IAgentTelemetry, flatFilename string) {
	telemetryFlat := agents.EncodeCollectionFlat()
	benchmark.SizeFlat = dumpToFile(telemetryFlat, flatFilename)
	benchmark.SizeFlatCompressed = compressFile(flatFilename)
}

func encodeAgentCollectionLookup(benchmark *Benchmark, agents factory.IAgentTelemetry, lookupFilename string) {
	telemetryFlat := agents.EncodeAgentCollectionLookup()
	benchmark.SizeLookup = dumpToFile(telemetryFlat, lookupFilename)
	benchmark.SizeLookupCompressed = compressFile(lookupFilename)
}

func fullServerCollection() Benchmark {
	var benchmark Benchmark

	agents := factory.NewServerCollection()
	benchmark.Size = 0

	flatFilename := "server-telemetry-info-flat.json"
	encodeCollectionFlat(&benchmark, agents, flatFilename)

	return benchmark
}

func fullAgentCollection(size int) Benchmark {
	var benchmark Benchmark

	agents := factory.NewAgentCollection(size)
	benchmark.Size = size

	flatFilename := fmt.Sprintf("agent-telemetry-info-flat-%d.json", size)
	encodeCollectionFlat(&benchmark, agents, flatFilename)

	return benchmark
}

func fullAgentSystemEnvCollection(size int) Benchmark {
	var benchmark Benchmark

	agents := factory.NewAgentSystemEnvCollection(size)
	benchmark.Size = size

	flatFilename := fmt.Sprintf("agent-telemetry-system-environment-flat-%d.json", size)
	encodeCollectionFlat(&benchmark, agents, flatFilename)

	lookupFilename := fmt.Sprintf("agent-telemetry-system-environment-lookup-%d.json", size)
	encodeAgentCollectionLookup(&benchmark, agents, lookupFilename)

	return benchmark
}

func fullAgentSoftwareEnvCollection(size int) Benchmark {
	var benchmark Benchmark

	agents := factory.NewAgentSoftwareEnvCollection(size)
	benchmark.Size = size

	flatFilename := fmt.Sprintf("agent-telemetry-software-environment-flat-%d.json", size)
	encodeCollectionFlat(&benchmark, agents, flatFilename)

	lookupFilename := fmt.Sprintf("agent-telemetry-software-environment-lookup-%d.json", size)
	encodeAgentCollectionLookup(&benchmark, agents, lookupFilename)

	return benchmark
}

func fullAgentCertCollection(size int) Benchmark {
	var benchmark Benchmark

	agents := factory.NewAgentCertCollection(size)
	benchmark.Size = size

	flatFilename := fmt.Sprintf("agent-telemetry-cert-flat-%d.json", size)
	encodeCollectionFlat(&benchmark, agents, flatFilename)

	lookupFilename := fmt.Sprintf("agent-telemetry-cert-lookup-%d.json", size)
	encodeAgentCollectionLookup(&benchmark, agents, lookupFilename)

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

func benchmarkAgentTelemetry(callback func(size int) Benchmark) {
	quit := make(chan bool)

	fmt.Println(">>")
	funcName := runtime.FuncForPC(reflect.ValueOf(callback).Pointer()).Name()
	fmt.Printf("Benchmark %v\n", funcName)
	go dots(quit)

	var benchmarkResult []Benchmark
	for i := 50; i < 10000; i = i * 2 {
		benchmark := callback(i)

		// calculate compression ratio
		benchmark.RatioFlatCompressed = float32(benchmark.SizeFlat) / float32(benchmark.SizeFlatCompressed)
		benchmark.RatioLookupCompressed = float32(benchmark.SizeLookup) / float32(benchmark.SizeLookupCompressed)

		benchmarkResult = append(benchmarkResult, benchmark)
	}

	// output result in csv
	csvBytes, err := csvutil.Marshal(benchmarkResult)
	quit <- true
	checkError(err)
	fmt.Println("")
	fmt.Println(string(csvBytes))
	fmt.Println("<<")
}

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	if *benchmark {
		benchmarkAgentTelemetry(fullAgentSystemEnvCollection)
		benchmarkAgentTelemetry(fullAgentSoftwareEnvCollection)
		benchmarkAgentTelemetry(fullAgentCertCollection)

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

	if *serverInfo {
		fmt.Println(">>")
		fmt.Println("Generate server collection")
		fullServerCollection()
		fmt.Println("<<")
	}

	if *agentInfo {
		fmt.Println(">>")
		fmt.Printf("Generate agent collection with size %v\n", *size)
		fullAgentCollection(*size)
		fmt.Println("<<")
	}

	if *agentSystemEnv {
		fmt.Println(">>")
		fmt.Printf("Generate agent system environment collection with size %v\n", *size)
		fullAgentSystemEnvCollection(*size)
		fmt.Println("<<")
	}

	if *agentSoftwareEnv {
		fmt.Println(">>")
		fmt.Printf("Generate agent software environment collection with size %v\n", *size)
		fullAgentSoftwareEnvCollection(*size)
		fmt.Println("<<")
	}

	if *agentCert {
		fmt.Println(">>")
		fmt.Printf("Generate agent cert collection with size %v\n", *size)
		fullAgentCertCollection(*size)
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
