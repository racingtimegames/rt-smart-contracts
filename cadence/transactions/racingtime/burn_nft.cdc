import NonFungibleToken from "../../contracts/lib/NonFungibleToken.cdc"
import RacingTime from "../../contracts/RacingTime.cdc"

transaction(burnID: UInt64) {
    prepare(signer: AuthAccount) {
		let collection <- signer.load<@RacingTime.Collection>(from: RacingTime.CollectionStoragePath)!

		let nft <- collection.withdraw(withdrawID: burnID)

		destroy nft

		signer.save(<-collection, to: RacingTime.CollectionStoragePath)
    }
}