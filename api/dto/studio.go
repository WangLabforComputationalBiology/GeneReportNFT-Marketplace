package dto

type CreateAllFromThirdPartyReq struct {
	ProfileId       string `json:"profileId" binding:"required"`
	GeneSharingName string `json:"gene_sharing_name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Tags            string `json:"tags"` // tags用分号分隔,ex:third party:wegene;...
}

type CreateAllFromThirdPartyResp struct {
	NewGeneSharingContractAddress string `json:"new_geneSharing_contract_address"`
	TransactionHash               string `json:"transaction_hash"`
}

type ProfileIdsResp struct {
	ProfileIds []string `json:"profile_ids"`
}
