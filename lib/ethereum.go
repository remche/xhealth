package lib

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/remche/xhealth/contracts"
)

func Query(config *Config) *big.Float {

	client, err := ethclient.Dial(config.Rpc)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(config.Address)
	lpContractAddress := common.HexToAddress(config.LpContract)
	marketContractAddress := common.HexToAddress(config.MarketContract)
	lpContract, err := contracts.NewLp(lpContractAddress, client)
	marketContract, err := contracts.NewMarket(marketContractAddress, client)

	debt, err := lpContract.UpdatedBorrowBy(nil, address)
	limit, err := marketContract.BorrowingLimit(nil, address)

	if err != nil {
		log.Fatal(err)
	}

	zero := big.NewInt(0)
	if debt.Cmp(zero) == 0 {
		fmt.Println(address)
		fmt.Println("No debt for this address or incorrect address")
		os.Exit(123)
	}

	var health *big.Float
	if limit.Cmp(zero) == 0 {
		health = big.NewFloat(-1)
	} else {
		health = new(big.Float).Quo(new(big.Float).SetInt(limit), new(big.Float).SetInt(debt))
	}

	return (health)
}
