package scripts

import (
	"context"
	"fmt"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/racingtimegames/rt-smart-contracts/cadence/examples/go/racingtime"
)

var getNftInfo string = fmt.Sprintf(`
import RacingTime from %s

pub struct RacingTimeNFTData {
	pub let id: UInt64
	pub let rewardID: UInt32
	pub let typeID: UInt32
	pub let serialNumber: UInt32
	pub let ipfs: String

	init(id: UInt64, rewardID: UInt32,initTypeID: UInt32, serialNumber: UInt32,ipfs: String) {
		self.id=id
		self.rewardID=rewardID
		self.typeID=initTypeID
		self.serialNumber=serialNumber
		self.ipfs=ipfs
	}
}

pub fun main(address:Address) : [RacingTimeNFTData] {
	var nfts: [RacingTimeNFTData] = []
    let account = getAccount(address)
	
	if let artCollection= account.getCapability(RacingTime.CollectionPublicPath).borrow<&{RacingTime.RacingTimeCollectionPublic}>()  {
		for id in artCollection.getIDs() {
			if let racingTimeNFT = artCollection.borrowRacingTime(id: id) {
				nfts.append(RacingTimeNFTData(id: id, rewardID: racingTimeNFT.data.rewardID, initTypeID: racingTimeNFT.data.typeID, serialNumber: racingTimeNFT.data.serialNumber, ipfs: racingTimeNFT.data.ipfs ))           
			}
		}
	}
	
    return nfts
}`, racingtime.ContractOwnAddress)

func AccountInfo(searchAddress string) cadence.Value {
	ctx := context.Background()
	result, err := racingtime.FlowClient.ExecuteScriptAtLatestBlock(ctx, []byte(getNftInfo), []cadence.Value{cadence.NewAddress(flow.HexToAddress(searchAddress))})
	if err != nil {
		panic(err)
	}

	return result
}
