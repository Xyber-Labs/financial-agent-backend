package agent_client

import "math/big"

type AgentClient struct {
}

func NewAgentClient() *AgentClient {
	return &AgentClient{}
}

type RequestActionResponse struct {
	PoolAddress string `json:"poolAddress"`
}

type IAgentClient interface {
	RequestAction(tokenAddress string, amount *big.Int, chainId *big.Int) (*RequestActionResponse, error)
}
