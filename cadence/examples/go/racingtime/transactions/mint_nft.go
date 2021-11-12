// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime"
)

var mintArt = fmt.Sprintf(`
import NonFungibleToken from %s
import RacingTime from %s
transaction(recipient: Address,typeID: UInt32, rewardID: UInt32,serialNumber: UInt32,ipfs: String) {
    let minter: &RacingTime.NFTMinter

    prepare(signer: AuthAccount) {
        self.minter = signer.borrow<&RacingTime.NFTMinter>(from: RacingTime.MinterStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
    }
    execute {
        let recipient = getAccount(recipient)
        let receiver = recipient
            .getCapability(RacingTime.CollectionPublicPath)!
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")
        self.minter.mintNFT(recipient: receiver, typeID: typeID, rewardID: rewardID,serialNumber: serialNumber,ipfs: ipfs )
    }
}`, racingtime.NonFungibleTokenAddress, racingtime.ContractOwnAddress)

type NFTInfo struct {
	RewardID     uint32
	TypeID       uint32
	SerialNumber uint32
	Ipfs         string
}

func MintNFT(nft *NFTInfo) *flow.TransactionResult {

	referenceBlock, err := racingtime.FlowClient.GetLatestBlock(racingtime.Ctx, false)
	if err != nil {
		panic(err)
	}

	acctAddress, acctKey, signer := lib.ServiceAccount(racingtime.FlowClient, racingtime.SigAlgo, racingtime.HashAlgo, racingtime.KeyIndex, racingtime.Address, racingtime.PrivateKey)

	tx := flow.NewTransaction().
		SetScript([]byte(mintArt)).
		SetGasLimit(100).
		SetProposalKey(acctAddress, acctKey.Index, acctKey.SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(acctAddress).
		AddAuthorizer(acctAddress)

	if err := tx.AddArgument(cadence.NewAddress(flow.HexToAddress(racingtime.Address))); err != nil {
		panic(err)
	}

	if err := tx.AddArgument(cadence.NewUInt32(nft.RewardID)); err != nil {
		panic(err)
	}

	if err := tx.AddArgument(cadence.NewUInt32(nft.TypeID)); err != nil {
		panic(err)
	}

	if err := tx.AddArgument(cadence.NewUInt32(nft.SerialNumber)); err != nil {
		panic(err)
	}
	_ipfs, err := cadence.NewString(nft.Ipfs)
	if err != nil {
		panic(err)
	}
	if err := tx.AddArgument(_ipfs); err != nil {
		panic(err)
	}

	if err := tx.SignEnvelope(acctAddress, acctKey.Index, signer); err != nil {
		panic(err)
	}

	if err := racingtime.FlowClient.SendTransaction(racingtime.Ctx, *tx); err != nil {
		panic(err)
	}

	return lib.WaitForSeal(racingtime.Ctx, racingtime.FlowClient, tx.ID())
}
