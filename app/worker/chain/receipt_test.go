package chain

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/Crossbell-Box/OperatorSync/app/worker/config"
	"github.com/Crossbell-Box/OperatorSync/app/worker/consts"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"go.uber.org/zap"
	"testing"
)

func TestReceipt(t *testing.T) {
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
	client, _, _, _ := Prepare()

	// Test
	rcpt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0xee78c9061159b17eeac2f30f19bfc7dc5b40dac0d9f6f7d681909ef5d546c596"))
	if err != nil {
		t.Logf("Error occurred: %v", err)
		t.Fail()
	} else {
		t.Logf("%+v", rcpt)
		for _, log := range rcpt.Logs {
			t.Logf("%+v", log)
			for _, topic := range log.Topics {
				t.Log(topic.Big().Int64())
			}

		}

	}
}
