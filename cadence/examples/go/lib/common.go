// SPDX-License-Identifier: Apache-2.0

package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
	"strconv"
	"time"
)

func WaitForSeal(ctx context.Context, c *client.Client, id flow.Identifier) *flow.TransactionResult {
	result, err := c.GetTransactionResult(ctx, id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Waiting for transaction %s to be sealed...\n", id)

	for result.Status != flow.TransactionStatusSealed {
		time.Sleep(time.Second)
		fmt.Print(".")
		result, err = c.GetTransactionResult(ctx, id)
		if err != nil {
			panic(err)
			return result
		}
		if result.Error != nil {
			return result
		}
	}

	fmt.Printf("%v", result)
	fmt.Printf("Transaction %s sealed\n", id)
	return result
}

func ServiceAccount(flowClient *client.Client, sigAlgo crypto.SignatureAlgorithm, hashAlgo crypto.HashAlgorithm, keyIndex int, address, privateKey string) (flow.Address, *flow.AccountKey, crypto.Signer) {
	_sigAlgo := crypto.StringToSignatureAlgorithm(sigAlgo.String())

	servicePrivateKeyHex := privateKey
	_privateKey, err := crypto.DecodePrivateKeyHex(_sigAlgo, servicePrivateKeyHex)
	if err != nil {
		panic(err)
	}

	_hashAlgo := crypto.StringToHashAlgorithm(hashAlgo.String())

	_address := flow.HexToAddress(address)

	_account, err := flowClient.GetAccount(context.Background(), _address)
	if err != nil {
		panic(err)
	}

	_accKey := _account.Keys[keyIndex]
	_signer := crypto.NewInMemorySigner(_privateKey, _hashAlgo)
	return _address, _accKey, _signer
}

func GetPublicKey(sigAlgo crypto.SignatureAlgorithm, privateKey string) crypto.PublicKey {
	_sigAlgo := crypto.StringToSignatureAlgorithm(sigAlgo.String())

	servicePrivateKeyHex := privateKey
	_privateKey, err := crypto.DecodePrivateKeyHex(_sigAlgo, servicePrivateKeyHex)
	if err != nil {
		panic(err)
	}
	return _privateKey.PublicKey()
}

func GetReferenceBlockId(flowClient *client.Client) flow.Identifier {
	block, err := flowClient.GetLatestBlock(context.Background(), false)
	if err != nil {
		panic(err)
	}

	return block.ID
}

func CadenceValueToJsonString(value cadence.Value) string {
	result := CadenceValueToInterface(value)
	json1, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(json1)
}

func CadenceValueToInterface(field cadence.Value) interface{} {
	switch field.(type) {
	case cadence.Dictionary:
		result := map[string]interface{}{}
		for _, item := range field.(cadence.Dictionary).Pairs {
			result[item.Key.String()] = CadenceValueToInterface(item.Value)
		}
		return result
	case cadence.Struct:
		result := map[string]interface{}{}
		subStructNames := field.(cadence.Struct).StructType.Fields
		for j, subField := range field.(cadence.Struct).Fields {
			result[subStructNames[j].Identifier] = CadenceValueToInterface(subField)
		}
		return result
	case cadence.Array:
		result := []interface{}{}
		for _, item := range field.(cadence.Array).Values {
			result = append(result, CadenceValueToInterface(item))
		}
		return result
	default:
		result, err := strconv.Unquote(field.String())
		if err != nil {
			return field.String()
		}
		return result
	}
}
