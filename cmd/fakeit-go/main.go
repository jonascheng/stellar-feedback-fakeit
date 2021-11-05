package main

import (
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentSystemEnv struct {
	SystemEnv myfakeit.SystemEnv `json:"system" xml:"system" fake:"{system}"`
	// AddressInfo *gofakeit.AddressInfo `fake:"{address}"`
	// PersonInfo  *gofakeit.PersonInfo  `fake:"{person}"`
	// UUID string `fake:"{uuid}"`
	// CreatedFormat time.Time `format:"MM/dd/yyyy"`
}

type BulkAgentSystemEnv struct {
	Agents []AgentSystemEnv `json:"agents" xml:"agents" fakesize:"3"`
}

func testStruct() {
	var agents BulkAgentSystemEnv
	err := gofakeit.Struct(&agents)
	if err != nil {
		panic(err)
	}

	x, err := json.Marshal(agents)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(x))
}

func main() {
	testStruct()
}
