// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime"
)

var (
	transferNft = fmt.Sprintf(`
	import NonFungibleToken from %s
	import RacingTime from %s

	transaction(recipient: Address, withdrawID: UInt64) {
    prepare(signer: AuthAccount) {
        let recipient = getAccount(recipient)
        let collectionRef = signer.borrow<&RacingTime.Collection>(from: RacingTime.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the owner's collection")
        let depositRef = recipient.getCapability(RacingTime.CollectionPublicPath)!.borrow<&{NonFungibleToken.CollectionPublic}>()!
        let nft <- collectionRef.withdraw(withdrawID: withdrawID)
        depositRef.deposit(token: <-nft)
    }
}`, racingtime.NonFungibleTokenAddress, racingtime.ContractOwnAddress)
)

// TransferRacingTimeNFT This transaction transfers a RacingTime NFT from one account to another.
func TransferRacingTimeNFT(toAddress string, nftId uint64) *flow.TransactionResult {
	referenceBlock, err := racingtime.FlowClient.GetLatestBlock(racingtime.Ctx, false)
	if err != nil {
		panic(err)
	}

	acctAddress, acctKey, signer := lib.ServiceAccount(racingtime.FlowClient, racingtime.SigAlgo, racingtime.HashAlgo, racingtime.KeyIndex, racingtime.Address, racingtime.PrivateKey)

	tx := flow.NewTransaction().
		SetScript([]byte(transferNft)).
		SetGasLimit(100).
		SetProposalKey(acctAddress, acctKey.Index, acctKey.SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(acctAddress).
		AddAuthorizer(acctAddress)

	if err := tx.AddArgument(cadence.NewAddress(flow.HexToAddress(toAddress))); err != nil {
		panic(err)
	}

	if err := tx.AddArgument(cadence.NewUInt64(nftId)); err != nil {
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
