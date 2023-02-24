// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataTypesCharacter is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCharacter struct {
	CharacterId *big.Int
	Handle      string
	Uri         string
	NoteCount   *big.Int
	SocialToken common.Address
	LinkModule  common.Address
}

// DataTypesCreateCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateCharacterData struct {
	To                 common.Address
	Handle             string
	Uri                string
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypesERC721Struct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesERC721Struct struct {
	TokenAddress  common.Address
	Erc721TokenId *big.Int
}

// DataTypesMintNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMintNoteData struct {
	CharacterId    *big.Int
	NoteId         *big.Int
	To             common.Address
	MintModuleData []byte
}

// DataTypesNote is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNote struct {
	LinkItemType [32]byte
	LinkKey      [32]byte
	ContentUri   string
	LinkModule   common.Address
	MintModule   common.Address
	MintNFT      common.Address
	Deleted      bool
	Locked       bool
}

// DataTypesNoteStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNoteStruct struct {
	CharacterId *big.Int
	NoteId      *big.Int
}

// DataTypesPostNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPostNoteData struct {
	CharacterId        *big.Int
	ContentUri         string
	LinkModule         common.Address
	LinkModuleInitData []byte
	MintModule         common.Address
	MintModuleInitData []byte
	Locked             bool
}

// DataTypescreateThenLinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypescreateThenLinkCharacterData struct {
	FromCharacterId *big.Int
	To              common.Address
	LinkType        [32]byte
}

// DataTypeslinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAddressData struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAnyUriData struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkCharacterData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkERC721Data struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	TokenId         *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkLinklistData struct {
	FromCharacterId *big.Int
	ToLinkListId    *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkNoteData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypessetLinkModule4AddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4AddressData struct {
	Account            common.Address
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4LinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4LinklistData struct {
	LinklistId         *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetMintModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetMintModule4NoteData struct {
	CharacterId        *big.Int
	NoteId             *big.Int
	MintModule         common.Address
	MintModuleInitData []byte
}

// DataTypesunlinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAddressData struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
}

// DataTypesunlinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAnyUriData struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
}

// DataTypesunlinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkCharacterData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkERC721Data struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	TokenId         *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkLinklistData struct {
	FromCharacterId *big.Int
	ToLinkListId    *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkNoteData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"ErrCharacterNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrHandleContainsInvalidCharacters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrHandleExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrHandleLengthInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotAddressOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotCharacterOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotEnoughPermission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotEnoughPermissionForThisNote\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoteIsDeleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoteLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoteNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrREC721NotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrSocialTokenExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrTargetAlreadyHasPrimaryCharacter\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OPERATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XSYNC_OPERATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createCharacter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.createThenLinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createThenLinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"deleteNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacter\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getCharacterByHandle\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacterUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLinkModule4Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4ERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4Linklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinklistContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinklistId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linkListId\",\"type\":\"uint256\"}],\"name\":\"getLinklistType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinklistUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getNote\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.Note\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorPermissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getOperators4Note\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"blocklist\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPrimaryCharacterId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissionBitMap\",\"type\":\"uint256\"}],\"name\":\"grantOperatorPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blocklist\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"grantOperators4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linklist_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFTImpl_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"periphery_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newbieVilla_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperatorAllowedForNote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"isPrimaryCharacter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"lockNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"characterIds\",\"type\":\"uint256[]\"}],\"name\":\"migrateOperatorSyncPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MintNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mintNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"postNote4Address\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"postNote4AnyUri\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"}],\"name\":\"postNote4Character\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ERC721Struct\",\"name\":\"erc721\",\"type\":\"tuple\"}],\"name\":\"postNote4ERC721\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"}],\"name\":\"postNote4Linklist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NoteStruct\",\"name\":\"note\",\"type\":\"tuple\"}],\"name\":\"postNote4Note\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setCharacterUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"setHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4AddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4LinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Linklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setLinklistUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setMintModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setMintModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setNoteUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryCharacterId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setSocialToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615d1c80620000216000396000f3fe608060405234801561001057600080fd5b50600436106104545760003560e01c806374f345cf11610241578063c053f6b81161013b578063dc17b6de116100c3578063ef0828ab11610087578063ef0828ab14610ab0578063f2ad807514610ac3578063f316bacd14610ad6578063f6479d7714610ae9578063fe9299fb14610afc57600080fd5b8063dc17b6de14610a28578063dca2713514610a3b578063e56f2fe414610a4e578063e985e9c514610a61578063ec81d19414610a9d57600080fd5b8063cd69fe611161010a578063cd69fe61146109a4578063d23b320b146109b7578063d70e10c6146109ca578063dabb0531146109f5578063db491e8014610a0857600080fd5b8063c053f6b81461095a578063c2a6fe3b1461096b578063c87b56dd1461097e578063cb8e757e1461099157600080fd5b8063983d2737116101c9578063a6e6178d1161018d578063a6e6178d146108fb578063a7ccb4bf1461090e578063af90b11214610921578063b88d4fde14610934578063b9d328451461094757600080fd5b8063983d2737146108875780639864c307146108a25780639a4dec18146108b55780639a50248d146108c8578063a22cb465146108e857600080fd5b806392f7070b1161021057806392f7070b1461081057806393f057e514610823578063952be0ef1461083657806395d89b411461086c57806395d9fa7d1461087457600080fd5b806374f345cf146107c4578063867884e6146107d75780638734bbfc146107ea5780638b4ca06a146107fd57600080fd5b806331b9d08c1161035257806347f94de7116102da5780635fb881831161029e5780635fb8818314610758578063628b644a1461076b5780636352211e1461077e5780636bf55d5f1461079157806370a08231146107b157600080fd5b806347f94de7146106eb57806349186953146106fe5780634f6ccce71461071f5780635932d53d146107325780635a936d101461074557600080fd5b80633a542db7116103215780633a542db71461068457806340ad34d81461069f57806342842e0e146106b257806342966c68146106c557806344b82a24146106d857600080fd5b806331b9d08c1461061f578063327b2a031461064b57806333f06ee61461065e578063388f50831461067157600080fd5b8063144a3e83116103e057806323b872dd116103a457806323b872dd146105aa57806328fbb805146105bd57806329c301c2146105d05780632abc6bf6146105e35780632f745c591461060c57600080fd5b8063144a3e831461053357806318160ddd14610546578063188b04b31461054e578063206657f2146105615780632209d1451461057457600080fd5b806308cb68ff1161042757806308cb68ff146104d4578063095ea7b3146104e95780630c4dd5f2146104fc5780630ff982441461050f5780631316529d1461052257600080fd5b806301ffc9a71461045957806304f3bcec1461048157806306fdde03146104ac578063081812fc146104c1575b600080fd5b61046c610467366004614551565b610b25565b60405190151581526020015b60405180910390f35b601754610494906001600160a01b031681565b6040516001600160a01b039091168152602001610478565b6104b4610b50565b60405161047891906145be565b6104946104cf3660046145d1565b610be2565b6104e76104e2366004614602565b610c7c565b005b6104e76104f736600461464b565b610d53565b6104e761050a366004614602565b610e68565b6104e761051d366004614677565b610f72565b60045b604051908152602001610478565b6104b46105413660046145d1565b611004565b600854610525565b6104e761055c3660046146a5565b61100f565b6104e761056f3660046146d9565b6110ad565b61049461058236600461464b565b6001600160a01b03918216600090815260106020908152604080832093835292905220541690565b6104e76105b8366004614711565b611145565b61046c6105cb366004614741565b611176565b6105256105de36600461478c565b61118d565b6105256105f13660046147c0565b6001600160a01b03166000908152600c602052604090205490565b61052561061a36600461464b565b61123b565b61049461062d3660046147c0565b6001600160a01b039081166000908152601160205260409020541690565b6105256106593660046147ef565b6112d1565b6104e761066c366004614875565b611491565b6104e761067f3660046146a5565b611578565b610494730f588318a494e4508a121a32b6670b5494ca335781565b6104e76106ad3660046148c0565b6115cd565b6104e76106c0366004614711565b611649565b6104e76106d33660046145d1565b611664565b6105256106e63660046148dc565b611701565b6104e76106f9366004614875565b611779565b61071161070c366004614920565b6117e0565b604051610478929190614986565b61052561072d3660046145d1565b611839565b6104e76107403660046149ef565b6118cc565b6104e7610753366004614677565b6119f6565b6104e76107663660046146a5565b611a88565b6104e7610779366004614a30565b611add565b61049461078c3660046145d1565b611b6a565b6107a461079f3660046145d1565b611be1565b6040516104789190614a82565b6105256107bf3660046147c0565b611bfb565b6104e76107d2366004614920565b611c82565b6104e76107e53660046148c0565b611cfc565b61046c6107f83660046145d1565b611d6e565b61052561080b3660046145d1565b611d9c565b61052561081e366004614a95565b611e08565b6104e7610831366004614677565b611e99565b610525610844366004614ae6565b60009182526019602090815260408084206001600160a01b0393909316845291905290205490565b6104b4611f0a565b6104e7610882366004614ae6565b611f19565b61049473da2423cea4f1047556e7a142f81a7ed50e93e16081565b6104e76108b03660046146a5565b611fab565b6105256108c33660046147ef565b612000565b6108db6108d6366004614b0b565b61210a565b6040516104789190614b40565b6104e76108f6366004614bcc565b6122b9565b6104e7610909366004614875565b6122c4565b61052561091c3660046146a5565b61233a565b61052561092f3660046148dc565b6123de565b6104e7610942366004614c65565b61241e565b6104e7610955366004614d25565b612456565b6013546001600160a01b0316610494565b6104e7610979366004614920565b6124c0565b6104b461098c3660046145d1565b612532565b6104e761099f366004614d25565b6125d7565b6105256109b2366004614d25565b612649565b6104e76109c53660046146a5565b612787565b6105256109d8366004614920565b6000918252600d6020908152604080842092845291905290205490565b6108db610a033660046145d1565b612812565b610a1b610a16366004614920565b612999565b6040516104789190614d59565b6104e7610a36366004614def565b612afa565b6104b4610a493660046145d1565b612b8b565b6104e7610a5c366004614e71565b612bfd565b61046c610a6f366004614f2b565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6104b4610aab3660046145d1565b612d8a565b6104e7610abe366004614602565b612daa565b6104e7610ad13660046145d1565b612e1b565b610525610ae4366004614f59565b612e6a565b6104e7610af7366004614677565b612fc8565b610494610b0a3660046145d1565b6000908152600f60205260409020546001600160a01b031690565b60006001600160e01b0319821663780e9d6360e01b1480610b4a5750610b4a82613011565b92915050565b606060008054610b5f90614fb4565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8b90614fb4565b8015610bd85780601f10610bad57610100808354040283529160200191610bd8565b820191906000526020600020905b815481529060010190602001808311610bbb57829003601f168201915b5050505050905090565b6000818152600260205260408120546001600160a01b0316610c605760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b610c8960208201826147c0565b6001600160a01b0316336001600160a01b031614610cba5760405163ca67421960e01b815260040160405180910390fd5b73__$69db18565ecbefece480e92cc5f8fb1274$__63dfc34f25610ce160208401846147c0565b610cf160408501602086016147c0565b610cfe6040860186614fe8565b60116040518663ffffffff1660e01b8152600401610d20959493929190615057565b60006040518083038186803b158015610d3857600080fd5b505af4158015610d4c573d6000803e3d6000fd5b5050505050565b6000610d5e82611b6a565b9050806001600160a01b0316836001600160a01b031603610dcb5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610c57565b336001600160a01b0382161480610de75750610de78133610a6f565b610e595760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610c57565b610e638383613061565b505050565b6013546040516367880d6160e11b8152823560048201526000916001600160a01b03169063cf101ac290602401602060405180830381865afa158015610eb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed69190615096565b9050610ee38160c16130cf565b73__$69db18565ecbefece480e92cc5f8fb1274$__636252159e8335610f0f60408601602087016147c0565b610f1c6040870187614fe8565b600f6040518663ffffffff1660e01b8152600401610f3e9594939291906150af565b60006040518083038186803b158015610f5657600080fd5b505af4158015610f6a573d6000803e3d6000fd5b505050505050565b610f7e813560b26130cf565b60135481356000818152600d6020908152604080832081870135808552908352928190205490516328066f9160e21b8152600481019490945290850135602484015260448301919091526001600160a01b039092166064820152608481019190915273__$e2c9a8f399964af47bb173600dd9e5f662$__9063a019be449060a401610d20565b6060610b4a82612532565b61101b813560b26130cf565b6110288160200135613160565b73__$e2c9a8f399964af47bb173600dd9e5f662$__639ec52a2382356020840135604085013561105b6060870187614fe8565b6013546020808a01356000908152600a9091526040908190206005015490516001600160e01b031960e08a901b168152610d209796959493926001600160a01b03908116921690600d906004016150da565b6110b88360026130cf565b604051631f8c0b6760e11b8152600481018490526001600160a01b038316602482015260448101829052601860648201526019608482015273__$bfdcc7011a2e80c167e18378382b8de19c$__90633f1816ce9060a4015b60006040518083038186803b15801561112857600080fd5b505af415801561113c573d6000803e3d6000fd5b50505050505050565b61114f3382613198565b61116b5760405162461bcd60e51b8152600401610c5790615127565b610e6383838361328f565b6000611183848484613436565b90505b9392505050565b600061119b823560ec6130cf565b81356000908152600a60205260408120600301805482906111bb9061518e565b91829055506040516342a34a5360e01b815290915073__$2873850e5b3f5697aafc2f5cd9f27cc884$__906342a34a539061120490869085906000908190600e906004016152b1565b60006040518083038186803b15801561121c57600080fd5b505af4158015611230573d6000803e3d6000fd5b509295945050505050565b600061124683611bfb565b82106112a85760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610c57565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b60006112df833560ca6130cf565b6112f96112ef60208401846147c0565b83602001356134c0565b6013546545524337323160d01b906000906001600160a01b0316632ea24efc8261132660208801886147c0565b6040516001600160e01b031960e085901b16815260048101929092526001600160a01b03166024820152602087013560448201526064016020604051808303816000875af115801561137c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a09190615096565b85356000908152600a602052604081206003018054929350909182906113c59061518e565b9182905550905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386866113f760208c018c6147c0565b8b6020013560405160200161142592919060609290921b6001600160601b0319168252601482015260340190565b604051602081830303815290604052600e6040518763ffffffff1660e01b8152600401611457969594939291906152f6565b60006040518083038186803b15801561146f57600080fd5b505af4158015611483573d6000803e3d6000fd5b509298975050505050505050565b6013546040516367880d6160e11b8152600481018590526000916001600160a01b03169063cf101ac290602401602060405180830381865afa1580156114db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ff9190615096565b905061150c8160b16130cf565b601354604051633c17845760e11b81526001600160a01b039091169063782f08ae9061154090879087908790600401615341565b600060405180830381600087803b15801561155a57600080fd5b505af115801561156e573d6000803e3d6000fd5b5050505050505050565b611584813560b96130cf565b60135460405163e4bb4f2360e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9163e4bb4f2391610d209185916001600160a01b0390911690600d9060040161535b565b6115d9813560b66130cf565b60135460408051635cebf08d60e11b81528335600482015260208401356024820152908301356044820152606083013560648201526001600160a01b039091166084820152600d60a482015273__$e2c9a8f399964af47bb173600dd9e5f662$__9063b9d7e11a9060c401610d20565b610e638383836040518060200160405280600081525061241e565b6000818152600a6020526040808220905161168291600101906153c9565b60408051918290039091206000818152600b6020908152838220829055858252600a9052918220828155909250906116bd60018301826144a0565b6116cb6002830160006144a0565b50600060038201556004810180546001600160a01b03199081169091556005909101805490911690556116fd82613558565b5050565b600061170f833560c86130cf565b82356000908152600a60205260408120600301805467131a5b9adb1a5cdd60c21b928592909182906117409061518e565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a60405160200161142591815260200190565b6117848360b06130cf565b6000838152600a602052604090206002016117a0828483615485565b50827f17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c663199783836040516117d3929190615544565b60405180910390a2505050565b6000828152601a6020908152604080832084845290915290206060908190611807906135b7565b6000858152601a60209081526040808320878452909152902090925061182f906002016135b7565b90505b9250929050565b600061184460085490565b82106118a75760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610c57565b600882815481106118ba576118ba615558565b90600052602060002001549050919050565b3373da2423cea4f1047556e7a142f81a7ed50e93e1601461191f5760405162461bcd60e51b815260206004820152600d60248201526c37b7363c9037b832b930ba37b960991b6044820152606401610c57565b60005b81811015610e635773__$bfdcc7011a2e80c167e18378382b8de19c$__633f1816ce84848481811061195657611956615558565b60405160e085901b6001600160e01b031916815260209190910290920135600483015250730f588318a494e4508a121a32b6670b5494ca3357602482015264400000003f60c61b6044820152601860648201526019608482015260a40160006040518083038186803b1580156119cb57600080fd5b505af41580156119df573d6000803e3d6000fd5b5050505080806119ee9061518e565b915050611922565b611a02813560be6130cf565b60135481356000818152600d60209081526040808320818701358085529083529281902054905163d862986560e01b8152600481019490945290850135602484015260448301919091526001600160a01b039092166064820152608481019190915273__$e2c9a8f399964af47bb173600dd9e5f662$__9063d86298659060a401610d20565b611a94813560bb6130cf565b6013546040516332f107bf60e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__916332f107bf91610d209185916001600160a01b0390911690600d9060040161556e565b611ae784846135c4565b611af18484613629565b611afb84846136a0565b6040516001626802bf60e01b0319815273__$2873850e5b3f5697aafc2f5cd9f27cc884$__9063ff97fd4190611b3e908790879087908790600e906004016155eb565b60006040518083038186803b158015611b5657600080fd5b505af415801561156e573d6000803e3d6000fd5b6000818152600260205260408120546001600160a01b031680610b4a5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610c57565b6000818152601860205260409020606090610b4a906135b7565b60006001600160a01b038216611c665760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610c57565b506001600160a01b031660009081526003602052604090205490565b611c8d8260c46130cf565b611c978282613629565b6000828152600e60209081526040808320848452825291829020600501805460ff60a81b1916600160a81b179055905182815283917f036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd3091015b60405180910390a25050565b611d08813560b86130cf565b60135481356000908152600d6020908152604080832060608601358452909152908190205490516317e9a8b160e31b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9263bf4d458892610d209286926001600160a01b0316919060040161560b565b600080611d7a83611b6a565b6001600160a01b03166000908152600c60205260409020549290921492915050565b60135460405162fba02760e01b8152600481018390526000916001600160a01b03169062fba02790602401602060405180830381865afa158015611de4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b4a9190615096565b6000611e16833560c76130cf565b82356000908152600a602052604081206003018054664164647265737360c81b926001600160a01b0386169290918290611e4f9061518e565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a604051602001611425919060609190911b6001600160601b031916815260140190565b611ea5813560ba6130cf565b60135481356000908152600d6020908152604080832081860135845290915290819020549051633d7f9b3d60e11b815273__$e2c9a8f399964af47bb173600dd9e5f662$__92637aff367a92610d209286926001600160a01b0316919060040161565d565b606060018054610b5f90614fb4565b611f248260016130cf565b6000828152600a60205260409020600401546001600160a01b031615611f5d5760405163fe6f50e560e01b815260040160405180910390fd5b6040516384b44a2f60e01b8152600481018390526001600160a01b0382166024820152600a604482015273__$d8529f0216812a2dec6b96adf1943c3537$__906384b44a2f90606401610f3e565b611fb7813560bd6130cf565b601354604051635fe5df1d60e11b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9163bfcbbe3a91610d209185916001600160a01b0390911690600d906004016156a5565b600061200e833560c96130cf565b601354604051635cb46be760e01b81526000600482018190528435602483015260208501356044830152634e6f746560e01b9290916001600160a01b0390911690635cb46be7906064016020604051808303816000875af1158015612077573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061209b9190615096565b85356000908152600a602052604081206003018054929350909182906120c09061518e565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a600001358b60200135604051602001611425929190918252602082015260400190565b6121126144da565b60008383604051612124929190615705565b604080519182900382206000818152600b602090815283822054808352600a82529184902060c086019094528354855260018401805493965091949392908401919061216f90614fb4565b80601f016020809104026020016040519081016040528092919081815260200182805461219b90614fb4565b80156121e85780601f106121bd576101008083540402835291602001916121e8565b820191906000526020600020905b8154815290600101906020018083116121cb57829003601f168201915b5050505050815260200160028201805461220190614fb4565b80601f016020809104026020016040519081016040528092919081815260200182805461222d90614fb4565b801561227a5780601f1061224f5761010080835404028352916020019161227a565b820191906000526020600020905b81548152906001019060200180831161225d57829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015295945050505050565b6116fd3383836136e5565b6122cf8360006130cf565b6122ef82826040516122e2929190615705565b60405180910390206137b3565b6122f982826137e0565b60405163130f361d60e01b815273__$d8529f0216812a2dec6b96adf1943c3537$__9063130f361d9061111090869086908690600b90600a90600401615715565b600061234b82356020840135613629565b73__$2873850e5b3f5697aafc2f5cd9f27cc884$__635be694158335602085013561237c60608701604088016147c0565b6123896060880188614fe8565b6014546040516001600160e01b031960e089901b1681526123c19695949392916001600160a01b031690600a90600e90600401615743565b602060405180830381865af4158015611de4573d6000803e3d6000fd5b60006123ec833560c66130cf565b82356000908152600a6020526040812060030180546821b430b930b1ba32b960b91b928592909182906117409061518e565b6124283383613198565b6124445760405162461bcd60e51b8152600401610c5790615127565b61245084848484613846565b50505050565b612462813560b56130cf565b61247481602001358260400135613629565b601354604051631619823f60e31b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9163b0cc11f891610d209185916001600160a01b0390911690600e90600d90600401615791565b6124cb8260c56130cf565b6124d58282613629565b6000828152600e60209081526040808320848452825291829020600501805460ff60a01b1916600160a01b179055905182815283917f4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf939101611cf0565b6000818152600a6020526040902060020180546060919061255290614fb4565b80601f016020809104026020016040519081016040528092919081815260200182805461257e90614fb4565b80156125cb5780601f106125a0576101008083540402835291602001916125cb565b820191906000526020600020905b8154815290600101906020018083116125ae57829003601f168201915b50505050509050919050565b6125e3813560b76130cf565b6126006125f660408301602084016147c0565b82604001356134c0565b601354604051638f3334ff60e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__91638f3334ff91610d209185916001600160a01b0390911690600d90600401615807565b600061266961265b6020840184614fe8565b6040516122e2929190615705565b61267e6126796020840184614fe8565b6137e0565b60126000815461268d9061518e565b918290555090506126aa6126a460208401846147c0565b82613879565b60405163470cc98960e11b815273__$d8529f0216812a2dec6b96adf1943c3537$__90638e199312906126e99085908590600b90600a9060040161586a565b60006040518083038186803b15801561270157600080fd5b505af4158015612715573d6000803e3d6000fd5b50600c92506000915061272d905060208501856147c0565b6001600160a01b03166001600160a01b03168152602001908152602001600020546000036127825780600c600061276760208601866147c0565b6001600160a01b031681526020810191909152604001600020555b919050565b612793813560c26130cf565b6127a281356020830135613629565b6127b1813560208301356136a0565b73__$69db18565ecbefece480e92cc5f8fb1274$__6320828a02823560208401356127e260608601604087016147c0565b6127ef6060870187614fe8565b600e6040518763ffffffff1660e01b8152600401610d2096959493929190615939565b61281a6144da565b600a60008381526020019081526020016000206040518060c00160405290816000820154815260200160018201805461285290614fb4565b80601f016020809104026020016040519081016040528092919081815260200182805461287e90614fb4565b80156128cb5780601f106128a0576101008083540402835291602001916128cb565b820191906000526020600020905b8154815290600101906020018083116128ae57829003601f168201915b505050505081526020016002820180546128e490614fb4565b80601f016020809104026020016040519081016040528092919081815260200182805461291090614fb4565b801561295d5780601f106129325761010080835404028352916020019161295d565b820191906000526020600020905b81548152906001019060200180831161294057829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015292915050565b60408051610100808201835260008083526020808401829052606084860181905284018290526080840182905260a0840182905260c0840182905260e08401829052868252600e81528482208683528152908490208451928301855280548352600181015491830191909152600281018054939492939192840191612a1d90614fb4565b80601f0160208091040260200160405190810160405280929190818152602001828054612a4990614fb4565b8015612a965780601f10612a6b57610100808354040283529160200191612a96565b820191906000526020600020905b815481529060010190602001808311612a7957829003601f168201915b505050918352505060038201546001600160a01b039081166020830152600483015481166040830152600590920154918216606082015260ff600160a01b8304811615156080830152600160a81b909204909116151560a090910152905092915050565b612b058660036130cf565b612b0f8686613629565b604051630afb883f60e41b815273__$bfdcc7011a2e80c167e18378382b8de19c$__9063afb883f090612b5390899089908990899089908990601a906004016159b8565b60006040518083038186803b158015612b6b57600080fd5b505af4158015612b7f573d6000803e3d6000fd5b50505050505050505050565b601354604051632b05429560e21b8152600481018390526060916001600160a01b03169063ac150a5490602401600060405180830381865afa158015612bd5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b4a9190810190615a00565b601454600290600160a81b900460ff16158015612c28575060145460ff808316600160a01b90920416105b612c8b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610c57565b6014805460ff60a81b1960ff8416600160a01b021661ffff60a01b1990911617600160a81b179055612cbf89898989613893565b601380546001600160a01b03199081166001600160a01b0388811691909117909255601480548216878416179055601580548216868416179055601b80549091169184169190911790556040514281527f400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf9060200160405180910390a16014805460ff60a81b1916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050565b6000818152600a6020526040902060010180546060919061255290614fb4565b612db6813560bc6130cf565b60135481356000908152600d602090815260408083208186013584529091529081902054905163bf5c00c160e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9263bf5c00c192610d209286926001600160a01b03169190600401615a6d565b612e24816138e4565b336000818152600c602052604080822080549085905590519092839285927fce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de9190a45050565b6000612e78843560cb6130cf565b601354604051633610bf0960e11b815265416e7955726960d01b916000916001600160a01b0390911690636c217e1290612eba90849089908990600401615341565b6020604051808303816000875af1158015612ed9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612efd9190615096565b86356000908152600a60205260408120600301805492935090918290612f229061518e565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53888386868b8b604051602001612f5b929190615705565b604051602081830303815290604052600e6040518763ffffffff1660e01b8152600401612f8d969594939291906152f6565b60006040518083038186803b158015612fa557600080fd5b505af4158015612fb9573d6000803e3d6000fd5b50929998505050505050505050565b612fd4813560b46130cf565b61300e8135612fe960408401602085016147c0565b836040013560405180604001604052806002815260200161060f60f31b81525061394a565b50565b60006001600160e01b031982166380ac58cd60e01b148061304257506001600160e01b03198216635b5e139f60e01b145b80610b4a57506301ffc9a760e01b6001600160e01b0319831614610b4a565b600081815260046020526040902080546001600160a01b0319166001600160a01b038416908117909155819061309682611b6a565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6130d882613b6e565b156130e1575050565b6015546001600160a01b0316330361311d576000828152601960209081526040808320328452909152902054600190821c81160361311d575050565b6000828152601960209081526040808320338452909152902054600190821c811603613147575050565b604051632c4bc2b960e21b815260040160405180910390fd5b6000818152600260205260409020546001600160a01b031661300e576040516375af0fc960e11b815260048101829052602401610c57565b6000818152600260205260408120546001600160a01b03166132115760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610c57565b600061321c83611b6a565b9050806001600160a01b0316846001600160a01b031614806132575750836001600160a01b031661324c84610be2565b6001600160a01b0316145b8061328757506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b949350505050565b826001600160a01b03166132a282611b6a565b6001600160a01b0316146133065760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610c57565b6001600160a01b0382166133685760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610c57565b613373838383613bce565b61337e600082613061565b6001600160a01b03831660009081526003602052604081208054600192906133a7908490615ac5565b90915550506001600160a01b03821660009081526003602052604081208054600192906133d5908490615ad8565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6000838152601a6020908152604080832085845290915281206134598184613c9e565b15613468576000915050611186565b6134756002820184613c9e565b15613484576001915050611186565b60008581526019602090815260408083206001600160a01b03871684529091529020546134b79060c31c60019081161490565b95945050505050565b6040516331a9108f60e11b8152600481018290526000906001600160a01b03841690636352211e90602401602060405180830381865afa158015613508573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061352c9190615aeb565b90506001600160a01b038116600003610e6357604051631b96896160e21b815260040160405180910390fd5b6135623382613198565b6135ae5760405162461bcd60e51b815260206004820152601b60248201527f4e4654426173653a204e6f744f776e65724f72417070726f76656400000000006044820152606401610c57565b61300e81613cc0565b6060600061118683613d67565b6135cd82613b6e565b156135d6575050565b6015546001600160a01b031633036135fc576135f3828232613436565b156135fc575050565b613607828233613436565b15613610575050565b604051631a1d1d4760e11b815260040160405180910390fd5b6000828152600e60209081526040808320848452909152902060050154600160a01b900460ff161561366e57604051631f0fc8f560e11b815260040160405180910390fd5b6000828152600a60205260409020600301548111156116fd576040516364783acb60e01b815260040160405180910390fd5b6000828152600e60209081526040808320848452909152902060050154600160a81b900460ff16156116fd57604051630bc06a0f60e21b815260040160405180910390fd5b816001600160a01b0316836001600160a01b0316036137465760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610c57565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6000818152600b60205260409020541561300e57604051631b659b9f60e21b815260040160405180910390fd5b8181601f8111806137f15750600381105b1561380f57604051636f819c2160e11b815260040160405180910390fd5b60005b81811015610d4c5761383e83838381811061382f5761382f615558565b9050013560f81c60f81b613dc2565b600101613812565b61385184848461328f565b61385d84848484613e70565b6124505760405162461bcd60e51b8152600401610c5790615b08565b6116fd828260405180602001604052806000815250613f71565b61389f84848484613fa4565b7f414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c3084848484426040516138d6959493929190615b5a565b60405180910390a150505050565b60006138ef82611b6a565b6015549091506001600160a01b0316331480156139145750326001600160a01b038216145b1561391d575050565b6001600160a01b0381163303613931575050565b604051631b0c476f60e11b815260040160405180910390fd5b6001600160a01b0383166000908152600c602052604090205415613981576040516317a4570160e01b815260040160405180910390fd5b6040516001600160601b0319606085901b1660208201526139ba90603401604051602081830303815290604052805190602001206137b3565b60006012600081546139cb9061518e565b918290555090506139dc8482613879565b73__$d8529f0216812a2dec6b96adf1943c3537$__638e1993126040518060a00160405280876001600160a01b0316815260200187604051602001613a34919060609190911b6001600160601b031916815260140190565b604051602081830303815290604052815260200160405180602001604052806000815250815260200160006001600160a01b031681526020016040518060200160405280600081525081525083600b600a6040518563ffffffff1660e01b8152600401613aa49493929190615b94565b60006040518083038186803b158015613abc57600080fd5b505af4158015613ad0573d6000803e3d6000fd5b5050506001600160a01b038086166000908152600c60205260408082208590556013549051639ec52a2360e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9450639ec52a2393613b37938b9388938b938b9392169190600d90600401615c2b565b60006040518083038186803b158015613b4f57600080fd5b505af4158015613b63573d6000803e3d6000fd5b505050505050505050565b600080613b7a83611b6a565b90506001600160a01b0381163303613b955750600192915050565b6015546001600160a01b031633148015613bb75750326001600160a01b038216145b15613bc55750600192915050565b50600092915050565b601b546001600160a01b03848116911614613c5b576000818152601860205260408120613bfa90613fbf565b600083815260186020526040812091925090613c15906135b7565b905060005b82811015613c5757613c4584838381518110613c3857613c38615558565b6020026020010151613fc9565b80613c4f8161518e565b915050613c1a565b5050505b6001600160a01b0383166000908152600c602052604090205415613c93576001600160a01b0383166000908152600c60205260408120555b610e63838383614002565b6001600160a01b03811660009081526001830160205260408120541515611186565b6000613ccb82611b6a565b9050613cd981600084613bce565b613ce4600083613061565b6001600160a01b0381166000908152600360205260408120805460019290613d0d908490615ac5565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b6060816000018054806020026020016040519081016040528092919081815260200182805480156125cb57602002820191906000526020600020905b815481526020019060010190808311613da35750505050509050919050565b600360fc1b6001600160f81b031982161080613deb5750603d60f91b6001600160f81b03198216115b80613e1b5750603960f81b6001600160f81b03198216118015613e1b5750606160f81b6001600160f81b03198216105b8015613e355750602d60f81b6001600160f81b0319821614155b8015613e4f5750605f60f81b6001600160f81b0319821614155b1561300e576040516001621693dd60e01b0319815260040160405180910390fd5b60006001600160a01b0384163b15613f6657604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290613eb4903390899088908890600401615c76565b6020604051808303816000875af1925050508015613eef575060408051601f3d908101601f19168201909252613eec91810190615cb3565b60015b613f4c573d808015613f1d576040519150601f19603f3d011682016040523d82523d6000602084013e613f22565b606091505b508051600003613f445760405162461bcd60e51b8152600401610c5790615b08565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050613287565b506001949350505050565b613f7b83836140ba565b613f886000848484613e70565b610e635760405162461bcd60e51b8152600401610c5790615b08565b6000613fb1848683615485565b506001610d4c828483615485565b6000610b4a825490565b60008281526019602090815260408083206001600160a01b0385168452825280832083905584835260189091529020610e639082614208565b6001600160a01b03831661405d5761405881600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b614080565b816001600160a01b0316836001600160a01b03161461408057614080838261421d565b6001600160a01b03821661409757610e63816142ba565b826001600160a01b0316826001600160a01b031614610e6357610e638282614369565b6001600160a01b0382166141105760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610c57565b6000818152600260205260409020546001600160a01b0316156141755760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610c57565b61418160008383613bce565b6001600160a01b03821660009081526003602052604081208054600192906141aa908490615ad8565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000611186836001600160a01b0384166143ad565b6000600161422a84611bfb565b6142349190615ac5565b600083815260076020526040902054909150808214614287576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b6008546000906142cc90600190615ac5565b600083815260096020526040812054600880549394509092849081106142f4576142f4615558565b90600052602060002001549050806008838154811061431557614315615558565b600091825260208083209091019290925582815260099091526040808220849055858252812055600880548061434d5761434d615cd0565b6001900381819060005260206000200160009055905550505050565b600061437483611bfb565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b600081815260018301602052604081205480156144965760006143d1600183615ac5565b85549091506000906143e590600190615ac5565b905081811461444a57600086600001828154811061440557614405615558565b906000526020600020015490508087600001848154811061442857614428615558565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061445b5761445b615cd0565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610b4a565b6000915050610b4a565b5080546144ac90614fb4565b6000825580601f106144bc575050565b601f01602090049060005260206000209081019061300e9190614522565b6040518060c001604052806000815260200160608152602001606081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b5b808211156145375760008155600101614523565b5090565b6001600160e01b03198116811461300e57600080fd5b60006020828403121561456357600080fd5b81356111868161453b565b60005b83811015614589578181015183820152602001614571565b50506000910152565b600081518084526145aa81602086016020860161456e565b601f01601f19169290920160200192915050565b6020815260006111866020830184614592565b6000602082840312156145e357600080fd5b5035919050565b6000606082840312156145fc57600080fd5b50919050565b60006020828403121561461457600080fd5b81356001600160401b0381111561462a57600080fd5b613287848285016145ea565b6001600160a01b038116811461300e57600080fd5b6000806040838503121561465e57600080fd5b823561466981614636565b946020939093013593505050565b60006060828403121561468957600080fd5b61118683836145ea565b6000608082840312156145fc57600080fd5b6000602082840312156146b757600080fd5b81356001600160401b038111156146cd57600080fd5b61328784828501614693565b6000806000606084860312156146ee57600080fd5b83359250602084013561470081614636565b929592945050506040919091013590565b60008060006060848603121561472657600080fd5b833561473181614636565b9250602084013561470081614636565b60008060006060848603121561475657600080fd5b8335925060208401359150604084013561476f81614636565b809150509250925092565b600060e082840312156145fc57600080fd5b60006020828403121561479e57600080fd5b81356001600160401b038111156147b457600080fd5b6132878482850161477a565b6000602082840312156147d257600080fd5b813561118681614636565b6000604082840312156145fc57600080fd5b6000806060838503121561480257600080fd5b82356001600160401b0381111561481857600080fd5b6148248582860161477a565b92505061182f84602085016147dd565b60008083601f84011261484657600080fd5b5081356001600160401b0381111561485d57600080fd5b60208301915083602082850101111561183257600080fd5b60008060006040848603121561488a57600080fd5b8335925060208401356001600160401b038111156148a757600080fd5b6148b386828701614834565b9497909650939450505050565b6000608082840312156148d257600080fd5b6111868383614693565b600080604083850312156148ef57600080fd5b82356001600160401b0381111561490557600080fd5b6149118582860161477a565b95602094909401359450505050565b6000806040838503121561493357600080fd5b50508035926020909101359150565b600081518084526020808501945080840160005b8381101561497b5781516001600160a01b031687529582019590820190600101614956565b509495945050505050565b6040815260006149996040830185614942565b82810360208401526134b78185614942565b60008083601f8401126149bd57600080fd5b5081356001600160401b038111156149d457600080fd5b6020830191508360208260051b850101111561183257600080fd5b60008060208385031215614a0257600080fd5b82356001600160401b03811115614a1857600080fd5b614a24858286016149ab565b90969095509350505050565b60008060008060608587031215614a4657600080fd5b843593506020850135925060408501356001600160401b03811115614a6a57600080fd5b614a7687828801614834565b95989497509550505050565b6020815260006111866020830184614942565b60008060408385031215614aa857600080fd5b82356001600160401b03811115614abe57600080fd5b614aca8582860161477a565b9250506020830135614adb81614636565b809150509250929050565b60008060408385031215614af957600080fd5b823591506020830135614adb81614636565b60008060208385031215614b1e57600080fd5b82356001600160401b03811115614b3457600080fd5b614a2485828601614834565b60208152815160208201526000602083015160c06040840152614b6660e0840182614592565b90506040840151601f19848303016060850152614b838282614592565b91505060608401516080840152608084015160018060a01b0380821660a08601528060a08701511660c086015250508091505092915050565b8035801515811461278257600080fd5b60008060408385031215614bdf57600080fd5b8235614bea81614636565b915061182f60208401614bbc565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715614c3657614c36614bf8565b604052919050565b60006001600160401b03821115614c5757614c57614bf8565b50601f01601f191660200190565b60008060008060808587031215614c7b57600080fd5b8435614c8681614636565b93506020850135614c9681614636565b92506040850135915060608501356001600160401b03811115614cb857600080fd5b8501601f81018713614cc957600080fd5b8035614cdc614cd782614c3e565b614c0e565b818152886020838501011115614cf157600080fd5b8160208401602083013760006020838301015280935050505092959194509250565b600060a082840312156145fc57600080fd5b600060208284031215614d3757600080fd5b81356001600160401b03811115614d4d57600080fd5b61328784828501614d13565b60208152815160208201526020820151604082015260006040830151610100806060850152614d8c610120850183614592565b9150606085015160018060a01b0380821660808701528060808801511660a0870152505060a0850151614dca60c08601826001600160a01b03169052565b5060c085015180151560e08601525060e0850151801515858301525090949350505050565b60008060008060008060808789031215614e0857600080fd5b863595506020870135945060408701356001600160401b0380821115614e2d57600080fd5b614e398a838b016149ab565b90965094506060890135915080821115614e5257600080fd5b50614e5f89828a016149ab565b979a9699509497509295939492505050565b60008060008060008060008060c0898b031215614e8d57600080fd5b88356001600160401b0380821115614ea457600080fd5b614eb08c838d01614834565b909a50985060208b0135915080821115614ec957600080fd5b50614ed68b828c01614834565b9097509550506040890135614eea81614636565b93506060890135614efa81614636565b92506080890135614f0a81614636565b915060a0890135614f1a81614636565b809150509295985092959890939650565b60008060408385031215614f3e57600080fd5b8235614f4981614636565b91506020830135614adb81614636565b600080600060408486031215614f6e57600080fd5b83356001600160401b0380821115614f8557600080fd5b614f918783880161477a565b94506020860135915080821115614fa757600080fd5b506148b386828701614834565b600181811c90821680614fc857607f821691505b6020821081036145fc57634e487b7160e01b600052602260045260246000fd5b6000808335601e19843603018112614fff57600080fd5b8301803591506001600160401b0382111561501957600080fd5b60200191503681900382131561183257600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03868116825285166020820152608060408201819052600090615084908301858761502e565b90508260608301529695505050505050565b6000602082840312156150a857600080fd5b5051919050565b8581526001600160a01b0385166020820152608060408201819052600090615084908301858761502e565b88815287602082015286604082015260e06060820152600061510060e08301878961502e565b6001600160a01b0395861660808401529390941660a082015260c001529695505050505050565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b6000600182016151a0576151a0615178565b5060010190565b6000808335601e198436030181126151be57600080fd5b83016020810192503590506001600160401b038111156151dd57600080fd5b80360382131561183257600080fd5b8035825260006151ff60208301836151a7565b60e0602086015261521460e08601828461502e565b915050604083013561522581614636565b6001600160a01b03818116604087015261524260608601866151a7565b9250868403606088015261525784848361502e565b9350506080850135915061526a82614636565b16608085015261527d60a08401846151a7565b85830360a087015261529083828461502e565b925050506152a060c08401614bbc565b151560c08501528091505092915050565b60c0815260006152c460c08301886151ec565b602083810197909752604083019590955250606081019290925281830360808301526000835260a09091015201919050565b60c08152600061530960c08301896151ec565b876020840152866040840152856060840152828103608084015261532d8186614592565b9150508260a0830152979650505050505050565b8381526040602082015260006134b760408301848661502e565b60608152833560608201526000602085013561537681614636565b6001600160a01b038181166080850152604087013560a085015261539d60608801886151a7565b9250608060c08601526153b460e08601848361502e565b96909116602085015250505060400152919050565b60008083546153d781614fb4565b600182811680156153ef576001811461540457615433565b60ff1984168752821515830287019450615433565b8760005260208060002060005b8581101561542a5781548a820152908401908201615411565b50505082870194505b50929695505050505050565b601f821115610e6357600081815260208120601f850160051c810160208610156154665750805b601f850160051c820191505b81811015610f6a57828155600101615472565b6001600160401b0383111561549c5761549c614bf8565b6154b0836154aa8354614fb4565b8361543f565b6000601f8411600181146154e457600085156154cc5750838201355b600019600387901b1c1916600186901b178355610d4c565b600083815260209020601f19861690835b8281101561551557868501358255602094850194600190920191016154f5565b50868210156155325760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152600061118360208301848661502e565b634e487b7160e01b600052603260045260246000fd5b6060815283356060820152600061558860208601866151a7565b60808085015261559c60e08501828461502e565b915050604086013560a08401526155b660608701876151a7565b848303605f190160c08601526155cd83828461502e565b6001600160a01b039790971660208601525050505060400152919050565b85815284602082015260806040820152600061508460808301858761502e565b8335815260c08101602085013561562181614636565b6001600160a01b039081166020840152604086810135908401526060958601359583019590955292909316608084015260a09092019190915290565b8335815260a08101602085013561567381614636565b6001600160a01b0390811660208401526040958601359583019590955292909316606084015260809092019190915290565b606081528335606082015260208401356080820152604084013560a082015260006156d360608601866151a7565b608060c08501526156e860e08501828461502e565b6001600160a01b0396909616602085015250505060400152919050565b8183823760009101908152919050565b85815260806020820152600061572f60808301868861502e565b604083019490945250606001529392505050565b888152876020820152600060018060a01b03808916604084015260e0606084015261577260e08401888a61502e565b951660808301525060a081019290925260c09091015295945050505050565b6080815284356080820152602085013560a0820152604085013560c0820152606085013560e082015260006157c960808701876151a7565b60a06101008501526157e06101208501828461502e565b6001600160a01b039790971660208501525050506040810192909252606090910152919050565b60608152833560608201526000602085013561582281614636565b60018060a01b038082166080850152604087013560a0850152606087013560c085015261585260808801886151a7565b925060a060e08601526153b46101008601848361502e565b608081526000853561587b81614636565b6001600160a01b03818116608085015261589860208901896151a7565b925060a0808601526158af6101208601848361502e565b9250506158bf60408901896151a7565b607f19808786030160c08801526158d785838561502e565b945060608b013592506158e983614636565b83831660e08801526158fe60808c018c6151a7565b945092508087860301610100880152505061591a83838361502e565b6020860198909852505050506040810192909252606090910152919050565b86815285602082015260018060a01b038516604082015260a06060820152600061596760a08301858761502e565b9050826080830152979650505050505050565b8183526000602080850194508260005b8581101561497b57813561599d81614636565b6001600160a01b03168752958201959082019060010161598a565b87815286602082015260a0604082015260006159d860a08301878961597a565b82810360608401526159eb81868861597a565b91505082608083015298975050505050505050565b600060208284031215615a1257600080fd5b81516001600160401b03811115615a2857600080fd5b8201601f81018413615a3957600080fd5b8051615a47614cd782614c3e565b818152856020838501011115615a5c57600080fd5b6134b782602083016020860161456e565b60608152833560608201526000615a8760208601866151a7565b60606080850152615a9c60c08501828461502e565b60409788013560a08601526001600160a01b039690961660208501525050509092019190915290565b81810381811115610b4a57610b4a615178565b80820180821115610b4a57610b4a615178565b600060208284031215615afd57600080fd5b815161118681614636565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b606081526000615b6e60608301878961502e565b8281036020840152615b8181868861502e565b9150508260408301529695505050505050565b60808152600060018060a01b03808751166080840152602087015160a080850152615bc3610120850182614592565b90506040880151607f19808684030160c0870152615be18383614592565b92508360608b01511660e087015260808a0151935080868403016101008701525050615c0d8183614592565b60208501979097525050506040810192909252606090910152919050565b87815286602082015285604082015260e060608201526000615c5060e0830187614592565b6001600160a01b0395861660808401529390941660a082015260c0015295945050505050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090615ca990830184614592565b9695505050505050565b600060208284031215615cc557600080fd5b81516111868161453b565b634e487b7160e01b600052603160045260246000fdfea264697066735822122025b55e71b31ddb1e527de765dba420d7f1fc5b2cd4a9189edcb2ee358275190d64736f6c63430008100033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(address)
