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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.createThenLinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createThenLinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"deleteNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacter\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getCharacterByHandle\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacterUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLinkModule4Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4ERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4Linklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinklistContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinklistId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linkListId\",\"type\":\"uint256\"}],\"name\":\"getLinklistType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinklistUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getNote\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.Note\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorPermissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorPermissions4Note\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPrimaryCharacterId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissionBitMap\",\"type\":\"uint256\"}],\"name\":\"grantOperatorPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissionBitMap\",\"type\":\"uint256\"}],\"name\":\"grantOperatorPermissions4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_linklistContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mintNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_periphery\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"isPrimaryCharacter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"lockNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"characterIds\",\"type\":\"uint256[]\"}],\"name\":\"migrateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MintNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mintNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"postNote4Address\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"postNote4AnyUri\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"}],\"name\":\"postNote4Character\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ERC721Struct\",\"name\":\"erc721\",\"type\":\"tuple\"}],\"name\":\"postNote4ERC721\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"}],\"name\":\"postNote4Linklist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NoteStruct\",\"name\":\"note\",\"type\":\"tuple\"}],\"name\":\"postNote4Note\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setCharacterUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"setHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4AddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4LinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Linklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setLinklistUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setMintModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setMintModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setNoteUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryCharacterId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setSocialToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615fb480620000216000396000f3fe608060405234801561001057600080fd5b506004361061045f5760003560e01c8063867884e61161024c578063c053f6b811610146578063e56f2fe4116100c3578063ef0828ab11610087578063ef0828ab14610ada578063f2ad807514610aed578063f316bacd14610b00578063f6479d7714610b13578063fe9299fb14610b2657600080fd5b8063e56f2fe414610a52578063e7a1c1c014610a65578063e985e9c514610a78578063ec81d19414610ab4578063ed84c34514610ac757600080fd5b8063d23b320b1161010a578063d23b320b146109ce578063d70e10c6146109e1578063dabb053114610a0c578063db491e8014610a1f578063dca2713514610a3f57600080fd5b8063c053f6b814610971578063c2a6fe3b14610982578063c87b56dd14610995578063cb8e757e146109a8578063cd69fe61146109bb57600080fd5b80639864c307116101d4578063a6e6178d11610198578063a6e6178d14610912578063a7ccb4bf14610925578063af90b11214610938578063b88d4fde1461094b578063b9d328451461095e57600080fd5b80639864c307146108a65780639a4dec18146108b95780639a50248d146108cc578063a22cb465146108ec578063a48047ba146108ff57600080fd5b806392f7070b1161021b57806392f7070b1461082f57806393f057e514610842578063952be0ef1461085557806395d89b411461088b57806395d9fa7d1461089357600080fd5b8063867884e6146107e35780638734bbfc146107f65780638b4ca06a1461080957806392ef445d1461081c57600080fd5b80632f745c591161035d57806347f94de7116102e55780636352211e116102a95780636352211e146107775780636bf55d5f1461078a57806370a08231146107aa57806374f345cf146107bd5780637d46a713146107d057600080fd5b806347f94de7146107185780634f6ccce71461072b5780635a936d101461073e5780635fb8818314610751578063628b644a1461076457600080fd5b8063388f50831161032c578063388f5083146106b957806340ad34d8146106cc57806342842e0e146106df57806342966c68146106f257806344b82a241461070557600080fd5b80632f745c591461065457806331b9d08c14610667578063327b2a031461069357806333f06ee6146106a657600080fd5b80631316529d116103eb5780632209d145116103af5780632209d1451461059257806323b872dd146105c857806329c301c2146105db5780632abc6bf6146105ee5780632d8723b81461061757600080fd5b80631316529d14610540578063144a3e831461055157806318160ddd14610564578063188b04b31461056c578063206657f21461057f57600080fd5b806308cb68ff1161043257806308cb68ff146104df578063095ea7b3146104f45780630c4dd5f2146105075780630d18d07a1461051a5780630ff982441461052d57600080fd5b806301ffc9a71461046457806304f3bcec1461048c57806306fdde03146104b7578063081812fc146104cc575b600080fd5b6104776104723660046148fe565b610b4f565b60405190151581526020015b60405180910390f35b60175461049f906001600160a01b031681565b6040516001600160a01b039091168152602001610483565b6104bf610b7a565b6040516104839190614973565b61049f6104da366004614986565b610c0c565b6104f26104ed3660046149b7565b610ca6565b005b6104f2610502366004614a10565b610d3f565b6104f26105153660046149b7565b610e55565b6104f2610528366004614a3c565b610f5f565b6104f261053b366004614a6c565b610fd5565b60045b604051908152602001610483565b6104bf61055f366004614986565b6110d6565b600854610543565b6104f261057a366004614a9a565b6110e1565b6104f261058d366004614ace565b6111e0565b61049f6105a0366004614a10565b6001600160a01b03918216600090815260106020908152604080832093835292905220541690565b6104f26105d6366004614b06565b611231565b6105436105e9366004614b48565b611262565b6105436105fc366004614b7c565b6001600160a01b03166000908152600c602052604090205490565b610543610625366004614b99565b6000928352601a602090815260408085209385529281528284206001600160a01b039290921684525290205490565b610543610662366004614a10565b611310565b61049f610675366004614b7c565b6001600160a01b039081166000908152601160205260409020541690565b6105436106a1366004614be4565b6113a6565b6104f26106b4366004614c7a565b61156b565b6104f26106c7366004614a9a565b611652565b6104f26106da366004614cc5565b6116a7565b6104f26106ed366004614b06565b611723565b6104f2610700366004614986565b61173e565b610543610713366004614ce1565b6117db565b6104f2610726366004614c7a565b611853565b610543610739366004614986565b6118ba565b6104f261074c366004614a6c565b61194d565b6104f261075f366004614a9a565b6119df565b6104f2610772366004614d25565b611a34565b61049f610785366004614986565b611ac4565b61079d610798366004614986565b611b3b565b6040516104839190614d77565b6105436107b8366004614b7c565b611b55565b6104f26107cb366004614dc4565b611bdc565b6104f26107de366004614a3c565b611c62565b6104f26107f1366004614cc5565b611cd3565b610477610804366004614986565b611d45565b610543610817366004614986565b611d73565b6104f261082a366004614de6565b611ddf565b61054361083d366004614e5a565b611ecb565b6104f2610850366004614a6c565b611f61565b610543610863366004614a3c565b60009182526019602090815260408084206001600160a01b0393909316845291905290205490565b6104bf611fd2565b6104f26108a1366004614a3c565b611fe1565b6104f26108b4366004614a9a565b61203a565b6105436108c7366004614be4565b61208f565b6108df6108da366004614ea0565b612199565b6040516104839190614ee1565b6104f26108fa366004614f6d565b612348565b61047761090d366004614a3c565b612353565b6104f2610920366004614c7a565b61238e565b610543610933366004614a9a565b61240f565b610543610946366004614ce1565b6124b3565b6104f2610959366004615084565b6124f3565b6104f261096c366004615101565b61252b565b6013546001600160a01b031661049f565b6104f2610990366004614dc4565b6125fc565b6104bf6109a3366004614986565b61267a565b6104f26109b6366004615101565b61271f565b6104f26109c9366004615101565b612791565b6104f26109dc366004614a9a565b6127a5565b6105436109ef366004614dc4565b6000918252600d6020908152604080842092845291905290205490565b6108df610a1a366004614986565b612832565b610a32610a2d366004614dc4565b6129b9565b6040516104839190615135565b6104bf610a4d366004614986565b612b1a565b6104f2610a603660046151cb565b612b8c565b6104f2610a73366004614a3c565b612d41565b610477610a86366004615285565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6104bf610ac2366004614986565b612e03565b6104f2610ad53660046152b3565b612e23565b6104f2610ae83660046149b7565b612e9f565b6104f2610afb366004614986565b612f10565b610543610b0e3660046152f2565b612f5f565b6104f2610b21366004614a6c565b6130bd565b61049f610b34366004614986565b6000908152600f60205260409020546001600160a01b031690565b60006001600160e01b0319821663780e9d6360e01b1480610b745750610b7482613103565b92915050565b606060008054610b899061534d565b80601f0160208091040260200160405190810160405280929190818152602001828054610bb59061534d565b8015610c025780601f10610bd757610100808354040283529160200191610c02565b820191906000526020600020905b815481529060010190602001808311610be557829003601f168201915b5050505050905090565b6000818152600260205260408120546001600160a01b0316610c8a5760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b73__$69db18565ecbefece480e92cc5f8fb1274$__63dfc34f25610ccd6020840184614b7c565b610cdd6040850160208601614b7c565b610cea6040860186615382565b60116040518663ffffffff1660e01b8152600401610d0c9594939291906153f1565b60006040518083038186803b158015610d2457600080fd5b505af4158015610d38573d6000803e3d6000fd5b5050505050565b6000610d4a82611ac4565b9050806001600160a01b0316836001600160a01b03161415610db85760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610c81565b336001600160a01b0382161480610dd45750610dd48133610a86565b610e465760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610c81565b610e508383613153565b505050565b6013546040516367880d6160e11b8152823560048201526000916001600160a01b03169063cf101ac290602401602060405180830381865afa158015610e9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec39190615430565b9050610ed08160c16131c1565b73__$69db18565ecbefece480e92cc5f8fb1274$__636252159e8335610efc6040860160208701614b7c565b610f096040870187615382565b600f6040518663ffffffff1660e01b8152600401610f2b959493929190615449565b60006040518083038186803b158015610f4357600080fd5b505af4158015610f57573d6000803e3d6000fd5b505050505050565b610f688261326d565b6000828152601860205260409020610f8090826132ee565b50610f8d8282600061330a565b806001600160a01b0316827faa9506a57073a80893a2d2fdd53f4bd49d281a8548f483ad230f2d5da7f6710c42604051610fc991815260200190565b60405180910390a35050565b610fe1813560b26131c1565b6040516331a9108f60e11b81528135600482015273__$e2c9a8f399964af47bb173600dd9e5f662$__90633019fadd9083903090636352211e90602401602060405180830381865afa15801561103b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105f9190615474565b60135485356000908152600d60209081526040808320818a0135845282529182902054825160e088901b6001600160e01b031916815286356004820152918601356024830152949091013560448201526001600160a01b0392831660648201529116608482015260a481019190915260c401610d0c565b6060610b748261267a565b6110ed813560b26131c1565b6110fa8160200135613369565b73__$e2c9a8f399964af47bb173600dd9e5f662$__63957f8a9882356020840135604085013561112d6060870187615382565b6040516331a9108f60e11b8152883560048201523090636352211e90602401602060405180830381865afa158015611169573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118d9190615474565b6013546020808b01356000908152600a9091526040908190206005015490516001600160e01b031960e08b901b168152610d0c989796959493926001600160a01b03908116921690600d90600401615491565b6111e98361326d565b8061120c57600083815260186020526040902061120690836132ee565b50611226565b600083815260186020526040902061122490836133c2565b505b610e5083838361330a565b61123b33826133d7565b6112575760405162461bcd60e51b8152600401610c81906154eb565b610e508383836134ca565b6000611270823560ec6131c1565b81356000908152600a602052604081206003018054829061129090615552565b91829055506040516342a34a5360e01b815290915073__$2873850e5b3f5697aafc2f5cd9f27cc884$__906342a34a53906112d990869085906000908190600e90600401615677565b60006040518083038186803b1580156112f157600080fd5b505af4158015611305573d6000803e3d6000fd5b509295945050505050565b600061131b83611b55565b821061137d5760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610c81565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b60006113b4833560ca6131c1565b6113ce6113c46020840184614b7c565b8360200135613671565b6013546545524337323160d01b906000906001600160a01b0316632ea24efc826113fb6020880188614b7c565b6040516001600160e01b031960e085901b16815260048101929092526001600160a01b03166024820152602087013560448201526064016020604051808303816000875af1158015611451573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114759190615430565b85356000908152600a6020526040812060030180549293509091829061149a90615552565b9182905550905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386866114cc60208c018c614b7c565b8b602001356040516020016114ff92919060609290921b6bffffffffffffffffffffffff19168252601482015260340190565b604051602081830303815290604052600e6040518763ffffffff1660e01b8152600401611531969594939291906156bc565b60006040518083038186803b15801561154957600080fd5b505af415801561155d573d6000803e3d6000fd5b509298975050505050505050565b6013546040516367880d6160e11b8152600481018590526000916001600160a01b03169063cf101ac290602401602060405180830381865afa1580156115b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d99190615430565b90506115e68160b16131c1565b601354604051633c17845760e11b81526001600160a01b039091169063782f08ae9061161a90879087908790600401615707565b600060405180830381600087803b15801561163457600080fd5b505af1158015611648573d6000803e3d6000fd5b5050505050505050565b61165e813560b96131c1565b60135460405163e4bb4f2360e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9163e4bb4f2391610d0c9185916001600160a01b0390911690600d9060040161572a565b6116b3813560b66131c1565b60135460408051635cebf08d60e11b81528335600482015260208401356024820152908301356044820152606083013560648201526001600160a01b039091166084820152600d60a482015273__$e2c9a8f399964af47bb173600dd9e5f662$__9063b9d7e11a9060c401610d0c565b610e50838383604051806020016040528060008152506124f3565b6000818152600a6020526040808220905161175c9160010190615798565b60408051918290039091206000818152600b6020908152838220829055858252600a90529182208281559092509061179760018301826147cd565b6117a56002830160006147cd565b50600060038201556004810180546001600160a01b03199081169091556005909101805490911690556117d782613721565b5050565b60006117e9833560c86131c1565b82356000908152600a60205260408120600301805467131a5b9adb1a5cdd60c21b9285929091829061181a90615552565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a6040516020016114ff91815260200190565b61185e8360b06131c1565b6000838152600a6020526040902061187a906002018383614807565b50827f17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c663199783836040516118ad929190615834565b60405180910390a2505050565b60006118c560085490565b82106119285760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610c81565b6008828154811061193b5761193b615848565b90600052602060002001549050919050565b611959813560be6131c1565b60135481356000818152600d60209081526040808320818701358085529083529281902054905163d862986560e01b8152600481019490945290850135602484015260448301919091526001600160a01b039092166064820152608481019190915273__$e2c9a8f399964af47bb173600dd9e5f662$__9063d86298659060a401610d0c565b6119eb813560bb6131c1565b6013546040516332f107bf60e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__916332f107bf91610d0c9185916001600160a01b0390911690600d9060040161585e565b611a3f8460c36131c1565b611a4b84846002613780565b611a558484613867565b6040516001626802bf60e01b0319815273__$2873850e5b3f5697aafc2f5cd9f27cc884$__9063ff97fd4190611a98908790879087908790600e906004016158db565b60006040518083038186803b158015611ab057600080fd5b505af4158015611648573d6000803e3d6000fd5b6000818152600260205260408120546001600160a01b031680610b745760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610c81565b6000818152601860205260409020606090610b749061391c565b60006001600160a01b038216611bc05760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610c81565b506001600160a01b031660009081526003602052604090205490565b611be78260c46131c1565b611bf382826003613780565b611bfd8282613867565b6000828152600e60209081526040808320848452825291829020600501805460ff60a81b1916600160a81b179055905182815283917f036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd3091015b60405180910390a25050565b611c6b8261326d565b6000828152601860205260409020611c8390826133c2565b50611c9782826001600160b01b031961330a565b806001600160a01b0316827f58f51b5bb567864de85c6a422b33491f2418924a44613d7b9341f58657bdd83342604051610fc991815260200190565b611cdf813560b86131c1565b60135481356000908152600d6020908152604080832060608601358452909152908190205490516317e9a8b160e31b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9263bf4d458892610d0c9286926001600160a01b031691906004016158fb565b600080611d5183611ac4565b6001600160a01b03166000908152600c60205260409020549290921492915050565b60135460405162fba02760e01b8152600481018390526000916001600160a01b03169062fba02790602401602060405180830381865afa158015611dbb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b749190615430565b60005b81811015610e50576000838383818110611dfe57611dfe615848565b6020908102929092013560008181526016909352604090922054919250506001600160a01b03168015611e57576000828152601860205260409020611e4390826133c2565b50611e5782826001600160b01b031961330a565b6000828152601860205260408120611e6e9061391c565b905060005b8151811015611eb657611ea684838381518110611e9257611e92615848565b602002602001015160b0600019901b61330a565b611eaf81615552565b9050611e73565b5050505080611ec490615552565b9050611de2565b6000611ed9833560c76131c1565b82356000908152600a602052604081206003018054664164647265737360c81b926001600160a01b0386169290918290611f1290615552565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a6040516020016114ff919060609190911b6bffffffffffffffffffffffff1916815260140190565b611f6d813560ba6131c1565b60135481356000908152600d6020908152604080832081860135845290915290819020549051633d7f9b3d60e11b815273__$e2c9a8f399964af47bb173600dd9e5f662$__92637aff367a92610d0c9286926001600160a01b0316919060040161594d565b606060018054610b899061534d565b611fec8260016131c1565b6040516384b44a2f60e01b8152600481018390526001600160a01b0382166024820152600a604482015273__$d8529f0216812a2dec6b96adf1943c3537$__906384b44a2f90606401610f2b565b612046813560bd6131c1565b601354604051635fe5df1d60e11b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9163bfcbbe3a91610d0c9185916001600160a01b0390911690600d90600401615995565b600061209d833560c96131c1565b601354604051635cb46be760e01b81526000600482018190528435602483015260208501356044830152634e6f746560e01b9290916001600160a01b0390911690635cb46be7906064016020604051808303816000875af1158015612106573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061212a9190615430565b85356000908152600a6020526040812060030180549293509091829061214f90615552565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878386868a600001358b602001356040516020016114ff929190918252602082015260400190565b6121a161488b565b600083836040516121b39291906159f5565b604080519182900382206000818152600b602090815283822054808352600a82529184902060c08601909452835485526001840180549396509194939290840191906121fe9061534d565b80601f016020809104026020016040519081016040528092919081815260200182805461222a9061534d565b80156122775780601f1061224c57610100808354040283529160200191612277565b820191906000526020600020905b81548152906001019060200180831161225a57829003601f168201915b505050505081526020016002820180546122909061534d565b80601f01602080910402602001604051908101604052809291908181526020018280546122bc9061534d565b80156123095780601f106122de57610100808354040283529160200191612309565b820191906000526020600020905b8154815290600101906020018083116122ec57829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015295945050505050565b6117d7338383613929565b60008281526019602090815260408083206001600160a01b03851684529091528120548015612383576001612386565b60005b949350505050565b6123998360006131c1565b60405163130f361d60e01b815273__$d8529f0216812a2dec6b96adf1943c3537$__9063130f361d906123da90869086908690600b90600a90600401615a05565b60006040518083038186803b1580156123f257600080fd5b505af4158015612406573d6000803e3d6000fd5b50505050505050565b600061242082356020840135613867565b73__$2873850e5b3f5697aafc2f5cd9f27cc884$__635be69415833560208501356124516060870160408801614b7c565b61245e6060880188615382565b6014546040516001600160e01b031960e089901b1681526124969695949392916001600160a01b031690600a90600e90600401615a33565b602060405180830381865af4158015611dbb573d6000803e3d6000fd5b60006124c1833560c66131c1565b82356000908152600a6020526040812060030180546821b430b930b1ba32b960b91b9285929091829061181a90615552565b6124fd33836133d7565b6125195760405162461bcd60e51b8152600401610c81906154eb565b612525848484846139f0565b50505050565b612537813560b56131c1565b61254981602001358260400135613867565b6040516331a9108f60e11b81528135600482015273__$e2c9a8f399964af47bb173600dd9e5f662$__9063018212d19083903090636352211e90602401602060405180830381865afa1580156125a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c79190615474565b6013546040516001600160e01b031960e086901b168152610d0c9392916001600160a01b031690600e90600d90600401615a81565b6126078260c56131c1565b61261382826004613780565b61261d8282613867565b6000828152600e60209081526040808320848452825291829020600501805460ff60a01b1916600160a01b179055905182815283917f4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf939101611c56565b6000818152600a6020526040902060020180546060919061269a9061534d565b80601f01602080910402602001604051908101604052809291908181526020018280546126c69061534d565b80156127135780601f106126e857610100808354040283529160200191612713565b820191906000526020600020905b8154815290600101906020018083116126f657829003601f168201915b50505050509050919050565b61272b813560b76131c1565b61274861273e6040830160208401614b7c565b8260400135613671565b601354604051638f3334ff60e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__91638f3334ff91610d0c9185916001600160a01b0390911690600d90600401615b00565b6127a261279d82615b63565b613a23565b50565b6127b1813560c26131c1565b6127c2813560208301356001613780565b6127d181356020830135613867565b73__$69db18565ecbefece480e92cc5f8fb1274$__6320828a02823560208401356128026060860160408701614b7c565b61280f6060870187615382565b600e6040518763ffffffff1660e01b8152600401610d0c96959493929190615c16565b61283a61488b565b600a60008381526020019081526020016000206040518060c0016040529081600082015481526020016001820180546128729061534d565b80601f016020809104026020016040519081016040528092919081815260200182805461289e9061534d565b80156128eb5780601f106128c0576101008083540402835291602001916128eb565b820191906000526020600020905b8154815290600101906020018083116128ce57829003601f168201915b505050505081526020016002820180546129049061534d565b80601f01602080910402602001604051908101604052809291908181526020018280546129309061534d565b801561297d5780601f106129525761010080835404028352916020019161297d565b820191906000526020600020905b81548152906001019060200180831161296057829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015292915050565b60408051610100808201835260008083526020808401829052606084860181905284018290526080840182905260a0840182905260c0840182905260e08401829052868252600e81528482208683528152908490208451928301855280548352600181015491830191909152600281018054939492939192840191612a3d9061534d565b80601f0160208091040260200160405190810160405280929190818152602001828054612a699061534d565b8015612ab65780601f10612a8b57610100808354040283529160200191612ab6565b820191906000526020600020905b815481529060010190602001808311612a9957829003601f168201915b505050918352505060038201546001600160a01b039081166020830152600483015481166040830152600590920154918216606082015260ff600160a01b8304811615156080830152600160a81b909204909116151560a090910152905092915050565b601354604051632b05429560e21b8152600481018390526060916001600160a01b03169063ac150a5490602401600060405180830381865afa158015612b64573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b749190810190615c57565b601454600160a81b900460ff1615808015612bb457506014546001600160a01b90910460ff16105b80612bd55750303b158015612bd55750601454600160a01b900460ff166001145b612c385760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610c81565b6014805460ff60a01b1916600160a01b1790558015612c65576014805460ff60a81b1916600160a81b1790555b612c7189898989613af8565b601380546001600160a01b03199081166001600160a01b0388811691909117909255601480548216878416179055601580548216868416179055601780549091169184169190911790556040514281527f400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf9060200160405180910390a18015612d36576014805460ff60a81b19169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b612d4a8261326d565b6001600160a01b038116612d9b5760008281526016602090815260408083205460189092529091206001600160a01b0390911690612d8890826132ee565b50612d958382600061330a565b50612dc7565b6000828152601860205260409020612db390826133c2565b50612dc782826001600160b01b031961330a565b806001600160a01b0316827f691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c41493142604051610fc991815260200190565b6000818152600a6020526040902060010180546060919061269a9061534d565b612e2c8461326d565b612e368484613867565b6000848152601a6020908152604080832086845282528083206001600160a01b0386168085529083529281902084905551838152859187917f488a41148fb8f04fec9e8e1f936444b9b70c0f084a5242092d8b4b01f6163d80910160405180910390a450505050565b612eab813560bc6131c1565b60135481356000908152600d602090815260408083208186013584529091529081902054905163bf5c00c160e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__9263bf5c00c192610d0c9286926001600160a01b03169190600401615cc4565b612f198161326d565b336000818152600c602052604080822080549085905590519092839285927fce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de9190a45050565b6000612f6d843560cb6131c1565b601354604051633610bf0960e11b815265416e7955726960d01b916000916001600160a01b0390911690636c217e1290612faf90849089908990600401615707565b6020604051808303816000875af1158015612fce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ff29190615430565b86356000908152600a6020526040812060030180549293509091829061301790615552565b919050819055905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53888386868b8b6040516020016130509291906159f5565b604051602081830303815290604052600e6040518763ffffffff1660e01b8152600401613082969594939291906156bc565b60006040518083038186803b15801561309a57600080fd5b505af41580156130ae573d6000803e3d6000fd5b50929998505050505050505050565b6130c9813560b46131c1565b6127a281356130de6040840160208501614b7c565b836040013560405180604001604052806002815260200161060f60f31b815250613b49565b60006001600160e01b031982166380ac58cd60e01b148061313457506001600160e01b03198216635b5e139f60e01b145b80610b7457506301ffc9a760e01b6001600160e01b0319831614610b74565b600081815260046020526040902080546001600160a01b0319166001600160a01b038416908117909155819061318882611ac4565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60006131cc83611ac4565b9050336001600160a01b03821614806132025750326001600160a01b03821614801561320257506015546001600160a01b031633145b8061322b57506000838152601960209081526040808320338452909152902054600190831c8116145b610e505760405162461bcd60e51b81526020600482015260136024820152722737ba22b737bab3b42832b936b4b9b9b4b7b760691b6044820152606401610c81565b600061327882611ac4565b9050336001600160a01b03821614806132ae5750326001600160a01b0382161480156132ae57506015546001600160a01b031633145b6117d75760405162461bcd60e51b81526020600482015260116024820152702737ba21b430b930b1ba32b927bbb732b960791b6044820152606401610c81565b6000613303836001600160a01b038416613dc6565b9392505050565b60008381526019602090815260408083206001600160a01b038616808552908352928190208490555183815285917f4b0dcce4c30a5691ff14b7d4a8a27e2738b66289dc90120ebbc092812941bd2791015b60405180910390a3505050565b6000818152600260205260409020546001600160a01b03166127a25760405162461bcd60e51b81526020600482015260126024820152714368617261637465724e6f7445786973747360701b6044820152606401610c81565b6000613303836001600160a01b038416613eb9565b6000818152600260205260408120546001600160a01b03166134505760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610c81565b600061345b83611ac4565b9050806001600160a01b0316846001600160a01b031614806134965750836001600160a01b031661348b84610c0c565b6001600160a01b0316145b8061238657506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff16612386565b826001600160a01b03166134dd82611ac4565b6001600160a01b0316146135415760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610c81565b6001600160a01b0382166135a35760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610c81565b6135ae838383613f08565b6135b9600082613153565b6001600160a01b03831660009081526003602052604081208054600192906135e2908490615d1c565b90915550506001600160a01b0382166000908152600360205260408120805460019290613610908490615d33565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6040516331a9108f60e11b8152600481018290526001600160a01b03831690636352211e90602401602060405180830381865afa1580156136b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136da9190615474565b6001600160a01b03166117d75760405162461bcd60e51b815260206004820152600f60248201526e5245433732314e6f7445786973747360881b6044820152606401610c81565b61372b33826133d7565b6137775760405162461bcd60e51b815260206004820152601b60248201527f4e4654426173653a204e6f744f776e65724f72417070726f76656400000000006044820152606401610c81565b6127a28161402c565b600061378b84611ac4565b9050336001600160a01b03821614806137c15750326001600160a01b0382161480156137c157506015546001600160a01b031633145b8061381b57506000848152601a602090815260408083208684528252808320338452909152902054600190831c8116148061381b57506000848152601a602090815260408083208684528252808320338452909152902054155b6125255760405162461bcd60e51b815260206004820152601e60248201527f4e6f74456e6f7567685065726d697373696f6e466f72546869734e6f746500006044820152606401610c81565b6000828152600e60209081526040808320848452909152902060050154600160a01b900460ff16156138cb5760405162461bcd60e51b815260206004820152600d60248201526c139bdd19525cd1195b195d1959609a1b6044820152606401610c81565b6000828152600a60205260409020600301548111156117d75760405162461bcd60e51b815260206004820152600d60248201526c4e6f74654e6f7445786973747360981b6044820152606401610c81565b60606000613303836140d3565b816001600160a01b0316836001600160a01b0316141561398b5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610c81565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910161335c565b6139fb8484846134ca565b613a078484848461412e565b6125255760405162461bcd60e51b8152600401610c8190615d4b565b6000601260008154613a3490615552565b91829055508251909150613a48908261422c565b604051632902741560e01b815273__$d8529f0216812a2dec6b96adf1943c3537$__90632902741590613a8a9085906001908690600b90600a90600401615d9d565b60006040518083038186803b158015613aa257600080fd5b505af4158015613ab6573d6000803e3d6000fd5b505083516001600160a01b03166000908152600c6020526040902054151591506117d790505790516001600160a01b03166000908152600c6020526040902055565b613b0484848484614246565b7f414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c308484848442604051613b3b959493929190615e46565b60405180910390a150505050565b6001600160a01b0383166000908152600c602052604090205415613bc55760405162461bcd60e51b815260206004820152602d60248201527f546172676574206164647265737320616c726561647920686173207072696d6160448201526c393c9031b430b930b1ba32b91760991b6064820152608401610c81565b6000601260008154613bd690615552565b91829055509050613be7848261422c565b73__$d8529f0216812a2dec6b96adf1943c3537$__63290274156040518060a00160405280876001600160a01b03168152602001613c2f886001600160a01b0316601461425f565b815260200160405180602001604052806000815250815260200160006001600160a01b0316815260200160405180602001604052806000815250815250600084600b600a6040518663ffffffff1660e01b8152600401613c93959493929190615d9d565b60006040518083038186803b158015613cab57600080fd5b505af4158015613cbf573d6000803e3d6000fd5b505050506001600160a01b0384166000908152600c602052604090819020829055516331a9108f60e11b81526004810186905273__$e2c9a8f399964af47bb173600dd9e5f662$__9063957f8a989087908490879087903090636352211e90602401602060405180830381865afa158015613d3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d629190615474565b6013546040516001600160e01b031960e089901b168152613d9a9695949392916001600160a01b031690600090600d90600401615e80565b60006040518083038186803b158015613db257600080fd5b505af4158015612d36573d6000803e3d6000fd5b60008181526001830160205260408120548015613eaf576000613dea600183615d1c565b8554909150600090613dfe90600190615d1c565b9050818114613e63576000866000018281548110613e1e57613e1e615848565b9060005260206000200154905080876000018481548110613e4157613e41615848565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080613e7457613e74615ed8565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610b74565b6000915050610b74565b6000818152600183016020526040812054613f0057508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610b74565b506000610b74565b6000818152601860205260408120613f1f906143fa565b600083815260186020526040812091925090613f3a9061391c565b905060005b82811015613fe857600084815260196020526040812083518290859085908110613f6b57613f6b615848565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002081905550613fd5828281518110613fac57613fac615848565b6020026020010151601860008781526020019081526020016000206132ee90919063ffffffff16565b5080613fe081615552565b915050613f3f565b506001600160a01b0385166000908152600c602052604090205415614021576001600160a01b0385166000908152600c60205260408120555b610d38858585614404565b600061403782611ac4565b905061404581600084613f08565b614050600083613153565b6001600160a01b0381166000908152600360205260408120805460019290614079908490615d1c565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b60608160000180548060200260200160405190810160405280929190818152602001828054801561271357602002820191906000526020600020905b81548152602001906001019080831161410f5750505050509050919050565b60006001600160a01b0384163b1561422157604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290614172903390899088908890600401615eee565b6020604051808303816000875af19250505080156141ad575060408051601f3d908101601f191682019092526141aa91810190615f2b565b60015b614207573d8080156141db576040519150601f19603f3d011682016040523d82523d6000602084013e6141e0565b606091505b5080516141ff5760405162461bcd60e51b8152600401610c8190615d4b565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050612386565b506001949350505050565b6117d78282604051806020016040528060008152506144bc565b61425260008585614807565b50610d3860018383614807565b6060600061426e836002615f48565b614279906002615d33565b6001600160401b0381111561429057614290614f99565b6040519080825280601f01601f1916602001820160405280156142ba576020820181803683370190505b509050600360fc1b816000815181106142d5576142d5615848565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061430457614304615848565b60200101906001600160f81b031916908160001a9053506000614328846002615f48565b614333906001615d33565b90505b60018111156143ab576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061436757614367615848565b1a60f81b82828151811061437d5761437d615848565b60200101906001600160f81b031916908160001a90535060049490941c936143a481615f67565b9050614336565b5083156133035760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610c81565b6000610b74825490565b6001600160a01b03831661445f5761445a81600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b614482565b816001600160a01b0316836001600160a01b0316146144825761448283826144ef565b6001600160a01b03821661449957610e508161458c565b826001600160a01b0316826001600160a01b031614610e5057610e50828261463b565b6144c6838361467f565b6144d3600084848461412e565b610e505760405162461bcd60e51b8152600401610c8190615d4b565b600060016144fc84611b55565b6145069190615d1c565b600083815260076020526040902054909150808214614559576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b60085460009061459e90600190615d1c565b600083815260096020526040812054600880549394509092849081106145c6576145c6615848565b9060005260206000200154905080600883815481106145e7576145e7615848565b600091825260208083209091019290925582815260099091526040808220849055858252812055600880548061461f5761461f615ed8565b6001900381819060005260206000200160009055905550505050565b600061464683611b55565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b6001600160a01b0382166146d55760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610c81565b6000818152600260205260409020546001600160a01b03161561473a5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610c81565b61474660008383613f08565b6001600160a01b038216600090815260036020526040812080546001929061476f908490615d33565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b5080546147d99061534d565b6000825580601f106147e9575050565b601f0160209004906000526020600020908101906127a291906148d3565b8280546148139061534d565b90600052602060002090601f016020900481019282614835576000855561487b565b82601f1061484e5782800160ff1982351617855561487b565b8280016001018555821561487b579182015b8281111561487b578235825591602001919060010190614860565b506148879291506148d3565b5090565b6040518060c001604052806000815260200160608152602001606081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b5b8082111561488757600081556001016148d4565b6001600160e01b0319811681146127a257600080fd5b60006020828403121561491057600080fd5b8135613303816148e8565b60005b8381101561493657818101518382015260200161491e565b838111156125255750506000910152565b6000815180845261495f81602086016020860161491b565b601f01601f19169290920160200192915050565b6020815260006133036020830184614947565b60006020828403121561499857600080fd5b5035919050565b6000606082840312156149b157600080fd5b50919050565b6000602082840312156149c957600080fd5b81356001600160401b038111156149df57600080fd5b6123868482850161499f565b6001600160a01b03811681146127a257600080fd5b8035614a0b816149eb565b919050565b60008060408385031215614a2357600080fd5b8235614a2e816149eb565b946020939093013593505050565b60008060408385031215614a4f57600080fd5b823591506020830135614a61816149eb565b809150509250929050565b600060608284031215614a7e57600080fd5b613303838361499f565b6000608082840312156149b157600080fd5b600060208284031215614aac57600080fd5b81356001600160401b03811115614ac257600080fd5b61238684828501614a88565b600080600060608486031215614ae357600080fd5b833592506020840135614af5816149eb565b929592945050506040919091013590565b600080600060608486031215614b1b57600080fd5b8335614b26816149eb565b92506020840135614af5816149eb565b600060e082840312156149b157600080fd5b600060208284031215614b5a57600080fd5b81356001600160401b03811115614b7057600080fd5b61238684828501614b36565b600060208284031215614b8e57600080fd5b8135613303816149eb565b600080600060608486031215614bae57600080fd5b83359250602084013591506040840135614bc7816149eb565b809150509250925092565b6000604082840312156149b157600080fd5b60008060608385031215614bf757600080fd5b82356001600160401b03811115614c0d57600080fd5b614c1985828601614b36565b925050614c298460208501614bd2565b90509250929050565b60008083601f840112614c4457600080fd5b5081356001600160401b03811115614c5b57600080fd5b602083019150836020828501011115614c7357600080fd5b9250929050565b600080600060408486031215614c8f57600080fd5b8335925060208401356001600160401b03811115614cac57600080fd5b614cb886828701614c32565b9497909650939450505050565b600060808284031215614cd757600080fd5b6133038383614a88565b60008060408385031215614cf457600080fd5b82356001600160401b03811115614d0a57600080fd5b614d1685828601614b36565b95602094909401359450505050565b60008060008060608587031215614d3b57600080fd5b843593506020850135925060408501356001600160401b03811115614d5f57600080fd5b614d6b87828801614c32565b95989497509550505050565b6020808252825182820181905260009190848201906040850190845b81811015614db85783516001600160a01b031683529284019291840191600101614d93565b50909695505050505050565b60008060408385031215614dd757600080fd5b50508035926020909101359150565b60008060208385031215614df957600080fd5b82356001600160401b0380821115614e1057600080fd5b818501915085601f830112614e2457600080fd5b813581811115614e3357600080fd5b8660208260051b8501011115614e4857600080fd5b60209290920196919550909350505050565b60008060408385031215614e6d57600080fd5b82356001600160401b03811115614e8357600080fd5b614e8f85828601614b36565b9250506020830135614a61816149eb565b60008060208385031215614eb357600080fd5b82356001600160401b03811115614ec957600080fd5b614ed585828601614c32565b90969095509350505050565b60208152815160208201526000602083015160c06040840152614f0760e0840182614947565b90506040840151601f19848303016060850152614f248282614947565b91505060608401516080840152608084015160018060a01b0380821660a08601528060a08701511660c086015250508091505092915050565b80358015158114614a0b57600080fd5b60008060408385031215614f8057600080fd5b8235614f8b816149eb565b9150614c2960208401614f5d565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715614fd157614fd1614f99565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614fff57614fff614f99565b604052919050565b60006001600160401b0382111561502057615020614f99565b50601f01601f191660200190565b600082601f83011261503f57600080fd5b813561505261504d82615007565b614fd7565b81815284602083860101111561506757600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806080858703121561509a57600080fd5b84356150a5816149eb565b935060208501356150b5816149eb565b92506040850135915060608501356001600160401b038111156150d757600080fd5b6150e38782880161502e565b91505092959194509250565b600060a082840312156149b157600080fd5b60006020828403121561511357600080fd5b81356001600160401b0381111561512957600080fd5b612386848285016150ef565b60208152815160208201526020820151604082015260006040830151610100806060850152615168610120850183614947565b9150606085015160018060a01b0380821660808701528060808801511660a0870152505060a08501516151a660c08601826001600160a01b03169052565b5060c085015180151560e08601525060e0850151801515858301525090949350505050565b60008060008060008060008060c0898b0312156151e757600080fd5b88356001600160401b03808211156151fe57600080fd5b61520a8c838d01614c32565b909a50985060208b013591508082111561522357600080fd5b506152308b828c01614c32565b9097509550506040890135615244816149eb565b93506060890135615254816149eb565b92506080890135615264816149eb565b915060a0890135615274816149eb565b809150509295985092959890939650565b6000806040838503121561529857600080fd5b82356152a3816149eb565b91506020830135614a61816149eb565b600080600080608085870312156152c957600080fd5b843593506020850135925060408501356152e2816149eb565b9396929550929360600135925050565b60008060006040848603121561530757600080fd5b83356001600160401b038082111561531e57600080fd5b61532a87838801614b36565b9450602086013591508082111561534057600080fd5b50614cb886828701614c32565b600181811c9082168061536157607f821691505b602082108114156149b157634e487b7160e01b600052602260045260246000fd5b6000808335601e1984360301811261539957600080fd5b8301803591506001600160401b038211156153b357600080fd5b602001915036819003821315614c7357600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0386811682528516602082015260806040820181905260009061541e90830185876153c8565b90508260608301529695505050505050565b60006020828403121561544257600080fd5b5051919050565b8581526001600160a01b038516602082015260806040820181905260009061541e90830185876153c8565b60006020828403121561548657600080fd5b8151613303816149eb565b60006101008b83528a60208401528960408401528060608401526154b8818401898b6153c8565b6001600160a01b03978816608085015295871660a084015250509190931660c082015260e0019190915295945050505050565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60006000198214156155665761556661553c565b5060010190565b6000808335601e1984360301811261558457600080fd5b83016020810192503590506001600160401b038111156155a357600080fd5b803603831315614c7357600080fd5b8035825260006155c5602083018361556d565b60e060208601526155da60e0860182846153c8565b91505060408301356155eb816149eb565b6001600160a01b038181166040870152615608606086018661556d565b9250868403606088015261561d8484836153c8565b93505060808501359150615630826149eb565b16608085015261564360a084018461556d565b85830360a08701526156568382846153c8565b9250505061566660c08401614f5d565b151560c08501528091505092915050565b60c08152600061568a60c08301886155b2565b602083810197909752604083019590955250606081019290925281830360808301526000835260a09091015201919050565b60c0815260006156cf60c08301896155b2565b87602084015286604084015285606084015282810360808401526156f38186614947565b9150508260a0830152979650505050505050565b8381526040602082015260006157216040830184866153c8565b95945050505050565b606081528335606082015260006020850135615745816149eb565b6001600160a01b038181166080850152604087013560a085015261576c606088018861556d565b9250608060c086015261578360e0860184836153c8565b96909116602085015250505060400152919050565b600080835481600182811c9150808316806157b457607f831692505b60208084108214156157d457634e487b7160e01b86526022600452602486fd5b8180156157e857600181146157f957615826565b60ff19861689528489019650615826565b60008a81526020902060005b8681101561581e5781548b820152908501908301615805565b505084890196505b509498975050505050505050565b6020815260006123866020830184866153c8565b634e487b7160e01b600052603260045260246000fd5b60608152833560608201526000615878602086018661556d565b60808085015261588c60e0850182846153c8565b915050604086013560a08401526158a6606087018761556d565b848303605f190160c08601526158bd8382846153c8565b6001600160a01b039790971660208601525050505060400152919050565b85815284602082015260806040820152600061541e6080830185876153c8565b8335815260c081016020850135615911816149eb565b6001600160a01b039081166020840152604086810135908401526060958601359583019590955292909316608084015260a09092019190915290565b8335815260a081016020850135615963816149eb565b6001600160a01b0390811660208401526040958601359583019590955292909316606084015260809092019190915290565b606081528335606082015260208401356080820152604084013560a082015260006159c3606086018661556d565b608060c08501526159d860e0850182846153c8565b6001600160a01b0396909616602085015250505060400152919050565b8183823760009101908152919050565b858152608060208201526000615a1f6080830186886153c8565b604083019490945250606001529392505050565b888152876020820152600060018060a01b03808916604084015260e06060840152615a6260e08401888a6153c8565b951660808301525060a081019290925260c09091015295945050505050565b60a08152853560a0820152602086013560c0820152604086013560e082015260608601356101008201526000615aba608088018861556d565b60a0610120850152615ad1610140850182846153c8565b6001600160a01b0398891660208601529690971660408401525050606081019290925260809091015292915050565b606081528335606082015260006020850135615b1b816149eb565b60018060a01b038082166080850152604087013560a0850152606087013560c0850152615b4b608088018861556d565b925060a060e0860152615783610100860184836153c8565b600060a08236031215615b7557600080fd5b615b7d614faf565b615b8683614a00565b815260208301356001600160401b0380821115615ba257600080fd5b615bae3683870161502e565b60208401526040850135915080821115615bc757600080fd5b615bd33683870161502e565b6040840152615be460608601614a00565b60608401526080850135915080821115615bfd57600080fd5b50615c0a3682860161502e565b60808301525092915050565b86815285602082015260018060a01b038516604082015260a060608201526000615c4460a0830185876153c8565b9050826080830152979650505050505050565b600060208284031215615c6957600080fd5b81516001600160401b03811115615c7f57600080fd5b8201601f81018413615c9057600080fd5b8051615c9e61504d82615007565b818152856020838501011115615cb357600080fd5b61572182602083016020860161491b565b60608152833560608201526000615cde602086018661556d565b60606080850152615cf360c0850182846153c8565b60409788013560a08601526001600160a01b039690961660208501525050509092019190915290565b600082821015615d2e57615d2e61553c565b500390565b60008219821115615d4657615d4661553c565b500190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60a08152600060018060a01b038088511660a0840152602088015160a060c0850152615dcd610140850182614947565b90506040890151609f19808684030160e0870152615deb8383614947565b92508360608c01511661010087015260808b0151935080868403016101208701525050615e188183614947565b92505050615e2a602083018715159052565b8460408301528360608301528260808301529695505050505050565b606081526000615e5a6060830187896153c8565b8281036020840152615e6d8186886153c8565b9150508260408301529695505050505050565b60006101008a8352896020840152886040840152806060840152615ea681840189614947565b6001600160a01b03978816608085015295871660a084015250509190931660c082015260e00191909152949350505050565b634e487b7160e01b600052603160045260246000fd5b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090615f2190830184614947565b9695505050505050565b600060208284031215615f3d57600080fd5b8151613303816148e8565b6000816000190483118215151615615f6257615f6261553c565b500290565b600081615f7657615f7661553c565b50600019019056fea2646970667358221220657f9ce63d49477996b1bc8c4ae2cee851fba5c3439d836e2364d044862d4d4a64736f6c634300080a0033",
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

