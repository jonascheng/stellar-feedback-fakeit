package builder

import (
	"github.com/brianvoe/gofakeit/v6"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/myfakeit"
)

type AgentSystemEnv struct {
	SystemInfo myfakeit.SystemInfo `json:"system" xml:"system" fake:"{system}"`
}

type AgentSystemEnvCollection struct {
	Agents []AgentSystemEnv `json:"agents" xml:"agents" fakesize:"3"`
}

func CollectAgentSystemEnv() AgentSystemEnvCollection {
	var agents AgentSystemEnvCollection
	err := gofakeit.Struct(&agents)
	if err != nil {
		panic(err)
	}
	return agents
}
