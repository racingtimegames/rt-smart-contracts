package transactions

import (
	"context"
	"fmt"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/templates"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/lib"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime"
	"io/ioutil"
)

func UpdateContract() *flow.TransactionResult {

	referenceBlock, err := racingtime.FlowClient.GetLatestBlock(context.Background(), false)
	if err != nil {
		panic(err)
	}

	serviceAcctAddr, serviceAcctKey, singer := lib.ServiceAccount(racingtime.FlowClient, racingtime.SigAlgo, racingtime.HashAlgo, racingtime.KeyIndex, racingtime.Address, racingtime.PrivateKey)

	contractPath := fmt.Sprintf("../../contracts/%s.cdc", racingtime.ContractsName)

	code, err := ioutil.ReadFile(contractPath)
	if err != nil {
		panic(err)
	}

	tx := templates.UpdateAccountContract(serviceAcctAddr, templates.Contract{
		Name:   racingtime.ContractsName,
		Source: string(code),
	})
	tx.SetProposalKey(
		serviceAcctAddr,
		serviceAcctKey.Index,
		serviceAcctKey.SequenceNumber,
	)
	tx.SetReferenceBlockID(referenceBlock.ID)
	tx.SetPayer(serviceAcctAddr)
	tx.SetGasLimit(9999)
	if err := tx.SignEnvelope(serviceAcctAddr, serviceAcctKey.Index, singer); err != nil {
		panic(err)
	}

	if err := racingtime.FlowClient.SendTransaction(racingtime.Ctx, *tx); err != nil {
		panic(err)
	}
	println(tx.ID().String())
	_rest := lib.WaitForSeal(racingtime.Ctx, racingtime.FlowClient, tx.ID())

	return _rest
}
