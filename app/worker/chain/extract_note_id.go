package chain

import (
	"bytes"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	POST_NOTE_TOPIC_0_HEX = "0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740" // keccak 'PostNote(uint256,uint256,bytes32,bytes32,bytes)'
)

var (
	postNoteTopic0HashBytes []byte
)

func init() {
	postNoteTopic0HashBytes = common.HexToHash(POST_NOTE_TOPIC_0_HEX).Bytes()
}

func ExtractNoteID(rcpt *ethTypes.Receipt) (int64, error) {
	for _, log := range rcpt.Logs {
		if bytes.Compare(log.Topics[0].Bytes(), postNoteTopic0HashBytes) == 0 {
			// Get the 3rd topic
			return log.Topics[2].Big().Int64(), nil
		}
	}

	// No matching log found
	return 0, errors.New("no matching log")

}
