package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/contracts/sharingPlatformContract"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appErrors"
	"github.com/ethereum/go-ethereum/common"
	"net/http"
)

var (
	StudioServ *StudioService
)

type StudioService struct {
	iStudioBase
}

func RegisterStudioService() {
	StudioServ = &StudioService{}
}

// 用户基础接口
type iStudioBase interface {
}

// CreateAllFromThirdPartyOnChain 从第三方平台创建（链上操作）
func (s *StudioService) CreateAllFromThirdPartyOnChain(userAddressHex string, req dto.CreateAllFromThirdPartyReq) (dto.CreateAllFromThirdPartyResp, error) {
	//验证对应的Metadata数据是否已经存在，geneSharing_contract_address，并根据profileId读取dataHash数组
	results, err := dao.GetMetadataDao().GetMetadataDetailByProfileId(req.ProfileId)
	if err != nil {
		return dto.CreateAllFromThirdPartyResp{}, appErrors.New(http.StatusInternalServerError, "请重试", err)
	}

	var dataHashBytes32s [][32]byte
	//验证预构建的Metadata的合法性
	for _, result := range results {
		// 十六进制字符串转字节存入32字节切片
		dataHashBytes32s = append(dataHashBytes32s, [32]byte(common.Hex2Bytes(result.DataHash)))
	}

	//链上数据创建
	newGeneSharingAddress, _, receipt, err := sharingPlatformContract.GetContractIns().CreateAllFromThirdParty(sharingPlatformContract.NewAdminTransactor(), common.HexToAddress(userAddressHex), dataHashBytes32s)
	// 由于某些失败（例如合约逻辑错误）可能在receipt中反映，而非 error，因此需要检查receipt状态,且fisco的receipt.Status为0表示成功
	if err != nil || receipt.Status != 0 {
		return dto.CreateAllFromThirdPartyResp{}, appErrors.New(http.StatusInternalServerError, "链上交易失败", err)
	}

	//var dataHashBytes32s [][32]byte
	//var tempDataHashBytes32 [32]byte
	//userAddressHex = "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"
	//copy(tempDataHashBytes32[:], common.Hex2Bytes("7E6cd50ffe51247535bfb9140f45f6d4df89c5A793d29a8f02bb3a913a707679"))
	//dataHashBytes32s = append(dataHashBytes32s, tempDataHashBytes32)
	//
	//newGeneSharingAddress, _, receipt, err := sharingPlatformContract.GetContractIns().CreateAllFromThirdParty(sharingPlatformContract.NewAdminTransactor(), common.HexToAddress(userAddressHex), dataHashBytes32s)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(newGeneSharingAddress.Hex())
	//fmt.Println(receipt)

	return dto.CreateAllFromThirdPartyResp{NewGeneSharingContractAddress: newGeneSharingAddress.Hex(), TransactionHash: receipt.TransactionHash}, nil
}

// CreateAllFromThirdPartyOffChain fixme CreateAllFromThirdPartyOffChain 从第三方平台创建（链下操作）
func (s *StudioService) CreateAllFromThirdPartyOffChain(userAddress string, req dto.CreateAllFromThirdPartyReq) error {
	return nil
}

func (s *StudioService) GetProfileIdsByUser(userAddress string) (dto.ProfileIdsResp, error) {
	results, err := dao.GetUserDao().GetProfileIdsByUser(userAddress)
	if err != nil {
		return dto.ProfileIdsResp{}, appErrors.New(http.StatusInternalServerError, "获取profile ids失败", err)
	}
	return dto.ProfileIdsResp{ProfileIds: results}, nil
}