// GetOperatorPermissions4Note is a free data retrieval call binding the contract method 0x2d8723b8.
//
// Solidity: function getOperatorPermissions4Note(uint256 characterId, uint256 noteId, address operator) view returns(uint256)
func (_Contract *ContractCaller) GetOperatorPermissions4Note(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorPermissions4Note", characterId, noteId, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorPermissions4Note is a free data retrieval call binding the contract method 0x2d8723b8.
//
// Solidity: function getOperatorPermissions4Note(uint256 characterId, uint256 noteId, address operator) view returns(uint256)
func (_Contract *ContractSession) GetOperatorPermissions4Note(characterId *big.Int, noteId *big.Int, operator common.Address) (*big.Int, error) {
	return _Contract.Contract.GetOperatorPermissions4Note(&_Contract.CallOpts, characterId, noteId, operator)
}

// GetOperatorPermissions4Note is a free data retrieval call binding the contract method 0x2d8723b8.
//
// Solidity: function getOperatorPermissions4Note(uint256 characterId, uint256 noteId, address operator) view returns(uint256)
func (_Contract *ContractCallerSession) GetOperatorPermissions4Note(characterId *big.Int, noteId *big.Int, operator common.Address) (*big.Int, error) {
	return _Contract.Contract.GetOperatorPermissions4Note(&_Contract.CallOpts, characterId, noteId, operator)
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

// IsOperator is a free data retrieval call binding the contract method 0xa48047ba.
//
// Solidity: function isOperator(uint256 characterId, address operator) view returns(bool)
func (_Contract *ContractCaller) IsOperator(opts *bind.CallOpts, characterId *big.Int, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isOperator", characterId, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xa48047ba.
//
// Solidity: function isOperator(uint256 characterId, address operator) view returns(bool)
func (_Contract *ContractSession) IsOperator(characterId *big.Int, operator common.Address) (bool, error) {
	return _Contract.Contract.IsOperator(&_Contract.CallOpts, characterId, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0xa48047ba.
//
// Solidity: function isOperator(uint256 characterId, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsOperator(characterId *big.Int, operator common.Address) (bool, error) {
	return _Contract.Contract.IsOperator(&_Contract.CallOpts, characterId, operator)
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

// AddOperator is a paid mutator transaction binding the contract method 0x7d46a713.
//
// Solidity: function addOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactor) AddOperator(opts *bind.TransactOpts, characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addOperator", characterId, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x7d46a713.
//
// Solidity: function addOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractSession) AddOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddOperator(&_Contract.TransactOpts, characterId, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x7d46a713.
//
// Solidity: function addOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactorSession) AddOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddOperator(&_Contract.TransactOpts, characterId, operator)
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
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
func (_Contract *ContractTransactor) CreateCharacter(opts *bind.TransactOpts, vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createCharacter", vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
func (_Contract *ContractSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateCharacter(&_Contract.TransactOpts, vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
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

// GrantOperatorPermissions4Note is a paid mutator transaction binding the contract method 0xed84c345.
//
// Solidity: function grantOperatorPermissions4Note(uint256 characterId, uint256 noteId, address operator, uint256 permissionBitMap) returns()
func (_Contract *ContractTransactor) GrantOperatorPermissions4Note(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "grantOperatorPermissions4Note", characterId, noteId, operator, permissionBitMap)
}

// GrantOperatorPermissions4Note is a paid mutator transaction binding the contract method 0xed84c345.
//
// Solidity: function grantOperatorPermissions4Note(uint256 characterId, uint256 noteId, address operator, uint256 permissionBitMap) returns()
func (_Contract *ContractSession) GrantOperatorPermissions4Note(characterId *big.Int, noteId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.GrantOperatorPermissions4Note(&_Contract.TransactOpts, characterId, noteId, operator, permissionBitMap)
}

// GrantOperatorPermissions4Note is a paid mutator transaction binding the contract method 0xed84c345.
//
// Solidity: function grantOperatorPermissions4Note(uint256 characterId, uint256 noteId, address operator, uint256 permissionBitMap) returns()
func (_Contract *ContractTransactorSession) GrantOperatorPermissions4Note(characterId *big.Int, noteId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.GrantOperatorPermissions4Note(&_Contract.TransactOpts, characterId, noteId, operator, permissionBitMap)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Contract *ContractSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Contract *ContractTransactorSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
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

// MigrateOperator is a paid mutator transaction binding the contract method 0x92ef445d.
//
// Solidity: function migrateOperator(uint256[] characterIds) returns()
func (_Contract *ContractTransactor) MigrateOperator(opts *bind.TransactOpts, characterIds []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "migrateOperator", characterIds)
}

// MigrateOperator is a paid mutator transaction binding the contract method 0x92ef445d.
//
// Solidity: function migrateOperator(uint256[] characterIds) returns()
func (_Contract *ContractSession) MigrateOperator(characterIds []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.MigrateOperator(&_Contract.TransactOpts, characterIds)
}

// MigrateOperator is a paid mutator transaction binding the contract method 0x92ef445d.
//
// Solidity: function migrateOperator(uint256[] characterIds) returns()
func (_Contract *ContractTransactorSession) MigrateOperator(characterIds []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.MigrateOperator(&_Contract.TransactOpts, characterIds)
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

// RemoveOperator is a paid mutator transaction binding the contract method 0x0d18d07a.
//
// Solidity: function removeOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactor) RemoveOperator(opts *bind.TransactOpts, characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeOperator", characterId, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x0d18d07a.
//
// Solidity: function removeOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractSession) RemoveOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveOperator(&_Contract.TransactOpts, characterId, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x0d18d07a.
//
// Solidity: function removeOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactorSession) RemoveOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveOperator(&_Contract.TransactOpts, characterId, operator)
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

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactor) SetOperator(opts *bind.TransactOpts, characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOperator", characterId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractSession) SetOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOperator(&_Contract.TransactOpts, characterId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Contract *ContractTransactorSession) SetOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOperator(&_Contract.TransactOpts, characterId, operator)
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
