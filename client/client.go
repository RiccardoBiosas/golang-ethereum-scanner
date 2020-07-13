package client

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Client struct {
	EthereumClient *ethclient.Client
}

func(c *Client) Mount() {
	notFound := godotenv.Load()
		if notFound != nil {
			log.Fatal("No .env file found")
		}
		infuraRopstenApiKey, exists := os.LookupEnv("INFURA_ROPSTEN_API_KEY")
		if !exists {
			log.Fatal("No infura ropsten api key")
		}
		var err error
		c.EthereumClient, err = ethclient.Dial(infuraRopstenApiKey)
		if err != nil {
			log.Fatal(err)
		}
}

