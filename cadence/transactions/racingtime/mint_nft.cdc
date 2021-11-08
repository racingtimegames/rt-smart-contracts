import NonFungibleToken from "../../contracts/lib/NonFungibleToken.cdc"
import RacingTime from "../../contracts/RacingTime.cdc"

// This transction uses the NFTMinter resource to mint a new NFT.
//
// It must be run with the account that has the minter resource
// stored at path /storage/NFTMinter.
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
}