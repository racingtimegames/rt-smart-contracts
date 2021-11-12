// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"encoding/hex"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime"
)

func CreateAddress(seed string, sigAlgo crypto.SignatureAlgorithm, hashAlgo crypto.HashAlgorithm) (accountAddr string, privateKey string) {

	privateKey1, err := crypto.GeneratePrivateKey(sigAlgo, []byte(seed))
	if err != nil {
		panic(err)
	}

	privateKey = hex.EncodeToString(privateKey1.Encode())

	myAcctKey := flow.NewAccountKey().
		SetPublicKey(privateKey1.PublicKey()).
		SetSigAlgo(privateKey1.Algorithm()).
		SetHashAlgo(hashAlgo).
		SetWeight(flow.AccountKeyWeightThreshold)

	serviceAcctAddr, serviceAcctKey, serviceSigner := lib.ServiceAccount(racingtime.FlowClient, racingtime.SigAlgo, racingtime.HashAlgo, racingtime.KeyIndex, racingtime.Address, racingtime.PrivateKey)
	referenceBlockID := lib.GetReferenceBlockId(racingtime.FlowClient)
	createAccountTx := templates.CreateAccount([]*flow.AccountKey{myAcctKey}, nil, serviceAcctAddr)
	createAccountTx.SetProposalKey(
		serviceAcctAddr,
		serviceAcctKey.Index,
		serviceAcctKey.SequenceNumber,
	)
	createAccountTx.SetReferenceBlockID(referenceBlockID)
	createAccountTx.SetPayer(serviceAcctAddr)

	if err := createAccountTx.SignEnvelope(serviceAcctAddr, serviceAcctKey.Index, serviceSigner); err != nil {
		panic(err)
	}

	if err := racingtime.FlowClient.SendTransaction(racingtime.Ctx, *createAccountTx); err != nil {
		panic(err)
	}

	accountCreationTxRes := lib.WaitForSeal(racingtime.Ctx, racingtime.FlowClient, createAccountTx.ID())
	var myAddress flow.Address
	for _, event := range accountCreationTxRes.Events {
		if event.Type == flow.EventAccountCreated {
			accountCreatedEvent := flow.AccountCreatedEvent(event)
			myAddress = accountCreatedEvent.Address()
		}
	}

	accountAddr = myAddress.Hex()
	return
}
