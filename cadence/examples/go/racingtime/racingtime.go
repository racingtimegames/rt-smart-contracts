// SPDX-License-Identifier: Apache-2.0

package racingtime

import (
	"context"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
	"google.golang.org/grpc"
)

const (
	ContractsName = "RacingTime"
)

var (
	Ctx        = context.Background()
	FlowClient *client.Client
)

var (
	SigAlgo    crypto.SignatureAlgorithm
	HashAlgo   crypto.HashAlgorithm
	KeyIndex   int
	Address    string
	PrivateKey string
)

var (
	Node                    = "access.testnet.nodes.onflow.org:9000"
	NonFungibleTokenAddress = "0x631e88ae7f1d7c20"
	ContractOwnAddress      = "0xe0e251b47ff622ba"
	//FungibleTokenAddress    = "0x9a0766d93b6608b7"
	//FlowTokenAddress        = "0x7e60df042a9c0868"
)

func init() {
	var err error
	FlowClient, err = client.New(Node, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
}
