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

// DataTypessetLinkModule4CharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4CharacterData struct {
	CharacterId        *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4ERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4ERC721Data struct {
	TokenAddress       common.Address
	TokenId            *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4LinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4LinklistData struct {
	LinklistId         *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4NoteData struct {
	CharacterId        *big.Int
	NoteId             *big.Int
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canCreate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.createThenLinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createThenLinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"deleteNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacter\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getCharacterByHandle\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacterUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLinkModule4Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4ERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4Linklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinklistContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinklistId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linkListId\",\"type\":\"uint256\"}],\"name\":\"getLinklistType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinklistUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getNote\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.Note\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPrimaryCharacterId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_linklistContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mintNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_periphery\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"isPrimaryCharacter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"lockNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MintNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mintNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"postNote4Address\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"postNote4AnyUri\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"}],\"name\":\"postNote4Character\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ERC721Struct\",\"name\":\"erc721\",\"type\":\"tuple\"}],\"name\":\"postNote4ERC721\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"}],\"name\":\"postNote4Linklist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NoteStruct\",\"name\":\"note\",\"type\":\"tuple\"}],\"name\":\"postNote4Note\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setCharacterUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"setHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4AddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4CharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Character\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4ERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4ERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4LinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Linklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setLinklistUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setMintModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setMintModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setNoteUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryCharacterId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setProfileUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setSocialToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615a6a80620000216000396000f3fe608060405234801561001057600080fd5b50600436106104495760003560e01c8063867884e611610241578063c87b56dd1161013b578063e56f2fe4116100c3578063f2ad807511610087578063f2ad807514610a37578063f316bacd14610a4a578063f6479d7714610a5d578063fd2d866f14610a70578063fe9299fb14610a8357600080fd5b8063e56f2fe4146109af578063e7a1c1c0146109c2578063e985e9c5146109d5578063ec81d19414610a11578063ef0828ab14610a2457600080fd5b8063d70e10c61161010a578063d70e10c61461092b578063dabb053114610956578063db491e8014610969578063db8c198d14610989578063dca271351461099c57600080fd5b8063c87b56dd146108f2578063cb8e757e14610905578063cd69fe61146108bb578063d23b320b1461091857600080fd5b80639a50248d116101c9578063b88d4fde1161018d578063b88d4fde14610895578063b9d32845146108a8578063bd5f69cb146108bb578063c053f6b8146108ce578063c2a6fe3b146108df57600080fd5b80639a50248d14610829578063a22cb46514610849578063a6e6178d1461085c578063a7ccb4bf1461086f578063af90b1121461088257600080fd5b806393f057e51161021057806393f057e5146107d557806395d89b41146107e857806395d9fa7d146107f05780639864c307146108035780639a4dec181461081657600080fd5b8063867884e6146107895780638734bbfc1461079c5780638b4ca06a146107af57806392f7070b146107c257600080fd5b806331b9d08c116103525780634f6ccce7116102da57806369492c971161029e57806369492c971461073a57806370a082311461074d57806374f345cf146107605780637c392b51146106c85780637daca6d01461077357600080fd5b80634f6ccce7146106db5780635a936d10146106ee5780635fb8818314610701578063628b644a146107145780636352211e1461072757600080fd5b806340ad34d81161032157806340ad34d81461067c57806342842e0e1461068f57806342966c68146106a257806344b82a24146106b557806347f94de7146106c857600080fd5b806331b9d08c14610617578063327b2a031461064357806333f06ee614610656578063388f50831461066957600080fd5b80631316529d116103d55780632209d145116103a45780632209d1451461057f57806323b872dd146105b557806329c301c2146105c85780632abc6bf6146105db5780632f745c591461060457600080fd5b80631316529d14610540578063144a3e831461055157806318160ddd14610564578063188b04b31461056c57600080fd5b8063081812fc1161041c578063081812fc146104df57806308cb68ff146104f2578063095ea7b3146105075780630c4dd5f21461051a5780630ff982441461052d57600080fd5b806301ffc9a71461044e57806304f3bcec1461047657806305f63c8a146104a157806306fdde03146104ca575b600080fd5b61046161045c3660046144a1565b610aac565b60405190151581526020015b60405180910390f35b601754610489906001600160a01b031681565b6040516001600160a01b03909116815260200161046d565b6104896104af3660046144be565b6000908152601660205260409020546001600160a01b031690565b6104d2610ad7565b60405161046d919061452f565b6104896104ed3660046144be565b610b69565b61050561050036600461455a565b610c03565b005b6105056105153660046145b3565b610c9c565b61050561052836600461455a565b610db2565b61050561053b3660046145df565b610e17565b60045b60405190815260200161046d565b6104d261055f3660046144be565b610f16565b600854610543565b61050561057a36600461460d565b610f21565b61048961058d3660046145b3565b6001600160a01b03918216600090815260106020908152604080832093835292905220541690565b6105056105c3366004614641565b61101e565b6105436105d6366004614694565b61104f565b6105436105e93660046146c8565b6001600160a01b03166000908152600c602052604090205490565b6105436106123660046145b3565b6110fb565b6104896106253660046146c8565b6001600160a01b039081166000908152601160205260409020541690565b6105436106513660046146f7565b611191565b61050561066436600461478d565b611354565b61050561067736600461460d565b6113c8565b61050561068a3660046147d8565b611482565b61050561069d366004614641565b6114fc565b6105056106b03660046144be565b611517565b6105436106c33660046147f4565b6115b4565b6105056106d636600461478d565b61162a565b6105436106e93660046144be565b61166a565b6105056106fc3660046145df565b6116fd565b61050561070f36600461460d565b61178d565b610505610722366004614838565b611847565b6104896107353660046144be565b6118d3565b61050561074836600461460d565b61194a565b61054361075b3660046146c8565b611a82565b61050561076e36600461488a565b611b09565b610461610781366004614997565b600192915050565b6105056107973660046147d8565b611b81565b6104616107aa3660046144be565b611bf1565b6105436107bd3660046144be565b611c1f565b6105436107d03660046149e8565b611c8b565b6105056107e33660046145df565b611d1f565b6104d2611d8e565b6105056107fe366004614a1d565b611d9d565b61050561081136600461460d565b611e23565b6105436108243660046146f7565b611edd565b61083c610837366004614a42565b611fe5565b60405161046d9190614a83565b610505610857366004614b0f565b612194565b61050561086a36600461478d565b61219f565b61054361087d36600461460d565b612215565b6105436108903660046147f4565b6122b9565b6105056108a3366004614b3b565b6122f7565b6105056108b6366004614bb8565b61232f565b6105056108c9366004614bb8565b6123fe565b6013546001600160a01b0316610489565b6105056108ed36600461488a565b612412565b6104d26109003660046144be565b612482565b610505610913366004614bb8565b612527565b61050561092636600461460d565b6125fe565b61054361093936600461488a565b6000918252600d6020908152604080842092845291905290205490565b61083c6109643660046144be565b612678565b61097c61097736600461488a565b6127ff565b60405161046d9190614bec565b61050561099736600461460d565b612960565b6104d26109aa3660046144be565b6129aa565b6105056109bd366004614c82565b612a1c565b6105056109d0366004614a1d565b612bd1565b6104616109e3366004614d3c565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6104d2610a1f3660046144be565b612be4565b610505610a3236600461455a565b612c04565b610505610a453660046144be565b612c73565b610543610a58366004614d6a565b612cc2565b610505610a6b3660046145df565b612e1e565b610505610a7e36600461455a565b612e58565b610489610a913660046144be565b6000908152600f60205260409020546001600160a01b031690565b60006001600160e01b0319821663780e9d6360e01b1480610ad15750610ad182612ed1565b92915050565b606060008054610ae690614dc5565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1290614dc5565b8015610b5f5780601f10610b3457610100808354040283529160200191610b5f565b820191906000526020600020905b815481529060010190602001808311610b4257829003601f168201915b5050505050905090565b6000818152600260205260408120546001600160a01b0316610be75760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b73__$ccaefba5abaae7361fb981898db80d4305$__63dfc34f25610c2a60208401846146c8565b610c3a60408501602086016146c8565b610c476040860186614dfa565b60116040518663ffffffff1660e01b8152600401610c69959493929190614e69565b60006040518083038186803b158015610c8157600080fd5b505af4158015610c95573d6000803e3d6000fd5b5050505050565b6000610ca7826118d3565b9050806001600160a01b0316836001600160a01b03161415610d155760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610bde565b336001600160a01b0382161480610d315750610d3181336109e3565b610da35760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610bde565b610dad8383612f21565b505050565b610dbc8135612f8f565b73__$ccaefba5abaae7361fb981898db80d4305$__636252159e8235610de860408501602086016146c8565b610df56040860186614dfa565b600f6040518663ffffffff1660e01b8152600401610c69959493929190614ea8565b610e21813561304f565b6040516331a9108f60e11b81528135600482015273__$5abde9af541dd76010bdab2fdbd2e12ad0$__90633019fadd9083903090636352211e90602401602060405180830381865afa158015610e7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9f9190614ed3565b60135485356000908152600d60209081526040808320818a0135845282529182902054825160e088901b6001600160e01b031916815286356004820152918601356024830152949091013560448201526001600160a01b0392831660648201529116608482015260a481019190915260c401610c69565b6060610ad182612482565b610f2b813561304f565b610f3881602001356130d0565b73__$5abde9af541dd76010bdab2fdbd2e12ad0$__63957f8a98823560208401356040850135610f6b6060870187614dfa565b6040516331a9108f60e11b8152883560048201523090636352211e90602401602060405180830381865afa158015610fa7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fcb9190614ed3565b6013546020808b01356000908152600a9091526040908190206005015490516001600160e01b031960e08b901b168152610c69989796959493926001600160a01b03908116921690600d90600401614ef0565b6110283382613129565b6110445760405162461bcd60e51b8152600401610bde90614f4a565b610dad838383613220565b600061105b82356133c7565b81356000908152600a602052604081206003018054829061107b90614fb1565b91829055506040516342a34a5360e01b815290915073__$28febe8707c2533835cf7587e52ef64d24$__906342a34a53906110c490869085906000908190600e906004016150d6565b60006040518083038186803b1580156110dc57600080fd5b505af41580156110f0573d6000803e3d6000fd5b509295945050505050565b600061110683611a82565b82106111685760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610bde565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b600061119d83356133c7565b6111b76111ad60208401846146c8565b8360200135613468565b6013546545524337323160d01b906000906001600160a01b0316632ea24efc826111e460208801886146c8565b6040516001600160e01b031960e085901b16815260048101929092526001600160a01b03166024820152602087013560448201526064016020604051808303816000875af115801561123a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061125e919061511b565b85356000908152600a6020526040812060030180549293509091829061128390614fb1565b9182905550905073__$28febe8707c2533835cf7587e52ef64d24$__6342a34a53878386866112b560208c018c6146c8565b8b602001356040516020016112e892919060609290921b6bffffffffffffffffffffffff19168252601482015260340190565b604051602081830303815290604052600e6040518763ffffffff1660e01b815260040161131a96959493929190615134565b60006040518083038186803b15801561133257600080fd5b505af4158015611346573d6000803e3d6000fd5b509298975050505050505050565b61135d83612f8f565b601354604051633c17845760e11b81526001600160a01b039091169063782f08ae906113919086908690869060040161517f565b600060405180830381600087803b1580156113ab57600080fd5b505af11580156113bf573d6000803e3d6000fd5b50505050505050565b6113d2813561304f565b6040516331a9108f60e11b81528135600482015273__$5abde9af541dd76010bdab2fdbd2e12ad0$__90631e3981939083903090636352211e90602401602060405180830381865afa15801561142c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114509190614ed3565b6013546040516001600160e01b031960e086901b168152610c699392916001600160a01b031690600d906004016151a2565b61148c81356133c7565b60135460408051635cebf08d60e11b81528335600482015260208401356024820152908301356044820152606083013560648201526001600160a01b039091166084820152600d60a482015273__$5abde9af541dd76010bdab2fdbd2e12ad0$__9063b9d7e11a9060c401610c69565b610dad838383604051806020016040528060008152506122f7565b6000818152600a602052604080822090516115359160010190615219565b60408051918290039091206000818152600b6020908152838220829055858252600a90529182208281559092509061157060018301826142fc565b61157e6002830160006142fc565b50600060038201556004810180546001600160a01b03199081169091556005909101805490911690556115b082613518565b5050565b60006115c083356133c7565b82356000908152600a60205260408120600301805467131a5b9adb1a5cdd60c21b928592909182906115f190614fb1565b919050819055905073__$28febe8707c2533835cf7587e52ef64d24$__6342a34a53878386868a6040516020016112e891815260200190565b610dad8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061357792505050565b600061167560085490565b82106116d85760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610bde565b600882815481106116eb576116eb6152b5565b90600052602060002001549050919050565b611707813561304f565b60135481356000818152600d60209081526040808320818701358085529083529281902054905163d862986560e01b8152600481019490945290850135602484015260448301919091526001600160a01b039092166064820152608481019190915273__$5abde9af541dd76010bdab2fdbd2e12ad0$__9063d86298659060a401610c69565b611797813561304f565b6040516331a9108f60e11b81528135600482015273__$5abde9af541dd76010bdab2fdbd2e12ad0$__906331ba880f9083903090636352211e90602401602060405180830381865afa1580156117f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118159190614ed3565b6013546040516001600160e01b031960e086901b168152610c699392916001600160a01b031690600d906004016152cb565b6118508461304f565b61185a84846135d6565b6040516001626802bf60e01b0319815273__$28febe8707c2533835cf7587e52ef64d24$__9063ff97fd419061189d908790879087908790600e90600401615352565b60006040518083038186803b1580156118b557600080fd5b505af41580156118c9573d6000803e3d6000fd5b5050505050505050565b6000818152600260205260408120546001600160a01b031680610ad15760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610bde565b61195760208201826146c8565b6040516331a9108f60e11b8152602083013560048201526001600160a01b039190911690636352211e90602401602060405180830381865afa1580156119a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c59190614ed3565b6001600160a01b0316336001600160a01b031614611a165760405162461bcd60e51b815260206004820152600e60248201526d2737ba22a9219b9918a7bbb732b960911b6044820152606401610bde565b73__$ccaefba5abaae7361fb981898db80d4305$__63384afeb3611a3d60208401846146c8565b6020840135611a5260608601604087016146c8565b611a5f6060870187614dfa565b60106040518763ffffffff1660e01b8152600401610c6996959493929190615372565b60006001600160a01b038216611aed5760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610bde565b506001600160a01b031660009081526003602052604090205490565b611b12826133c7565b611b1c82826135d6565b6000828152600e60209081526040808320848452825291829020600501805460ff60a81b1916600160a81b179055905182815283917f036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd3091015b60405180910390a25050565b611b8b813561304f565b60135481356000908152600d6020908152604080832060608601358452909152908190205490516317e9a8b160e31b815273__$5abde9af541dd76010bdab2fdbd2e12ad0$__9263bf4d458892610c699286926001600160a01b031691906004016153b9565b600080611bfd836118d3565b6001600160a01b03166000908152600c60205260409020549290921492915050565b60135460405162fba02760e01b8152600481018390526000916001600160a01b03169062fba02790602401602060405180830381865afa158015611c67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad1919061511b565b6000611c9783356133c7565b82356000908152600a602052604081206003018054664164647265737360c81b926001600160a01b0386169290918290611cd090614fb1565b919050819055905073__$28febe8707c2533835cf7587e52ef64d24$__6342a34a53878386868a6040516020016112e8919060609190911b6bffffffffffffffffffffffff1916815260140190565b611d29813561304f565b60135481356000908152600d6020908152604080832081860135845290915290819020549051633d7f9b3d60e11b815273__$5abde9af541dd76010bdab2fdbd2e12ad0$__92637aff367a92610c699286926001600160a01b0316919060040161540b565b606060018054610ae690614dc5565b611da68261304f565b6040516384b44a2f60e01b8152600481018390526001600160a01b0382166024820152600a604482015273__$5f4b6d0e8d26e72e412687f6a7f2becc77$__906384b44a2f9060640160006040518083038186803b158015611e0757600080fd5b505af4158015611e1b573d6000803e3d6000fd5b505050505050565b611e2d813561304f565b6040516331a9108f60e11b81528135600482015273__$5abde9af541dd76010bdab2fdbd2e12ad0$__906372c61ab09083903090636352211e90602401602060405180830381865afa158015611e87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eab9190614ed3565b6013546040516001600160e01b031960e086901b168152610c699392916001600160a01b031690600d90600401615453565b6000611ee983356133c7565b601354604051635cb46be760e01b81526000600482018190528435602483015260208501356044830152634e6f746560e01b9290916001600160a01b0390911690635cb46be7906064016020604051808303816000875af1158015611f52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f76919061511b565b85356000908152600a60205260408120600301805492935090918290611f9b90614fb1565b919050819055905073__$28febe8707c2533835cf7587e52ef64d24$__6342a34a53878386868a600001358b602001356040516020016112e8929190918252602082015260400190565b611fed614336565b60008383604051611fff9291906154bc565b604080519182900382206000818152600b602090815283822054808352600a82529184902060c086019094528354855260018401805493965091949392908401919061204a90614dc5565b80601f016020809104026020016040519081016040528092919081815260200182805461207690614dc5565b80156120c35780601f10612098576101008083540402835291602001916120c3565b820191906000526020600020905b8154815290600101906020018083116120a657829003601f168201915b505050505081526020016002820180546120dc90614dc5565b80601f016020809104026020016040519081016040528092919081815260200182805461210890614dc5565b80156121555780601f1061212a57610100808354040283529160200191612155565b820191906000526020600020905b81548152906001019060200180831161213857829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015295945050505050565b6115b033838361368b565b6121a88361304f565b60405163130f361d60e01b815273__$5f4b6d0e8d26e72e412687f6a7f2becc77$__9063130f361d906121e990869086908690600b90600a906004016154cc565b60006040518083038186803b15801561220157600080fd5b505af41580156113bf573d6000803e3d6000fd5b6000612226823560208401356135d6565b73__$28febe8707c2533835cf7587e52ef64d24$__635be694158335602085013561225760608701604088016146c8565b6122646060880188614dfa565b6014546040516001600160e01b031960e089901b16815261229c9695949392916001600160a01b031690600a90600e906004016154fa565b602060405180830381865af4158015611c67573d6000803e3d6000fd5b60006122c583356133c7565b82356000908152600a6020526040812060030180546821b430b930b1ba32b960b91b928592909182906115f190614fb1565b6123013383613129565b61231d5760405162461bcd60e51b8152600401610bde90614f4a565b6123298484848461375a565b50505050565b61233981356133c7565b61234b816020013582604001356135d6565b6040516331a9108f60e11b81528135600482015273__$5abde9af541dd76010bdab2fdbd2e12ad0$__9063018212d19083903090636352211e90602401602060405180830381865afa1580156123a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123c99190614ed3565b6013546040516001600160e01b031960e086901b168152610c699392916001600160a01b031690600e90600d90600401615548565b61240f61240a826155c7565b61378d565b50565b61241b826133c7565b61242582826135d6565b6000828152600e60209081526040808320848452825291829020600501805460ff60a01b1916600160a01b179055905182815283917f4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf939101611b75565b6000818152600a602052604090206002018054606091906124a290614dc5565b80601f01602080910402602001604051908101604052809291908181526020018280546124ce90614dc5565b801561251b5780601f106124f05761010080835404028352916020019161251b565b820191906000526020600020905b8154815290600101906020018083116124fe57829003601f168201915b50505050509050919050565b612531813561304f565b61254e61254460408301602084016146c8565b8260400135613468565b6040516331a9108f60e11b81528135600482015273__$5abde9af541dd76010bdab2fdbd2e12ad0$__9063e7a15c419083903090636352211e90602401602060405180830381865afa1580156125a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125cc9190614ed3565b6013546040516001600160e01b031960e086901b168152610c699392916001600160a01b031690600d9060040161567a565b612608813561304f565b612617813560208301356135d6565b73__$ccaefba5abaae7361fb981898db80d4305$__6320828a028235602084013561264860608601604087016146c8565b6126556060870187614dfa565b600e6040518763ffffffff1660e01b8152600401610c69969594939291906156df565b612680614336565b600a60008381526020019081526020016000206040518060c0016040529081600082015481526020016001820180546126b890614dc5565b80601f01602080910402602001604051908101604052809291908181526020018280546126e490614dc5565b80156127315780601f1061270657610100808354040283529160200191612731565b820191906000526020600020905b81548152906001019060200180831161271457829003601f168201915b5050505050815260200160028201805461274a90614dc5565b80601f016020809104026020016040519081016040528092919081815260200182805461277690614dc5565b80156127c35780601f10612798576101008083540402835291602001916127c3565b820191906000526020600020905b8154815290600101906020018083116127a657829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015292915050565b60408051610100808201835260008083526020808401829052606084860181905284018290526080840182905260a0840182905260c0840182905260e08401829052868252600e8152848220868352815290849020845192830185528054835260018101549183019190915260028101805493949293919284019161288390614dc5565b80601f01602080910402602001604051908101604052809291908181526020018280546128af90614dc5565b80156128fc5780601f106128d1576101008083540402835291602001916128fc565b820191906000526020600020905b8154815290600101906020018083116128df57829003601f168201915b505050918352505060038201546001600160a01b039081166020830152600483015481166040830152600590920154918216606082015260ff600160a01b8304811615156080830152600160a81b909204909116151560a090910152905092915050565b61296a813561304f565b612979813560208301356135d6565b73__$ccaefba5abaae7361fb981898db80d4305$__631f2ffb698235602084013561264860608601604087016146c8565b601354604051632b05429560e21b8152600481018390526060916001600160a01b03169063ac150a5490602401600060405180830381865afa1580156129f4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ad1919081019061570d565b601454600160a81b900460ff1615808015612a4457506014546001600160a01b90910460ff16105b80612a655750303b158015612a655750601454600160a01b900460ff166001145b612ac85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610bde565b6014805460ff60a01b1916600160a01b1790558015612af5576014805460ff60a81b1916600160a81b1790555b612b0189898989613860565b601380546001600160a01b03199081166001600160a01b0388811691909117909255601480548216878416179055601580548216868416179055601780549091169184169190911790556040514281527f400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf9060200160405180910390a18015612bc6576014805460ff60a81b19169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b612bda8261304f565b6115b082826138b1565b6000818152600a602052604090206001018054606091906124a290614dc5565b612c0e813561304f565b60135481356000908152600d602090815260408083208186013584529091529081902054905163bf5c00c160e01b815273__$5abde9af541dd76010bdab2fdbd2e12ad0$__9263bf5c00c192610c699286926001600160a01b0316919060040161577a565b612c7c8161304f565b336000818152600c602052604080822080549085905590519092839285927fce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de9190a45050565b6000612cce84356133c7565b601354604051633610bf0960e11b815265416e7955726960d01b916000916001600160a01b0390911690636c217e1290612d109084908990899060040161517f565b6020604051808303816000875af1158015612d2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d53919061511b565b86356000908152600a60205260408120600301805492935090918290612d7890614fb1565b919050819055905073__$28febe8707c2533835cf7587e52ef64d24$__6342a34a53888386868b8b604051602001612db19291906154bc565b604051602081830303815290604052600e6040518763ffffffff1660e01b8152600401612de396959493929190615134565b60006040518083038186803b158015612dfb57600080fd5b505af4158015612e0f573d6000803e3d6000fd5b50929998505050505050505050565b61240f8135612e3360408401602085016146c8565b836040013560405180604001604052806002815260200161060f60f31b815250613915565b612e62813561304f565b73__$5f4b6d0e8d26e72e412687f6a7f2becc77$__631dc831338235612e8e60408501602086016146c8565b612e9b6040860186614dfa565b86356000908152600a60205260409081902090516001600160e01b031960e088901b168152610c69959493929190600401614ea8565b60006001600160e01b031982166380ac58cd60e01b1480612f0257506001600160e01b03198216635b5e139f60e01b145b80610ad157506301ffc9a760e01b6001600160e01b0319831614610ad1565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190612f56826118d3565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6013546040516331a9108f60e11b8152600481018390526001600160a01b0390911690636352211e90602401602060405180830381865afa158015612fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ffc9190614ed3565b6001600160a01b0316336001600160a01b03161461240f5760405162461bcd60e51b815260206004820152601060248201526f2737ba2634b735a634b9ba27bbb732b960811b6044820152606401610bde565b600061305a826118d3565b9050336001600160a01b03821614806130905750326001600160a01b03821614801561309057506015546001600160a01b031633145b6115b05760405162461bcd60e51b81526020600482015260116024820152702737ba21b430b930b1ba32b927bbb732b960791b6044820152606401610bde565b6000818152600260205260409020546001600160a01b031661240f5760405162461bcd60e51b81526020600482015260126024820152714368617261637465724e6f7445786973747360701b6044820152606401610bde565b6000818152600260205260408120546001600160a01b03166131a25760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610bde565b60006131ad836118d3565b9050806001600160a01b0316846001600160a01b031614806131e85750836001600160a01b03166131dd84610b69565b6001600160a01b0316145b8061321857506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b949350505050565b826001600160a01b0316613233826118d3565b6001600160a01b0316146132975760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610bde565b6001600160a01b0382166132f95760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610bde565b613304838383613b9b565b61330f600082612f21565b6001600160a01b03831660009081526003602052604081208054600192906133389084906157d2565b90915550506001600160a01b03821660009081526003602052604081208054600192906133669084906157e9565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b60006133d2826118d3565b9050336001600160a01b038216148061340157506000828152601660205260409020546001600160a01b031633145b806130905750326001600160a01b03821614801561309057506015546001600160a01b031633146115b05760405162461bcd60e51b81526020600482015260116024820152702737ba21b430b930b1ba32b927bbb732b960791b6044820152606401610bde565b6040516331a9108f60e11b8152600481018290526001600160a01b03831690636352211e90602401602060405180830381865afa1580156134ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134d19190614ed3565b6001600160a01b03166115b05760405162461bcd60e51b815260206004820152600f60248201526e5245433732314e6f7445786973747360881b6044820152606401610bde565b6135223382613129565b61356e5760405162461bcd60e51b815260206004820152601b60248201527f4e4654426173653a204e6f744f776e65724f72417070726f76656400000000006044820152606401610bde565b61240f81613c06565b613580826133c7565b6000828152600a6020908152604090912082516135a59260029092019184019061437e565b50817f17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c663199782604051611b75919061452f565b6000828152600e60209081526040808320848452909152902060050154600160a01b900460ff161561363a5760405162461bcd60e51b815260206004820152600d60248201526c139bdd19525cd1195b195d1959609a1b6044820152606401610bde565b6000828152600a60205260409020600301548111156115b05760405162461bcd60e51b815260206004820152600d60248201526c4e6f74654e6f7445786973747360981b6044820152606401610bde565b816001600160a01b0316836001600160a01b031614156136ed5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610bde565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b613765848484613220565b61377184848484613cad565b6123295760405162461bcd60e51b8152600401610bde90615801565b60125461379b9060016157e9565b601281905581516137ab91613dab565b601254604051632902741560e01b815273__$5f4b6d0e8d26e72e412687f6a7f2becc77$__916329027415916137ef91859160019190600b90600a90600401615853565b60006040518083038186803b15801561380757600080fd5b505af415801561381b573d6000803e3d6000fd5b505082516001600160a01b03166000908152600c60205260409020541515915061240f90505760125490516001600160a01b03166000908152600c6020526040902055565b61386c84848484613ef9565b7f414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c3084848484426040516138a39594939291906158fc565b60405180910390a150505050565b60008281526016602090815260409182902080546001600160a01b0319166001600160a01b038516908117909155915142815284917f691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931910160405180910390a35050565b61391e8461304f565b6001600160a01b0383166000908152600c60205260409020541561399a5760405162461bcd60e51b815260206004820152602d60248201527f546172676574206164647265737320616c726561647920686173207072696d6160448201526c393c9031b430b930b1ba32b91760991b6064820152608401610bde565b60006012600081546139ab90614fb1565b918290555090506139bc8482613dab565b73__$5f4b6d0e8d26e72e412687f6a7f2becc77$__63290274156040518060a00160405280876001600160a01b03168152602001613a04886001600160a01b03166014613f12565b815260200160405180602001604052806000815250815260200160006001600160a01b0316815260200160405180602001604052806000815250815250600084600b600a6040518663ffffffff1660e01b8152600401613a68959493929190615853565b60006040518083038186803b158015613a8057600080fd5b505af4158015613a94573d6000803e3d6000fd5b505050506001600160a01b0384166000908152600c602052604090819020829055516331a9108f60e11b81526004810186905273__$5abde9af541dd76010bdab2fdbd2e12ad0$__9063957f8a989087908490879087903090636352211e90602401602060405180830381865afa158015613b13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b379190614ed3565b6013546040516001600160e01b031960e089901b168152613b6f9695949392916001600160a01b031690600090600d90600401615936565b60006040518083038186803b158015613b8757600080fd5b505af4158015612bc6573d6000803e3d6000fd5b6000818152601660205260409020546001600160a01b031615613bc357613bc38160006138b1565b6001600160a01b0383166000908152600c602052604090205415613bfb576001600160a01b0383166000908152600c60205260408120555b610dad8383836140b4565b6000613c11826118d3565b9050613c1f81600084613b9b565b613c2a600083612f21565b6001600160a01b0381166000908152600360205260408120805460019290613c539084906157d2565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b60006001600160a01b0384163b15613da057604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290613cf190339089908890889060040161598e565b6020604051808303816000875af1925050508015613d2c575060408051601f3d908101601f19168201909252613d29918101906159cb565b60015b613d86573d808015613d5a576040519150601f19603f3d011682016040523d82523d6000602084013e613d5f565b606091505b508051613d7e5760405162461bcd60e51b8152600401610bde90615801565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050613218565b506001949350505050565b6001600160a01b038216613e015760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610bde565b6000818152600260205260409020546001600160a01b031615613e665760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610bde565b613e7260008383613b9b565b6001600160a01b0382166000908152600360205260408120805460019290613e9b9084906157e9565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b613f0560008585614402565b50610c9560018383614402565b60606000613f218360026159e8565b613f2c9060026157e9565b6001600160401b03811115613f4357613f436148ac565b6040519080825280601f01601f191660200182016040528015613f6d576020820181803683370190505b509050600360fc1b81600081518110613f8857613f886152b5565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110613fb757613fb76152b5565b60200101906001600160f81b031916908160001a9053506000613fdb8460026159e8565b613fe69060016157e9565b90505b600181111561405e576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061401a5761401a6152b5565b1a60f81b828281518110614030576140306152b5565b60200101906001600160f81b031916908160001a90535060049490941c9361405781615a07565b9050613fe9565b5083156140ad5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610bde565b9392505050565b6001600160a01b03831661410f5761410a81600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b614132565b816001600160a01b0316836001600160a01b03161461413257614132838261416c565b6001600160a01b03821661414957610dad81614209565b826001600160a01b0316826001600160a01b031614610dad57610dad82826142b8565b6000600161417984611a82565b61418391906157d2565b6000838152600760205260409020549091508082146141d6576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b60085460009061421b906001906157d2565b60008381526009602052604081205460088054939450909284908110614243576142436152b5565b906000526020600020015490508060088381548110614264576142646152b5565b600091825260208083209091019290925582815260099091526040808220849055858252812055600880548061429c5761429c615a1e565b6001900381819060005260206000200160009055905550505050565b60006142c383611a82565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b50805461430890614dc5565b6000825580601f10614318575050565b601f01602090049060005260206000209081019061240f9190614476565b6040518060c001604052806000815260200160608152602001606081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b82805461438a90614dc5565b90600052602060002090601f0160209004810192826143ac57600085556143f2565b82601f106143c557805160ff19168380011785556143f2565b828001600101855582156143f2579182015b828111156143f25782518255916020019190600101906143d7565b506143fe929150614476565b5090565b82805461440e90614dc5565b90600052602060002090601f01602090048101928261443057600085556143f2565b82601f106144495782800160ff198235161785556143f2565b828001600101855582156143f2579182015b828111156143f257823582559160200191906001019061445b565b5b808211156143fe5760008155600101614477565b6001600160e01b03198116811461240f57600080fd5b6000602082840312156144b357600080fd5b81356140ad8161448b565b6000602082840312156144d057600080fd5b5035919050565b60005b838110156144f25781810151838201526020016144da565b838111156123295750506000910152565b6000815180845261451b8160208601602086016144d7565b601f01601f19169290920160200192915050565b6020815260006140ad6020830184614503565b60006060828403121561455457600080fd5b50919050565b60006020828403121561456c57600080fd5b81356001600160401b0381111561458257600080fd5b61321884828501614542565b6001600160a01b038116811461240f57600080fd5b80356145ae8161458e565b919050565b600080604083850312156145c657600080fd5b82356145d18161458e565b946020939093013593505050565b6000606082840312156145f157600080fd5b6140ad8383614542565b60006080828403121561455457600080fd5b60006020828403121561461f57600080fd5b81356001600160401b0381111561463557600080fd5b613218848285016145fb565b60008060006060848603121561465657600080fd5b83356146618161458e565b925060208401356146718161458e565b929592945050506040919091013590565b600060e0828403121561455457600080fd5b6000602082840312156146a657600080fd5b81356001600160401b038111156146bc57600080fd5b61321884828501614682565b6000602082840312156146da57600080fd5b81356140ad8161458e565b60006040828403121561455457600080fd5b6000806060838503121561470a57600080fd5b82356001600160401b0381111561472057600080fd5b61472c85828601614682565b92505061473c84602085016146e5565b90509250929050565b60008083601f84011261475757600080fd5b5081356001600160401b0381111561476e57600080fd5b60208301915083602082850101111561478657600080fd5b9250929050565b6000806000604084860312156147a257600080fd5b8335925060208401356001600160401b038111156147bf57600080fd5b6147cb86828701614745565b9497909650939450505050565b6000608082840312156147ea57600080fd5b6140ad83836145fb565b6000806040838503121561480757600080fd5b82356001600160401b0381111561481d57600080fd5b61482985828601614682565b95602094909401359450505050565b6000806000806060858703121561484e57600080fd5b843593506020850135925060408501356001600160401b0381111561487257600080fd5b61487e87828801614745565b95989497509550505050565b6000806040838503121561489d57600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156148e4576148e46148ac565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614912576149126148ac565b604052919050565b60006001600160401b03821115614933576149336148ac565b50601f01601f191660200190565b600082601f83011261495257600080fd5b81356149656149608261491a565b6148ea565b81815284602083860101111561497a57600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156149aa57600080fd5b82356001600160401b038111156149c057600080fd5b6149cc85828601614941565b92505060208301356149dd8161458e565b809150509250929050565b600080604083850312156149fb57600080fd5b82356001600160401b03811115614a1157600080fd5b6149cc85828601614682565b60008060408385031215614a3057600080fd5b8235915060208301356149dd8161458e565b60008060208385031215614a5557600080fd5b82356001600160401b03811115614a6b57600080fd5b614a7785828601614745565b90969095509350505050565b60208152815160208201526000602083015160c06040840152614aa960e0840182614503565b90506040840151601f19848303016060850152614ac68282614503565b91505060608401516080840152608084015160018060a01b0380821660a08601528060a08701511660c086015250508091505092915050565b803580151581146145ae57600080fd5b60008060408385031215614b2257600080fd5b8235614b2d8161458e565b915061473c60208401614aff565b60008060008060808587031215614b5157600080fd5b8435614b5c8161458e565b93506020850135614b6c8161458e565b92506040850135915060608501356001600160401b03811115614b8e57600080fd5b614b9a87828801614941565b91505092959194509250565b600060a0828403121561455457600080fd5b600060208284031215614bca57600080fd5b81356001600160401b03811115614be057600080fd5b61321884828501614ba6565b60208152815160208201526020820151604082015260006040830151610100806060850152614c1f610120850183614503565b9150606085015160018060a01b0380821660808701528060808801511660a0870152505060a0850151614c5d60c08601826001600160a01b03169052565b5060c085015180151560e08601525060e0850151801515858301525090949350505050565b60008060008060008060008060c0898b031215614c9e57600080fd5b88356001600160401b0380821115614cb557600080fd5b614cc18c838d01614745565b909a50985060208b0135915080821115614cda57600080fd5b50614ce78b828c01614745565b9097509550506040890135614cfb8161458e565b93506060890135614d0b8161458e565b92506080890135614d1b8161458e565b915060a0890135614d2b8161458e565b809150509295985092959890939650565b60008060408385031215614d4f57600080fd5b8235614d5a8161458e565b915060208301356149dd8161458e565b600080600060408486031215614d7f57600080fd5b83356001600160401b0380821115614d9657600080fd5b614da287838801614682565b94506020860135915080821115614db857600080fd5b506147cb86828701614745565b600181811c90821680614dd957607f821691505b6020821081141561455457634e487b7160e01b600052602260045260246000fd5b6000808335601e19843603018112614e1157600080fd5b8301803591506001600160401b03821115614e2b57600080fd5b60200191503681900382131561478657600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b03868116825285166020820152608060408201819052600090614e969083018587614e40565b90508260608301529695505050505050565b8581526001600160a01b0385166020820152608060408201819052600090614e969083018587614e40565b600060208284031215614ee557600080fd5b81516140ad8161458e565b60006101008b83528a6020840152896040840152806060840152614f17818401898b614e40565b6001600160a01b03978816608085015295871660a084015250509190931660c082015260e0019190915295945050505050565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b6000600019821415614fc557614fc5614f9b565b5060010190565b6000808335601e19843603018112614fe357600080fd5b83016020810192503590506001600160401b0381111561500257600080fd5b80360383131561478657600080fd5b8035825260006150246020830183614fcc565b60e0602086015261503960e086018284614e40565b915050604083013561504a8161458e565b6001600160a01b0381811660408701526150676060860186614fcc565b9250868403606088015261507c848483614e40565b9350506080850135915061508f8261458e565b1660808501526150a260a0840184614fcc565b85830360a08701526150b5838284614e40565b925050506150c560c08401614aff565b151560c08501528091505092915050565b60c0815260006150e960c0830188615011565b602083810197909752604083019590955250606081019290925281830360808301526000835260a09091015201919050565b60006020828403121561512d57600080fd5b5051919050565b60c08152600061514760c0830189615011565b876020840152866040840152856060840152828103608084015261516b8186614503565b9150508260a0830152979650505050505050565b838152604060208201526000615199604083018486614e40565b95945050505050565b6080815284356080820152600060208601356151bd8161458e565b6001600160a01b0381811660a0850152604088013560c08501526151e46060890189614fcc565b9250608060e08601526151fc61010086018483614e40565b978216602086015250949094166040830152506060015250919050565b600080835481600182811c91508083168061523557607f831692505b602080841082141561525557634e487b7160e01b86526022600452602486fd5b818015615269576001811461527a576152a7565b60ff198616895284890196506152a7565b60008a81526020902060005b8681101561529f5781548b820152908501908301615286565b505084890196505b509498975050505050505050565b634e487b7160e01b600052603260045260246000fd5b608081528435608082015260006152e56020870187614fcc565b608060a08501526152fb61010085018284614e40565b915050604087013560c08401526153156060880188614fcc565b848303607f190160e086015261532c838284614e40565b6001600160a01b0398891660208701529690971660408501525050506060015292915050565b858152846020820152608060408201526000614e96608083018587614e40565b6001600160a01b038781168252602082018790528516604082015260a0606082018190526000906153a69083018587614e40565b9050826080830152979650505050505050565b8335815260c0810160208501356153cf8161458e565b6001600160a01b039081166020840152604086810135908401526060958601359583019590955292909316608084015260a09092019190915290565b8335815260a0810160208501356154218161458e565b6001600160a01b0390811660208401526040958601359583019590955292909316606084015260809092019190915290565b6080815284356080820152602085013560a0820152604085013560c082015260006154816060870187614fcc565b608060e085015261549761010085018284614e40565b6001600160a01b03978816602086015295909616604084015250506060015292915050565b8183823760009101908152919050565b8581526080602082015260006154e6608083018688614e40565b604083019490945250606001529392505050565b888152876020820152600060018060a01b03808916604084015260e0606084015261552960e08401888a614e40565b951660808301525060a081019290925260c09091015295945050505050565b60a08152853560a0820152602086013560c0820152604086013560e0820152606086013561010082015260006155816080880188614fcc565b60a061012085015261559861014085018284614e40565b6001600160a01b0398891660208601529690971660408401525050606081019290925260809091015292915050565b600060a082360312156155d957600080fd5b6155e16148c2565b6155ea836145a3565b815260208301356001600160401b038082111561560657600080fd5b61561236838701614941565b6020840152604085013591508082111561562b57600080fd5b61563736838701614941565b6040840152615648606086016145a3565b6060840152608085013591508082111561566157600080fd5b5061566e36828601614941565b60808301525092915050565b6080815284356080820152600060208601356156958161458e565b6001600160a01b0381811660a0850152604088013560c0850152606088013560e08501526156c66080890189614fcc565b925060a06101008601526151fc61012086018483614e40565b86815285602082015260018060a01b038516604082015260a0606082015260006153a660a083018587614e40565b60006020828403121561571f57600080fd5b81516001600160401b0381111561573557600080fd5b8201601f8101841361574657600080fd5b80516157546149608261491a565b81815285602083850101111561576957600080fd5b6151998260208301602086016144d7565b606081528335606082015260006157946020860186614fcc565b606060808501526157a960c085018284614e40565b60409788013560a08601526001600160a01b039690961660208501525050509092019190915290565b6000828210156157e4576157e4614f9b565b500390565b600082198211156157fc576157fc614f9b565b500190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60a08152600060018060a01b038088511660a0840152602088015160a060c0850152615883610140850182614503565b90506040890151609f19808684030160e08701526158a18383614503565b92508360608c01511661010087015260808b01519350808684030161012087015250506158ce8183614503565b925050506158e0602083018715159052565b8460408301528360608301528260808301529695505050505050565b606081526000615910606083018789614e40565b8281036020840152615923818688614e40565b9150508260408301529695505050505050565b60006101008a835289602084015288604084015280606084015261595c81840189614503565b6001600160a01b03978816608085015295871660a084015250509190931660c082015260e00191909152949350505050565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906159c190830184614503565b9695505050505050565b6000602082840312156159dd57600080fd5b81516140ad8161448b565b6000816000190483118215151615615a0257615a02614f9b565b500290565b600081615a1657615a16614f9b565b506000190190565b634e487b7160e01b600052603160045260246000fdfea264697066735822122063add97a540a5a9d6293ee8e8ae88a24eda86f90f746e18c6a41f9db745a390564736f6c634300080a0033",
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

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Contract *ContractCaller) CanCreate(opts *bind.CallOpts, handle string, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "canCreate", handle, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Contract *ContractSession) CanCreate(handle string, account common.Address) (bool, error) {
	return _Contract.Contract.CanCreate(&_Contract.CallOpts, handle, account)
}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Contract *ContractCallerSession) CanCreate(handle string, account common.Address) (bool, error) {
	return _Contract.Contract.CanCreate(&_Contract.CallOpts, handle, account)
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

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Contract *ContractCaller) GetOperator(opts *bind.CallOpts, characterId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperator", characterId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Contract *ContractSession) GetOperator(characterId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetOperator(&_Contract.CallOpts, characterId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Contract *ContractCallerSession) GetOperator(characterId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetOperator(&_Contract.CallOpts, characterId)
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

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Contract *ContractTransactor) CreateProfile(opts *bind.TransactOpts, vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createProfile", vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Contract *ContractSession) CreateProfile(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateProfile(&_Contract.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) CreateProfile(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contract.Contract.CreateProfile(&_Contract.TransactOpts, vars)
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

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4Character(opts *bind.TransactOpts, vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4Character", vars)
}

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4Character(vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Character(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4Character(vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Character(&_Contract.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4ERC721(opts *bind.TransactOpts, vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4ERC721", vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4ERC721(&_Contract.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4ERC721(&_Contract.TransactOpts, vars)
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

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactor) SetLinkModule4Note(opts *bind.TransactOpts, vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLinkModule4Note", vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Note(&_Contract.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contract *ContractTransactorSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Contract.Contract.SetLinkModule4Note(&_Contract.TransactOpts, vars)
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

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Contract *ContractTransactor) SetProfileUri(opts *bind.TransactOpts, profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setProfileUri", profileId, newUri)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Contract *ContractSession) SetProfileUri(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetProfileUri(&_Contract.TransactOpts, profileId, newUri)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Contract *ContractTransactorSession) SetProfileUri(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetProfileUri(&_Contract.TransactOpts, profileId, newUri)
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
