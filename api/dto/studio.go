package dto

type CreateAllFromThirdPartyReq struct {
	Profile         Profile // 第三方文件信息
	GeneSharingName string  `json:"gene_sharing_name"`
	Description     string  `json:"description"`
	Tags            string  `json:"tags"` // tags用分号分隔,ex:third party:wegene;...
}

type CreateAllFromThirdPartyResp struct {
	NewGeneSharingContractAddress string `json:"new_geneSharing_contract_address"`
	TransactionHash               string `json:"transaction_hash"`
}

type ProfileIdsResp struct {
	ProfileIds []string `json:"profile_ids"`
}
