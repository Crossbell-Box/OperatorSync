package chain

import (
	"context"
	"errors"
	crossbellContract "github.com/Crossbell-Box/OperatorSync/app/worker/chain/contract"
	"github.com/Crossbell-Box/OperatorSync/app/worker/consts"
	"github.com/Crossbell-Box/OperatorSync/app/worker/global"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
)

func PostNoteForCharacter(characterIdStr string, metadataUri string) (string, error) {
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
	tx, err = contractInstance.PostNote(
		operatorAuth,
		crossbellContract.DataTypesPostNoteData{
			CharacterId: big.NewInt(characterId),
			ContentUri:  metadataUri,
		},
	)
	if err != nil {
		global.Logger.Errorf("Failed to create transaction: %s", err.Error())
		return "", err
	}

	// Wait till transaction finish or timeout
	ctx, cancel := context.WithTimeout(context.Background(), consts.CONFIG_TRANSACTION_TIMEOUT)
	defer cancel()
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			global.Logger.Errorf("Transaction %s timed out!", tx.Hash().Hex())
		} else {
			global.Logger.Errorf("Transaction %s failed with error: %s", tx.Hash().Hex(), err.Error())
		}
		return "", err
	}

	return tx.Hash().Hex(), nil
}
