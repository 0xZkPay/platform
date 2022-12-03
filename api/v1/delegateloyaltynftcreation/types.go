package delegateloyaltynftcreation

type DelegateLoyaltyNFTCreationRequest struct {
	CreatorAddress string `json:"creatorAddress" binding:"required,hexadecimal"`
	MetaDataHash   string `json:"metadataUrl" binding:"required"`
}

type DelegateLoyaltyNFTCreationPayload struct {
	TransactionHash string `json:"transactionHash"`
}
