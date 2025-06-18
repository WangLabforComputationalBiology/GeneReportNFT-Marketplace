package models

type GeneSharingMetadata struct {
	GeneSharingContractAddress string   `json:"geneSharing_contract_address"`
	MetadataHash               [32]byte `json:"metadata_hash"`
}
