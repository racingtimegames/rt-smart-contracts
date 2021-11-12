// SPDX-License-Identifier: Apache-2.0

package t_test

import (
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime/scripts"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime/transactions"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	seed := "WYf13zjfhNhmT6S734pphp8hEZwhCf3A1KGMm7S3hxEXmxyt5fFXGXA9zeDR7wEt"

	racingtime.Address = ""
	racingtime.PrivateKey = ""

	acc, prv := transactions.CreateAddress(seed, crypto.ECDSA_secp256k1, crypto.SHA3_256)

	println("New Account Address:", acc)
	println("New Account Signature Algorithm:", crypto.ECDSA_secp256k1.String())
	println("New Account Hash Algorithm:", crypto.SHA3_256.String())
	println("New Account Private Key:", prv)
}

func TestSetupAccount(t *testing.T) {
	racingtime.Address = ""
	racingtime.PrivateKey = ""
	racingtime.SigAlgo = crypto.ECDSA_P256
	racingtime.HashAlgo = crypto.SHA3_256
	racingtime.KeyIndex = 0

	_result := transactions.SetupAccount()
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result.Status)
}

func TestRacingTimeDeployContracts(t *testing.T) {
	racingtime.Address = ""
	racingtime.PrivateKey = ""
	racingtime.SigAlgo = crypto.ECDSA_secp256k1
	racingtime.HashAlgo = crypto.SHA3_256
	racingtime.KeyIndex = 0

	_result := transactions.DeployContract()
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result.Status)
}

func TestRacingTimeUpdateContracts(t *testing.T) {
	racingtime.Address = ""
	racingtime.PrivateKey = ""
	racingtime.SigAlgo = crypto.ECDSA_secp256k1
	racingtime.HashAlgo = crypto.SHA3_256
	racingtime.KeyIndex = 0

	_result := transactions.UpdateContract()
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result.Status)
}

func TestMintNFT(t *testing.T) {
	racingtime.Address = ""
	racingtime.PrivateKey = ""
	racingtime.SigAlgo = crypto.ECDSA_secp256k1
	racingtime.HashAlgo = crypto.SHA3_256
	racingtime.KeyIndex = 0

	//test nft info
	nft := &transactions.NFTInfo{RewardID: 50003,
		TypeID:       2,
		SerialNumber: 1,
		Ipfs:         "https://bafkreie3o25bs7uxcijlo5sa2mtbetvvp544hvrypt7nthl23ahxuuttzm.ipfs.dweb.link"}

	_result := transactions.MintNFT(nft)
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result)
}

func TestAccountNFT(t *testing.T) {
	searchAddress := ""

	_result := scripts.AccountInfo(searchAddress)

	println(lib.CadenceValueToJsonString(_result))
}

func TestTransferNFT(t *testing.T) {
	racingtime.Address = ""
	racingtime.PrivateKey = ""
	racingtime.SigAlgo = crypto.ECDSA_secp256k1
	racingtime.HashAlgo = crypto.SHA3_256
	racingtime.KeyIndex = 0

	ToAddress := "0xf0da3bfb6f06d7d4"
	NFTId := uint64(2)

	_result := transactions.TransferRacingTimeNFT(ToAddress, NFTId)
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result)
}

func TestBurnNFT(t *testing.T) {
	racingtime.Address = ""
	racingtime.PrivateKey = ""
	racingtime.SigAlgo = crypto.ECDSA_secp256k1
	racingtime.HashAlgo = crypto.SHA3_256
	racingtime.KeyIndex = 0

	NFTId := uint64(2)

	_result := transactions.BurnRacingTimeNFT(NFTId)
	if _result.Error != nil {
		println(_result.Error.Error())
	}
	println(_result)
}
