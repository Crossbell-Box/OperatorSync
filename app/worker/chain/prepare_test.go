package chain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/consts"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"go.uber.org/zap"
	"math/big"
	"testing"
)

func TestPrepare(t *testing.T) {

	// Init settings
	config.Config.CrossbellChainID = consts.CONFIG_DEFAULT_CROSSBELL_CHAIN_ID
	config.Config.CrossbellJsonRPC = consts.CONFIG_DEFAULT_CROSSBELL_JSON_RPC
	config.Config.CrossbellContractAddress = consts.CONFIG_DEFAULT_CROSSBELL_CONTRACT_ADDRESS // Default
	randPrivKey, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	randPrivHex := hex.EncodeToString(crypto.FromECDSA(randPrivKey))
	t.Log("Random private key (Hex): ", randPrivHex)
	config.Config.EthereumPrivateKey = randPrivHex

	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // Unable to handle errors here
	global.Logger = logger.Sugar()

	// Prepare contract instance
	_, contractInstance, _, err := Prepare()
	if err != nil {
		global.Logger.Errorf("Failed to prepare eth contract instance")
		t.Fail()
	}

	// Test some read-only func
	result, err := contractInstance.GetCharacter(&bind.CallOpts{}, big.NewInt(19))
	resultJson, _ := json.MarshalIndent(&result, "", "  ")
	t.Log(string(resultJson))
}
