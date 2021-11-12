// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"fmt"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime"
)

var setupAccountContracts = fmt.Sprintf(`
import NonFungibleToken from %s
import RacingTime from %s
transaction {
    prepare(signer: AuthAccount) {
        if signer.borrow<&RacingTime.Collection>(from: RacingTime.CollectionStoragePath) == nil {
            let collection <- RacingTime.createEmptyCollection()
            
            signer.save(<-collection, to: RacingTime.CollectionStoragePath)
            signer.link<&RacingTime.Collection{NonFungibleToken.CollectionPublic, RacingTime.RacingTimeCollectionPublic}>(RacingTime.CollectionPublicPath, target: RacingTime.CollectionStoragePath)
        }
    }
}`, racingtime.NonFungibleTokenAddress, racingtime.ContractOwnAddress)

func SetupAccount() *flow.TransactionResult {
	referenceBlock, err := racingtime.FlowClient.GetLatestBlock(racingtime.Ctx, false)
	if err != nil {
		panic(err)
	}

	acctAddress, acctKey, signer := lib.ServiceAccount(racingtime.FlowClient, racingtime.SigAlgo, racingtime.HashAlgo, racingtime.KeyIndex, racingtime.Address, racingtime.PrivateKey)

	tx := flow.NewTransaction().
		SetScript([]byte(setupAccountContracts)).
		SetGasLimit(100).
		SetProposalKey(acctAddress, acctKey.Index, acctKey.SequenceNumber).
		SetReferenceBlockID(referenceBlock.ID).
		SetPayer(acctAddress).
		AddAuthorizer(acctAddress)

	if err := tx.SignEnvelope(acctAddress, acctKey.Index, signer); err != nil {
		panic(err)
	}

	if err := racingtime.FlowClient.SendTransaction(racingtime.Ctx, *tx); err != nil {
		panic(err)
	}

	return lib.WaitForSeal(racingtime.Ctx, racingtime.FlowClient, tx.ID())
}
