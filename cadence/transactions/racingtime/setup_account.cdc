import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import RacingTime from "../../contracts/RacingTime.cdc"

transaction {
    prepare(signer: AuthAccount) {
        if signer.borrow<&RacingTime.Collection>(from: RacingTime.CollectionStoragePath) == nil {
            let collection <- RacingTime.createEmptyCollection()

            signer.save(<-collection, to: RacingTime.CollectionStoragePath)
            signer.link<&RacingTime.Collection{NonFungibleToken.CollectionPublic, RacingTime.RacingTimeCollectionPublic}>(RacingTime.CollectionPublicPath, target: RacingTime.CollectionStoragePath)
        }
    }
}