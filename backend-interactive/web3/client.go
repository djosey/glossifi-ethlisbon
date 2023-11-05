package client

import (
	"ethlisbon-2023/configs"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Client() (client *ethclient.Client) {
	ETH_CLIENT := configs.GoDotEnvVariable("ETHCLIENT")
	client, err := ethclient.Dial(ETH_CLIENT)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		return
	}
	return client
}
