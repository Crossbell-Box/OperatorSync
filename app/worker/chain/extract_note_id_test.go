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

func TestExtractNoteID(t *testing.T) {
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

	// Notes pending test
	tx2NoteID := map[string]int64{
		"0x4c2283c9ed7c775c3be59ef9e32b32939b971545c057412d723b4c7705ac2e2b": 138, // postNote
		"0x2cf5e952aadea7254819f90904ec34be46c7a7a8e2aba0ae9c01d7b0655c0026": 260, //postNote4AnyUri
		"0xe846dcec2849663b8cc9064b0982788be835e21db045e7d0b16ec6f4f36a2a3a": 270,
	}

	// Test
	for tx, wantedNoteId := range tx2NoteID {
		rcpt, err := client.TransactionReceipt(context.Background(), common.HexToHash(tx))
		if err != nil {
			t.Logf("Failed to get transaction: %v", err)
			continue
		}
		extractedNoteId, err := ExtractNoteID(rcpt)
		if err != nil {
			t.Logf("Failed to extract note id: %v", err)
			continue
		}

		if extractedNoteId != wantedNoteId {
			t.Logf("Invalid note id: want %d, got %d", wantedNoteId, extractedNoteId)
		} else {
			t.Logf("Note id matches.")
		}
	}

}
