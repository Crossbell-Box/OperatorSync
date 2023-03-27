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

func PostNoteForCharacter(characterIdStr string, metadataUri string, forURI string, forNoteCharacterId int64, forNoteNoteId int64) (string, int64, int64, error) {
	characterId, err := strconv.ParseInt(characterIdStr, 10, 64)
	if err != nil {
		global.Logger.Errorf("Failed to parse character id with error: %s", err.Error())
		return "", 0, 0, err
	}

	// Prepare contract instance
	client, contractInstance, operatorAuth, err := Prepare()
	if err != nil {
		global.Logger.Errorf("Failed to prepare eth contract instance")
		return "", 0, 0, err
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
		return "", 0, 0, err
	}

	// Wait till transaction finish or timeout
	ctx, cancel := context.WithTimeout(context.Background(), consts.CONFIG_TRANSACTION_TIMEOUT)
	defer cancel()
	rcpt, err := bind.WaitMined(ctx, client, tx) // Warn: Not stable -> for other evm chains, we should query this later (like for transaction revert etc.)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			global.Logger.Errorf("Transaction %s timed out!", tx.Hash().Hex())
		}
		return "", 0, 0, err
	}
	if rcpt.Status == ethTypes.ReceiptStatusFailed {
		global.Logger.Errorf("Transaction %s failed", tx.Hash().Hex())
		return "", 0, 0, fmt.Errorf("transaction failed")
	}

	// Get return value
	noteId, err := ExtractNoteID(rcpt)
	if err != nil {
		global.Logger.Errorf("Failed to extract note id for transaction %s with error: %v", tx.Hash().Hex(), err)
		return "", 0, 0, err
	}

	return tx.Hash().Hex(), characterId, noteId, nil
}
