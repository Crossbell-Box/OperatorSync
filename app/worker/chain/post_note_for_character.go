package chain

import (
	"context"
	"errors"
	"fmt"
	crossbellContract "github.com/Crossbell-Box/OperatorSync/app/worker/chain/contract"
	"github.com/Crossbell-Box/OperatorSync/app/worker/consts"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
)

func PostNoteForCharacter(characterIdStr string, metadataUri string, forURI string, forNoteCharacterId int64, forNoteNoteId int64) (string, error) {
	characterId, err := strconv.ParseInt(characterIdStr, 10, 64)
	if err != nil {
		global.Logger.Errorf("Failed to parse character id with error: %s", err.Error())
		return "", err
	}

	// Prepare contract instance
	client, contractInstance, operatorAuth, err := Prepare()
	if err != nil {
		global.Logger.Errorf("Failed to prepare eth contract instance")
		return "", err
	}

	var tx *ethTypes.Transaction
	noteData := crossbellContract.DataTypesPostNoteData{
		CharacterId: big.NewInt(characterId),
		ContentUri:  metadataUri,
	}
	if forURI != "" {
		if forNoteCharacterId > 0 && forNoteNoteId > 0 {
			// For note
			tx, err = contractInstance.PostNote4Note(
				operatorAuth,
				noteData,
				crossbellContract.DataTypesNoteStruct{
					CharacterId: big.NewInt(forNoteCharacterId),
					NoteId:      big.NewInt(forNoteNoteId),
				},
			)
		} else {
			// For normal URI
			tx, err = contractInstance.PostNote4AnyUri(
				operatorAuth,
				noteData,
				forURI,
			)
		}
	} else {
		tx, err = contractInstance.PostNote(
			operatorAuth,
			noteData,
		)
	}
	if err != nil {
		global.Logger.Errorf("Failed to create transaction: %s", err.Error())
		return "", err
	}

	// Wait till transaction finish or timeout
	ctx, cancel := context.WithTimeout(context.Background(), consts.CONFIG_TRANSACTION_TIMEOUT)
	defer cancel()
	rcpt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			global.Logger.Errorf("Transaction %s timed out!", tx.Hash().Hex())
		}
		return "", err
	}
	if rcpt.Status == ethTypes.ReceiptStatusFailed {
		global.Logger.Errorf("Transaction %s failed", tx.Hash().Hex())
		return "", fmt.Errorf("transaction failed")
	}

	return tx.Hash().Hex(), nil
}
