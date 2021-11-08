import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import RacingTime from "../../contracts/RacingTime.cdc"

transaction(recipient: Address, withdrawID: UInt64) {

    prepare(signer: AuthAccount) {

        let recipient = getAccount(recipient)

        let collectionRef = signer.borrow<&RacingTime.Collection>(from: RacingTime.CollectionStoragePath)
            ?? panic("Could not borrow a reference to the owner's collection")

        let depositRef = recipient.getCapability(RacingTime.CollectionPublicPath)!.borrow<&{NonFungibleToken.CollectionPublic}>()!

        let nft <- collectionRef.withdraw(withdrawID: withdrawID)

        depositRef.deposit(token: <-nft)
    }
}