func (_Contract *ContractCaller) OPERATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "OPERATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(address)
func (_Contract *ContractSession) OPERATOR() (common.Address, error) {
	return _Contract.Contract.OPERATOR(&_Contract.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(address)
func (_Contract *ContractCallerSession) OPERATOR() (common.Address, error) {
	return _Contract.Contract.OPERATOR(&_Contract.CallOpts)
}

// XSYNCOPERATOR is a free data retrieval call binding the contract method 0x3a542db7.
//
// Solidity: function XSYNC_OPERATOR() view returns(address)
func (_Contract *ContractCaller) XSYNCOPERATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "XSYNC_OPERATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XSYNCOPERATOR is a free data retrieval call binding the contract method 0x3a542db7.
//
// Solidity: function XSYNC_OPERATOR() view returns(address)
func (_Contract *ContractSession) XSYNCOPERATOR() (common.Address, error) {
	return _Contract.Contract.XSYNCOPERATOR(&_Contract.CallOpts)
}

// XSYNCOPERATOR is a free data retrieval call binding the contract method 0x3a542db7.
//
// Solidity: function XSYNC_OPERATOR() view returns(address)
func (_Contract *ContractCallerSession) XSYNCOPERATOR() (common.Address, error) {
	return _Contract.Contract.XSYNCOPERATOR(&_Contract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractCaller) GetCharacter(opts *bind.CallOpts, characterId *big.Int) (DataTypesCharacter, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCharacter", characterId)

	if err != nil {
		return *new(DataTypesCharacter), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCharacter)).(*DataTypesCharacter)

	return out0, err

}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractSession) GetCharacter(characterId *big.Int) (DataTypesCharacter, error) {
	return _Contract.Contract.GetCharacter(&_Contract.CallOpts, characterId)
}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractCallerSession) GetCharacter(characterId *big.Int) (DataTypesCharacter, error) {
	return _Contract.Contract.GetCharacter(&_Contract.CallOpts, characterId)
}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractCaller) GetCharacterByHandle(opts *bind.CallOpts, handle string) (DataTypesCharacter, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCharacterByHandle", handle)

	if err != nil {
		return *new(DataTypesCharacter), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCharacter)).(*DataTypesCharacter)

	return out0, err

}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractSession) GetCharacterByHandle(handle string) (DataTypesCharacter, error) {
	return _Contract.Contract.GetCharacterByHandle(&_Contract.CallOpts, handle)
}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contract *ContractCallerSession) GetCharacterByHandle(handle string) (DataTypesCharacter, error) {
	return _Contract.Contract.GetCharacterByHandle(&_Contract.CallOpts, handle)
}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contract *ContractCaller) GetCharacterUri(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCharacterUri", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contract *ContractSession) GetCharacterUri(characterId *big.Int) (string, error) {
	return _Contract.Contract.GetCharacterUri(&_Contract.CallOpts, characterId)
}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contract *ContractCallerSession) GetCharacterUri(characterId *big.Int) (string, error) {
	return _Contract.Contract.GetCharacterUri(&_Contract.CallOpts, characterId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contract *ContractCaller) GetHandle(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getHandle", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contract *ContractSession) GetHandle(characterId *big.Int) (string, error) {
	return _Contract.Contract.GetHandle(&_Contract.CallOpts, characterId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contract *ContractCallerSession) GetHandle(characterId *big.Int) (string, error) {
	return _Contract.Contract.GetHandle(&_Contract.CallOpts, characterId)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contract *ContractCaller) GetLinkModule4Address(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinkModule4Address", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contract *ContractSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4Address(&_Contract.CallOpts, account)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contract *ContractCallerSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4Address(&_Contract.CallOpts, account)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetLinkModule4ERC721(opts *bind.CallOpts, tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinkModule4ERC721", tokenAddress, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4ERC721(&_Contract.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4ERC721(&_Contract.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetLinkModule4Linklist(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinkModule4Linklist", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4Linklist(&_Contract.CallOpts, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetLinkModule4Linklist(&_Contract.CallOpts, tokenId)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contract *ContractCaller) GetLinklistContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinklistContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contract *ContractSession) GetLinklistContract() (common.Address, error) {
	return _Contract.Contract.GetLinklistContract(&_Contract.CallOpts)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contract *ContractCallerSession) GetLinklistContract() (common.Address, error) {
	return _Contract.Contract.GetLinklistContract(&_Contract.CallOpts)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contract *ContractCaller) GetLinklistId(opts *bind.CallOpts, characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinklistId", characterId, linkType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contract *ContractSession) GetLinklistId(characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Contract.Contract.GetLinklistId(&_Contract.CallOpts, characterId, linkType)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contract *ContractCallerSession) GetLinklistId(characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Contract.Contract.GetLinklistId(&_Contract.CallOpts, characterId, linkType)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contract *ContractCaller) GetLinklistType(opts *bind.CallOpts, linkListId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinklistType", linkListId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contract *ContractSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Contract.Contract.GetLinklistType(&_Contract.CallOpts, linkListId)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contract *ContractCallerSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Contract.Contract.GetLinklistType(&_Contract.CallOpts, linkListId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contract *ContractCaller) GetLinklistUri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLinklistUri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contract *ContractSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Contract.Contract.GetLinklistUri(&_Contract.CallOpts, tokenId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contract *ContractCallerSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Contract.Contract.GetLinklistUri(&_Contract.CallOpts, tokenId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contract *ContractCaller) GetNote(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getNote", characterId, noteId)

	if err != nil {
		return *new(DataTypesNote), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNote)).(*DataTypesNote)

	return out0, err

}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contract *ContractSession) GetNote(characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Contract.Contract.GetNote(&_Contract.CallOpts, characterId, noteId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contract *ContractCallerSession) GetNote(characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Contract.Contract.GetNote(&_Contract.CallOpts, characterId, noteId)
}

// GetOperatorPermissions is a free data retrieval call binding the contract method 0x952be0ef.
//
// Solidity: function getOperatorPermissions(uint256 characterId, address operator) view returns(uint256)
func (_Contract *ContractCaller) GetOperatorPermissions(opts *bind.CallOpts, characterId *big.Int, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorPermissions", characterId, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorPermissions is a free data retrieval call binding the contract method 0x952be0ef.
//
// Solidity: function getOperatorPermissions(uint256 characterId, address operator) view returns(uint256)
func (_Contract *ContractSession) GetOperatorPermissions(characterId *big.Int, operator common.Address) (*big.Int, error) {
	return _Contract.Contract.GetOperatorPermissions(&_Contract.CallOpts, characterId, operator)
}

// GetOperatorPermissions is a free data retrieval call binding the contract method 0x952be0ef.
//
// Solidity: function getOperatorPermissions(uint256 characterId, address operator) view returns(uint256)
func (_Contract *ContractCallerSession) GetOperatorPermissions(characterId *big.Int, operator common.Address) (*big.Int, error) {
	return _Contract.Contract.GetOperatorPermissions(&_Contract.CallOpts, characterId, operator)
}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contract *ContractCaller) GetOperators(opts *bind.CallOpts, characterId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperators", characterId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contract *ContractSession) GetOperators(characterId *big.Int) ([]common.Address, error) {
	return _Contract.Contract.GetOperators(&_Contract.CallOpts, characterId)
}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contract *ContractCallerSession) GetOperators(characterId *big.Int) ([]common.Address, error) {
	return _Contract.Contract.GetOperators(&_Contract.CallOpts, characterId)
}

// GetOperators4Note is a free data retrieval call binding the contract method 0x49186953.
//
// Solidity: function getOperators4Note(uint256 characterId, uint256 noteId) view returns(address[] blocklist, address[] allowlist)
func (_Contract *ContractCaller) GetOperators4Note(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int) (struct {
	Blocklist []common.Address
	Allowlist []common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperators4Note", characterId, noteId)

	outstruct := new(struct {
		Blocklist []common.Address
		Allowlist []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Blocklist = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Allowlist = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetOperators4Note is a free data retrieval call binding the contract method 0x49186953.
//
// Solidity: function getOperators4Note(uint256 characterId, uint256 noteId) view returns(address[] blocklist, address[] allowlist)
func (_Contract *ContractSession) GetOperators4Note(characterId *big.Int, noteId *big.Int) (struct {
	Blocklist []common.Address
	Allowlist []common.Address
}, error) {
	return _Contract.Contract.GetOperators4Note(&_Contract.CallOpts, characterId, noteId)
}

// GetOperators4Note is a free data retrieval call binding the contract method 0x49186953.
//
// Solidity: function getOperators4Note(uint256 characterId, uint256 noteId) view returns(address[] blocklist, address[] allowlist)
func (_Contract *ContractCallerSession) GetOperators4Note(characterId *big.Int, noteId *big.Int) (struct {
	Blocklist []common.Address
	Allowlist []common.Address
}, error) {
	return _Contract.Contract.GetOperators4Note(&_Contract.CallOpts, characterId, noteId)
}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contract *ContractCaller) GetPrimaryCharacterId(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPrimaryCharacterId", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contract *ContractSession) GetPrimaryCharacterId(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetPrimaryCharacterId(&_Contract.CallOpts, account)
}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contract *ContractCallerSession) GetPrimaryCharacterId(account common.Address) (*big.Int, error) {
	return _Contract.Contract.GetPrimaryCharacterId(&_Contract.CallOpts, account)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contract *ContractCaller) GetRevision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRevision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contract *ContractSession) GetRevision() (*big.Int, error) {
	return _Contract.Contract.GetRevision(&_Contract.CallOpts)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contract *ContractCallerSession) GetRevision() (*big.Int, error) {
	return _Contract.Contract.GetRevision(&_Contract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// IsOperatorAllowedForNote is a free data retrieval call binding the contract method 0x28fbb805.
//
// Solidity: function isOperatorAllowedForNote(uint256 characterId, uint256 noteId, address operator) view returns(bool)
func (_Contract *ContractCaller) IsOperatorAllowedForNote(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isOperatorAllowedForNote", characterId, noteId, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorAllowedForNote is a free data retrieval call binding the contract method 0x28fbb805.
//
// Solidity: function isOperatorAllowedForNote(uint256 characterId, uint256 noteId, address operator) view returns(bool)
func (_Contract *ContractSession) IsOperatorAllowedForNote(characterId *big.Int, noteId *big.Int, operator common.Address) (bool, error) {
	return _Contract.Contract.IsOperatorAllowedForNote(&_Contract.CallOpts, characterId, noteId, operator)
}

// IsOperatorAllowedForNote is a free data retrieval call binding the contract method 0x28fbb805.
//
// Solidity: function isOperatorAllowedForNote(uint256 characterId, uint256 noteId, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsOperatorAllowedForNote(characterId *big.Int, noteId *big.Int, operator common.Address) (bool, error) {
	return _Contract.Contract.IsOperatorAllowedForNote(&_Contract.CallOpts, characterId, noteId, operator)
}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contract *ContractCaller) IsPrimaryCharacter(opts *bind.CallOpts, characterId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isPrimaryCharacter", characterId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contract *ContractSession) IsPrimaryCharacter(characterId *big.Int) (bool, error) {
	return _Contract.Contract.IsPrimaryCharacter(&_Contract.CallOpts, characterId)
}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contract *ContractCallerSession) IsPrimaryCharacter(characterId *big.Int) (bool, error) {
	return _Contract.Contract.IsPrimaryCharacter(&_Contract.CallOpts, characterId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractSession) Resolver() (common.Address, error) {
	return _Contract.Contract.Resolver(&_Contract.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractCallerSession) Resolver() (common.Address, error) {
	return _Contract.Contract.Resolver(&_Contract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenByIndex(&_Contract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenByIndex(&_Contract.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenOfOwnerByIndex(&_Contract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenOfOwnerByIndex(&_Contract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contract *ContractCaller) TokenURI(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenURI", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contract *ContractSession) TokenURI(characterId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, characterId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contract *ContractCallerSession) TokenURI(characterId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, characterId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Burn(&_Contract.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Burn(&_Contract.TransactOpts, tokenId)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns(uint256 characterId)
func (_Contract *ContractTransactor) CreateCharacter(opts *bind.TransactOpts, vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createCharacter", vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns(uint256 characterId)
func (_Contract *ContractSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateCharacter(&_Contract.TransactOpts, vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns(uint256 characterId)
func (_Contract *ContractTransactorSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateCharacter(&_Contract.TransactOpts, vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Contract *ContractTransactor) CreateThenLinkCharacter(opts *bind.TransactOpts, vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createThenLinkCharacter", vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Contract *ContractSession) CreateThenLinkCharacter(vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateThenLinkCharacter(&_Contract.TransactOpts, vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) CreateThenLinkCharacter(vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateThenLinkCharacter(&_Contract.TransactOpts, vars)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractTransactor) DeleteNote(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deleteNote", characterId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractSession) DeleteNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeleteNote(&_Contract.TransactOpts, characterId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractTransactorSession) DeleteNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeleteNote(&_Contract.TransactOpts, characterId, noteId)
}

// GrantOperatorPermissions is a paid mutator transaction binding the contract method 0x206657f2.
//
// Solidity: function grantOperatorPermissions(uint256 characterId, address operator, uint256 permissionBitMap) returns()
func (_Contract *ContractTransactor) GrantOperatorPermissions(opts *bind.TransactOpts, characterId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "grantOperatorPermissions", characterId, operator, permissionBitMap)
}

// GrantOperatorPermissions is a paid mutator transaction binding the contract method 0x206657f2.
//
// Solidity: function grantOperatorPermissions(uint256 characterId, address operator, uint256 permissionBitMap) returns()
func (_Contract *ContractSession) GrantOperatorPermissions(characterId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.GrantOperatorPermissions(&_Contract.TransactOpts, characterId, operator, permissionBitMap)
}

// GrantOperatorPermissions is a paid mutator transaction binding the contract method 0x206657f2.
//
// Solidity: function grantOperatorPermissions(uint256 characterId, address operator, uint256 permissionBitMap) returns()
func (_Contract *ContractTransactorSession) GrantOperatorPermissions(characterId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.GrantOperatorPermissions(&_Contract.TransactOpts, characterId, operator, permissionBitMap)
}

// GrantOperators4Note is a paid mutator transaction binding the contract method 0xdc17b6de.
//
// Solidity: function grantOperators4Note(uint256 characterId, uint256 noteId, address[] blocklist, address[] allowlist) returns()
func (_Contract *ContractTransactor) GrantOperators4Note(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int, blocklist []common.Address, allowlist []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "grantOperators4Note", characterId, noteId, blocklist, allowlist)
}

// GrantOperators4Note is a paid mutator transaction binding the contract method 0xdc17b6de.
//
// Solidity: function grantOperators4Note(uint256 characterId, uint256 noteId, address[] blocklist, address[] allowlist) returns()
func (_Contract *ContractSession) GrantOperators4Note(characterId *big.Int, noteId *big.Int, blocklist []common.Address, allowlist []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GrantOperators4Note(&_Contract.TransactOpts, characterId, noteId, blocklist, allowlist)
}

// GrantOperators4Note is a paid mutator transaction binding the contract method 0xdc17b6de.
//
// Solidity: function grantOperators4Note(uint256 characterId, uint256 noteId, address[] blocklist, address[] allowlist) returns()
func (_Contract *ContractTransactorSession) GrantOperators4Note(characterId *big.Int, noteId *big.Int, blocklist []common.Address, allowlist []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GrantOperators4Note(&_Contract.TransactOpts, characterId, noteId, blocklist, allowlist)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string name_, string symbol_, address linklist_, address mintNFTImpl_, address periphery_, address newbieVilla_) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, linklist_ common.Address, mintNFTImpl_ common.Address, periphery_ common.Address, newbieVilla_ common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", name_, symbol_, linklist_, mintNFTImpl_, periphery_, newbieVilla_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string name_, string symbol_, address linklist_, address mintNFTImpl_, address periphery_, address newbieVilla_) returns()
func (_Contract *ContractSession) Initialize(name_ string, symbol_ string, linklist_ common.Address, mintNFTImpl_ common.Address, periphery_ common.Address, newbieVilla_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, name_, symbol_, linklist_, mintNFTImpl_, periphery_, newbieVilla_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string name_, string symbol_, address linklist_, address mintNFTImpl_, address periphery_, address newbieVilla_) returns()
func (_Contract *ContractTransactorSession) Initialize(name_ string, symbol_ string, linklist_ common.Address, mintNFTImpl_ common.Address, periphery_ common.Address, newbieVilla_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, name_, symbol_, linklist_, mintNFTImpl_, periphery_, newbieVilla_)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkAddress(opts *bind.TransactOpts, vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkAddress", vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contract.Contract.LinkAddress(&_Contract.TransactOpts, vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contract.Contract.LinkAddress(&_Contract.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkAnyUri(opts *bind.TransactOpts, vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkAnyUri", vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contract.Contract.LinkAnyUri(&_Contract.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contract.Contract.LinkAnyUri(&_Contract.TransactOpts, vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkCharacter(opts *bind.TransactOpts, vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkCharacter", vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkCharacter(vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.LinkCharacter(&_Contract.TransactOpts, vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkCharacter(vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.LinkCharacter(&_Contract.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkERC721(opts *bind.TransactOpts, vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkERC721", vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.LinkERC721(&_Contract.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.LinkERC721(&_Contract.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkLinklist(opts *bind.TransactOpts, vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkLinklist", vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contract.Contract.LinkLinklist(&_Contract.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contract.Contract.LinkLinklist(&_Contract.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactor) LinkNote(opts *bind.TransactOpts, vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "linkNote", vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contract.Contract.LinkNote(&_Contract.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contract *ContractTransactorSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contract.Contract.LinkNote(&_Contract.TransactOpts, vars)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractTransactor) LockNote(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockNote", characterId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractSession) LockNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockNote(&_Contract.TransactOpts, characterId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contract *ContractTransactorSession) LockNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockNote(&_Contract.TransactOpts, characterId, noteId)
}

// MigrateOperatorSyncPermissions is a paid mutator transaction binding the contract method 0x5932d53d.
//
// Solidity: function migrateOperatorSyncPermissions(uint256[] characterIds) returns()
func (_Contract *ContractTransactor) MigrateOperatorSyncPermissions(opts *bind.TransactOpts, characterIds []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "migrateOperatorSyncPermissions", characterIds)
}

// MigrateOperatorSyncPermissions is a paid mutator transaction binding the contract method 0x5932d53d.
//
// Solidity: function migrateOperatorSyncPermissions(uint256[] characterIds) returns()
func (_Contract *ContractSession) MigrateOperatorSyncPermissions(characterIds []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.MigrateOperatorSyncPermissions(&_Contract.TransactOpts, characterIds)
}

// MigrateOperatorSyncPermissions is a paid mutator transaction binding the contract method 0x5932d53d.
//
// Solidity: function migrateOperatorSyncPermissions(uint256[] characterIds) returns()
func (_Contract *ContractTransactorSession) MigrateOperatorSyncPermissions(characterIds []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.MigrateOperatorSyncPermissions(&_Contract.TransactOpts, characterIds)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Contract *ContractTransactor) MintNote(opts *bind.TransactOpts, vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintNote", vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Contract *ContractSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contract.Contract.MintNote(&_Contract.TransactOpts, vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Contract *ContractTransactorSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contract.Contract.MintNote(&_Contract.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Contract *ContractTransactor) PostNote(opts *bind.TransactOpts, vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote", vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Contract *ContractSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contract.Contract.PostNote(&_Contract.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contract.Contract.PostNote(&_Contract.TransactOpts, vars)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Contract *ContractTransactor) PostNote4Address(opts *bind.TransactOpts, noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4Address", noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Contract *ContractSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Address(&_Contract.TransactOpts, noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Address(&_Contract.TransactOpts, noteData, ethAddress)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Contract *ContractTransactor) PostNote4AnyUri(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4AnyUri", postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Contract *ContractSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4AnyUri(&_Contract.TransactOpts, postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4AnyUri(&_Contract.TransactOpts, postNoteData, uri)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Contract *ContractTransactor) PostNote4Character(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4Character", postNoteData, toCharacterId)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Contract *ContractSession) PostNote4Character(postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Character(&_Contract.TransactOpts, postNoteData, toCharacterId)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4Character(postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Character(&_Contract.TransactOpts, postNoteData, toCharacterId)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Contract *ContractTransactor) PostNote4ERC721(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4ERC721", postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Contract *ContractSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4ERC721(&_Contract.TransactOpts, postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4ERC721(&_Contract.TransactOpts, postNoteData, erc721)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Contract *ContractTransactor) PostNote4Linklist(opts *bind.TransactOpts, noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4Linklist", noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Contract *ContractSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Linklist(&_Contract.TransactOpts, noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Linklist(&_Contract.TransactOpts, noteData, toLinklistId)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Contract *ContractTransactor) PostNote4Note(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postNote4Note", postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Contract *ContractSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Note(&_Contract.TransactOpts, postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Contract *ContractTransactorSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contract.Contract.PostNote4Note(&_Contract.TransactOpts, postNoteData, note)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contract *ContractTransactor) SetCharacterUri(opts *bind.TransactOpts, characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCharacterUri", characterId, newUri)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contract *ContractSession) SetCharacterUri(characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetCharacterUri(&_Contract.TransactOpts, characterId, newUri)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contract *ContractTransactorSession) SetCharacterUri(characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetCharacterUri(&_Contract.TransactOpts, characterId, newUri)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contract *ContractTransactor) SetHandle(opts *bind.TransactOpts, characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setHandle", characterId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contract *ContractSession) SetHandle(characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contract.Contract.SetHandle(&_Contract.TransactOpts, characterId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contract *ContractTransactorSession) SetHandle(characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contract.Contract.SetHandle(&_Contract.TransactOpts, characterId, newHandle)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4Address(opts *bind.TransactOpts, vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4Address", vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Address(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Address(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4Linklist(opts *bind.TransactOpts, vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4Linklist", vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Linklist(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Linklist(&_Contract.TransactOpts, vars)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contract *ContractTransactor) SetLinklistUri(opts *bind.TransactOpts, linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinklistUri", linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contract *ContractSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.Contract.SetLinklistUri(&_Contract.TransactOpts, linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contract *ContractTransactorSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contract.Contract.SetLinklistUri(&_Contract.TransactOpts, linklistId, uri)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetMintModule4Note(opts *bind.TransactOpts, vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMintModule4Note", vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contract.Contract.SetMintModule4Note(&_Contract.TransactOpts, vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contract.Contract.SetMintModule4Note(&_Contract.TransactOpts, vars)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contract *ContractTransactor) SetNoteUri(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setNoteUri", characterId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contract *ContractSession) SetNoteUri(characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetNoteUri(&_Contract.TransactOpts, characterId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contract *ContractTransactorSession) SetNoteUri(characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetNoteUri(&_Contract.TransactOpts, characterId, noteId, newUri)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contract *ContractTransactor) SetPrimaryCharacterId(opts *bind.TransactOpts, characterId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPrimaryCharacterId", characterId)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contract *ContractSession) SetPrimaryCharacterId(characterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPrimaryCharacterId(&_Contract.TransactOpts, characterId)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contract *ContractTransactorSession) SetPrimaryCharacterId(characterId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPrimaryCharacterId(&_Contract.TransactOpts, characterId)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contract *ContractTransactor) SetSocialToken(opts *bind.TransactOpts, characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSocialToken", characterId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contract *ContractSession) SetSocialToken(characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSocialToken(&_Contract.TransactOpts, characterId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contract *ContractTransactorSession) SetSocialToken(characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSocialToken(&_Contract.TransactOpts, characterId, tokenAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkAddress(opts *bind.TransactOpts, vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkAddress", vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkAddress(&_Contract.TransactOpts, vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkAddress(&_Contract.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkAnyUri(opts *bind.TransactOpts, vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkAnyUri", vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkAnyUri(&_Contract.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkAnyUri(&_Contract.TransactOpts, vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkCharacter(opts *bind.TransactOpts, vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkCharacter", vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkCharacter(vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkCharacter(&_Contract.TransactOpts, vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkCharacter(vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkCharacter(&_Contract.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkERC721(opts *bind.TransactOpts, vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkERC721", vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkERC721(&_Contract.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkERC721(&_Contract.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkLinklist(opts *bind.TransactOpts, vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkLinklist", vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkLinklist(&_Contract.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkLinklist(&_Contract.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactor) UnlinkNote(opts *bind.TransactOpts, vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlinkNote", vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkNote(&_Contract.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contract *ContractTransactorSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contract.Contract.UnlinkNote(&_Contract.TransactOpts, vars)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contract contract.
type ContractApprovalForAllIterator struct {
	Event *ContractApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApprovalForAll represents a ApprovalForAll event raised by the Contract contract.
type ContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalForAllIterator{contract: _Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApprovalForAll)
				if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) ParseApprovalForAll(log types.Log) (*ContractApprovalForAll, error) {
	event := new(ContractApprovalForAll)
	if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
