package main

import (
	"encoding/json"
	"fmt"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/builder"
)

func main() {
	agents := builder.CollectAgentSystemEnv()
	jsonBytes, err := json.Marshal(agents)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
