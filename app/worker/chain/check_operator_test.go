package chain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/consts"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"go.uber.org/zap"
	"testing"
)

func TestCheckOperator(t *testing.T) {
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

	// Test
	t.Log(CheckOperator("451"))
}
