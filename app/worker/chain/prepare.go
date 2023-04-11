package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	crossbellContract "github.com/Crossbell-Box/contracts.go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var (
	//!\\ Internal variables, please do not use them outside
	_initialized bool
	_client      *ethclient.Client
	_address     common.Address
	_auth        *bind.TransactOpts
	_instance    *crossbellContract.Contracts
)

func init() {
	// Package initialization
	_initialized = false
}

func update() error {
	if gasPrice, err := _client.SuggestGasPrice(context.Background()); err != nil {
		global.Logger.Errorf("Failed to get eth gas price with error: %s", err.Error())
		// Keep current gas price
	} else {
		// Update gas price
		_auth.GasPrice = gasPrice
	}

	nonce, err := _client.PendingNonceAt(context.Background(), _address)
	if err != nil {
		global.Logger.Errorf("Failed to get nonce with error: %s", err.Error())
		return err
	}
	_auth.Nonce = big.NewInt(int64(nonce))

	return nil
}

func Prepare() (*ethclient.Client, *crossbellContract.Contracts, *bind.TransactOpts, error) {
	if _initialized {
		// Already initialized, just update
		err := update()
		if err == nil {
			return _client, _instance, _auth, nil
		}
	}

	global.Logger.Debug("Start initializing Crossbell contract instance...")

	var err error

	_client, err = ethclient.Dial(config.Config.CrossbellJsonRPC)
	if err != nil {
		global.Logger.Errorf("Failed to connect to Crossbell JSON RPC with error: %s", err.Error())
		return nil, nil, nil, err
	}

	_instance, err = crossbellContract.NewContracts(common.HexToAddress(config.Config.CrossbellContractAddress), _client)
	if err != nil {
		global.Logger.Errorf("Failed to create Crossbell Contract instance with error: %s", err.Error())
		return nil, nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.Config.EthereumPrivateKey)
	if err != nil {
		global.Logger.Errorf("Failed to initialize Ethereum Private Key with error: %s", err.Error())
		return nil, nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		global.Logger.Errorf("Failed to cast publickey to ECDSA")
		return nil, nil, nil, fmt.Errorf("error casting public key to ECDSA")
	}

	_address = crypto.PubkeyToAddress(*publicKeyECDSA)

	global.Logger.Debugf("Operator address: %s", _address.Hex())

	_auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.Config.CrossbellChainID))
	if err != nil {
		global.Logger.Errorf("Failed to bind ethereum key with error: %s", err.Error())
		return nil, nil, nil, err
	}
	_auth.Value = big.NewInt(0)     // in wei
	_auth.GasLimit = uint64(300000) // in units

	err = update()
	if err != nil {
		return nil, nil, nil, err
	}

	global.Logger.Debug("Crossbell Ethereum account initialized successfully")
	_initialized = true
	return _client, _instance, _auth, nil
}
