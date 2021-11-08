import RacingTime from "../../contracts/RacingTime.cdc"

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
}