package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var (
	initialized bool
	client      *ethclient.Client
	auth        *bind.TransactOpts
)

func init() {
	// Package initialization
	initialized = false
}

func Prepare() (*ethclient.Client, *bind.TransactOpts, error) {
	if initialized {
		// Already initialized, just update gas price
		if gasPrice, err := client.SuggestGasPrice(context.Background()); err != nil {
			global.Logger.Errorf("Failed to get eth gas price with error: %s", err.Error())
			// Keep current gas price
		} else {
			// Update gas price
			auth.GasPrice = gasPrice
		}

		return client, auth, nil
	}

	var err error

	client, err = ethclient.Dial(config.Config.CrossbellJsonRPC)
	if err != nil {
		global.Logger.Errorf("Failed to connect to Crossbell JSON RPC with error: %s", err.Error())
		return nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.Config.EthereumPrivateKey)
	if err != nil {
		global.Logger.Errorf("Failed to initialize Ethereum Private Key with error: %s", err.Error())
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		global.Logger.Errorf("Failed to cast publickey to ECDSA")
		return nil, nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		global.Logger.Errorf("Failed to get nonce with error: %s", err.Error())
		return nil, nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		global.Logger.Errorf("Failed to get eth gas price with error: %s", err.Error())
		return nil, nil, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.Config.CrossbellChainID))
	if err != nil {
		global.Logger.Errorf("Failed to bind ethereum key with error: %s", err.Error())
		return nil, nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	global.Logger.Debug("Crossbell Ethereum account initialized successfully")
	initialized = true
	return client, auth, nil
}
