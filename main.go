package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	lp "github.com/remche/xhealth/liquiditypool"
	market "github.com/remche/xhealth/market"
	"gopkg.in/yaml.v2"
)

type CONFIG struct {
	Rpc    string `yaml:"rpc"`
	Lp     string `yaml:"lp"`
	Market string `yaml:"market"`
}

func readConfig(configPath string) CONFIG {
	config := CONFIG{}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}

func parseArgs(args []string) common.Address {
	if len(args) != 2 {
		log.Fatalf("Usage: %v address", args[0])
	}

	address := common.HexToAddress(os.Args[1])

	return address
}

func main() {
	config := readConfig("./config.yaml")
	account := parseArgs(os.Args)

	client, err := ethclient.Dial(config.Rpc)
	if err != nil {
		log.Fatal(err)
	}

	lpContractAddress := common.HexToAddress(config.Lp)
	marketContractAddress := common.HexToAddress(config.Market)
	lpContract, err := lp.NewLp(lpContractAddress, client)
	marketContract, err := market.NewMarket(marketContractAddress, client)

	debt, err := lpContract.UpdatedBorrowBy(nil, account)
	limit, err := marketContract.BorrowingLimit(nil, account)

	if err != nil {
		log.Fatal(err)
	}

	zero := big.NewInt(0)
	if debt.Cmp(zero) == 0 {
		fmt.Println("No debt for this address or incorrect address")
		os.Exit(123)
	}

	var health *big.Float
	if limit.Cmp(zero) == 0 {
		health = big.NewFloat(-1)
	} else {
		health = new(big.Float).Quo(new(big.Float).SetInt(limit), new(big.Float).SetInt(debt))
	}
	fmt.Println(health, debt, limit)
}
