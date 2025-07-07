package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/contracts/sharingPlatformContract"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/internal/models"
	"GeneReport_platform/pkg/appErrors"
	"context"
	"encoding/json"
	"errors"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"math/big"
	"net/http"
	"reflect"
	"time"
)

var (
	MetadataServ *MetadataService
)

type MetadataService struct {
	iMetadataBase
}

// Metadata基础接口
type iMetadataBase interface {
	GetMetadataOverviewByOwner(owner string) (dto.GetMetadataOverviewByOwnerResp, error)
}

func RegisterMetadataService() {
	MetadataServ = &MetadataService{}
}

/*fill your method here*/

func (m *MetadataService) GetMetadataOverviewByOwner(owner string) (dto.GetMetadataOverviewByOwnerResp, error) {
	targetMetadatas, err := dao.GetMetadataDao().GetMetadataOverviewByOwner(owner)
	if err != nil {
		return dto.GetMetadataOverviewByOwnerResp{}, appErrors.New(503, "获取Metadata概述信息失败", err)
	}
	// 映射转换dto
	var toResp dto.GetMetadataOverviewByOwnerResp
	_ = mapstructure.Decode(targetMetadatas, &toResp)
	return toResp, nil
}

func (m *MetadataService) GetAllMetadataOverview(page int) (dto.GetAllMetadataOverviewResp, error) {
	results, pageNum, err := dao.GetMetadataDao().GetAllMetadataOverview(page)
	if err != nil {
		return dto.GetAllMetadataOverviewResp{}, appErrors.New(503, "获取Metadata概览失败", err)
	}

	return dto.GetAllMetadataOverviewResp{
		MultiMetadata: results,
		PageNum:       pageNum,
	}, nil
}

func (m *MetadataService) GetMetadataDetailByDataHash(dataHash string) (map[string]interface{}, error) {
	metadata, err := dao.GetMetadataDao().GetMetadataOverviewByDataHash(dataHash)
	if err != nil {
		return nil, appErrors.New(503, "获取Metadata详细信息失败", err)
	}
	//检查当前metadata是否被隐藏
	if metadata.IsHidden == true {
		return nil, appErrors.New(http.StatusForbidden, "该表现型数据当前为隐藏状态")
	}

	results, err := m.GetDataImpl(metadata.ProfileID, metadata.Category)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetDataImpl 获取Metadata的详情数据
func (m *MetadataService) GetDataImpl(profileId, category string) (map[string]interface{}, error) {
	unique := dto.UniqueProfiles{}
	//在数据库查出profileId一样记录
	if profileId != "" {
		configs.DB.Where("profile_id = ?", profileId).Find(&unique)
	}
	if unique.Status == 0 {
		return nil, appErrors.New(http.StatusOK, "数据还没处理完成，请稍等!")
	}

	// 获取结构体类型
	t, ok := dto.GetStructType(category)
	if !ok {
		return nil, appErrors.New(http.StatusBadRequest, "无效的 category", nil)
	}

	// 创建切片用于查询结果
	sliceType := reflect.SliceOf(t)
	results := reflect.New(sliceType).Interface()

	// GORM 查询
	err := configs.DB.Model(reflect.New(t).Interface()).
		Where("profile_id = ?", profileId).
		Order("report_id desc").
		Find(results).Error
	if err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "查询失败", err)
	}

	// 检查结果是否为空
	resultValue := reflect.ValueOf(results).Elem()
	if resultValue.Len() == 0 {
		return map[string]interface{}{"data": []interface{}{}}, nil // 返回空数组
	}

	// 转为 JSON 并反序列化为 map[string]interface{}
	tempData, err := json.Marshal(resultValue.Interface()) // 序列化切片值而非指针
	if err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "序列化失败", err)
	}

	var resultArray []map[string]interface{}
	if err := json.Unmarshal(tempData, &resultArray); err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "反序列化失败", err)
	}

	// 包装为 map[string]interface{}
	return map[string]interface{}{"details": resultArray}, nil
}

//	func GenerateEncryptedZIP(data []models.GenoType, w io.Writer) ([]byte, *io.PipeReader, error) {
//		// 生成 AES 密钥和 IV
//		key := make([]byte, 32) // 256 位
//		if _, err := rand.Read(key); err != nil {
//			return nil, nil, err
//		}
//		iv := make([]byte, aes.BlockSize)
//		if _, err := rand.Read(iv); err != nil {
//			return nil, nil, err
//		}
//		block, err := aes.NewCipher(key)
//		if err != nil {
//			return nil, nil, err
//		}
//		stream := cipher.NewCTR(block, iv)
//
//		// 格式化 private.key 内容
//		keyContent := fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", key)
//
//		// 创建管道
//		pr, pw := io.Pipe()
//
//		// 异步生成和打包 ZIP
//		go func() {
//			defer pw.Close()
//			zw := zip.NewWriter(pw)
//
//			// 添加加密的 XLSX
//			w, err := zw.Create("data.xlsx")
//			if err != nil {
//				pw.CloseWithError(err)
//				return
//			}
//			ew := cipher.StreamWriter{S: stream, W: w}
//			if err := utils.GenerateXLSX(ew, data); err != nil {
//				pw.CloseWithError(err)
//				return
//			}
//
//			// 添加 private.key
//			keyW, err := zw.Create("private.key")
//			if err != nil {
//				pw.CloseWithError(err)
//				return
//			}
//			if _, err := keyW.Write([]byte(keyContent)); err != nil {
//				pw.CloseWithError(err)
//				return
//			}
//
//			if err := zw.Close(); err != nil {
//				pw.CloseWithError(err)
//			}
//		}()
//
//		return iv, pr, nil
//	}
func (m *MetadataService) GetGenoTypeZip(dataHash, viewer, pubKey string) (dto.GetGenoTypeZipResp, error) {
	// 根据 hash 取 metadata
	metadata, err := dao.GetMetadataDao().GetMetadataOverviewByDataHash(dataHash)
	if err != nil {
		return dto.GetGenoTypeZipResp{}, appErrors.New(503, "获取Metadata详细信息失败", err)
	}

	// 查看者为Metadata的拥有者，直接通过校验
	if metadata.Owner == viewer {
		shortURL, err := DownloadServ.GenerateDownloadLink(dataHash, viewer, pubKey)
		if err != nil {
			return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusInternalServerError, "短链接服务繁忙，请稍后再试", err)
		}
		return dto.GetGenoTypeZipResp{DownloadURL: shortURL, AccessStatus: true}, nil
	}

	// 查看者非Metadata拥有者
	// 检查是否通过机构认证
	user, err := dao.GetUserDao().GetUser(viewer)
	if err != nil {
		return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}
	if user.Email == "UNKNOWN" {
		return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusForbidden, "请先进行机构邮箱认证")
	}

	// 检查metadata当前是否可共享
	if metadata.IsSharable == false {
		return dto.GetGenoTypeZipResp{}, appErrors.New(403, "该基因型数据当前非共享", err)
	}

	//用户是否有viewAccess
	isHave, err := dao.GetActivityDao().IsViewAccessExist(dataHash, viewer)
	if err != nil {
		return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}
	if !isHave {
		return dto.GetGenoTypeZipResp{}, appErrors.NewWithData(http.StatusForbidden, "您的访问权限不存在", dto.GetGenoTypeZipResp{DownloadURL: "", AccessStatus: false})
	}

	//若有 检查用户的viewAccess状态
	activity, err := dao.GetActivityDao().GetLatestViewAccess(viewer, metadata.DataHash)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}
	// viewAccess状态为过期状态
	if errors.Is(err, gorm.ErrRecordNotFound) || activity.Expiry < time.Now().Unix() {
		return dto.GetGenoTypeZipResp{}, appErrors.NewWithData(http.StatusForbidden, "您的访问权限已过期", dto.GetGenoTypeZipResp{DownloadURL: "", AccessStatus: false})
	}

	//链上验证
	status, err := sharingPlatformContract.NewContractIns().VerifyViewAccess(&bind.CallOpts{Pending: false}, common.HexToHash(dataHash), common.HexToAddress(viewer))
	if err != nil {
		return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusInternalServerError, "链上服务繁忙，请稍后再试", err)
	}
	switch status.Int64() {
	case 0:
		//验证通过，发放短链接
		shortURL, err := DownloadServ.GenerateDownloadLink(dataHash, viewer, pubKey)
		if err != nil {
			return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusInternalServerError, "短链接服务繁忙，请稍后再试", err)
		}
		return dto.GetGenoTypeZipResp{DownloadURL: shortURL, AccessStatus: true}, nil
	case 1:
		return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusBadRequest, "该Metadata不存在", err)
	case 2:
		return dto.GetGenoTypeZipResp{}, appErrors.NewWithData(http.StatusBadRequest, "链上查看许可不存在", dto.GetGenoTypeZipResp{DownloadURL: "", AccessStatus: false}, err)
	case 3:
		return dto.GetGenoTypeZipResp{}, appErrors.NewWithData(http.StatusBadRequest, "您的链上查看许可已过期", dto.GetGenoTypeZipResp{DownloadURL: "", AccessStatus: false}, err)
	case 4:
		return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusBadRequest, "您的查看许可已被封禁", err)
	default:
		return dto.GetGenoTypeZipResp{}, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}

}

func (m *MetadataService) ObtainViewAccess(txHash, userAddress, pubKey string) (dto.NewViewAccessResp, error) {
	// 检查是否通过机构认证
	user, err := dao.GetUserDao().GetUser(userAddress)
	if err != nil {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}
	if user.Email == "UNKNOWN" {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusUnauthorized, "请先进行机构邮箱认证")
	}

	//获取交易回执
	receipt, err := sharingPlatformContract.NewChainClient().GetTransactionReceipt(context.Background(), common.HexToHash(txHash), true)
	if err != nil {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}
	//检查交易状态
	if receipt == nil || receipt.Status != 0 {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusBadRequest, "您的链上交易失败或不存在，请检查")
	}
	//检查回执的合约地址是否正确以及发起者是否为用户
	if receipt.To != sharingPlatformContract.PlatformContractAddressHex || receipt.From != userAddress {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusBadRequest, "恶意行为，多次进行该操作将会被封禁")
	}

	//解析交易回执的input
	methodName, params, err := sharingPlatformContract.DecodeInputData(receipt)
	if methodName != "obtainViewAccess" {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusBadRequest, "交易所调用的方法错误")
	}

	//根据解码参数进行类型断言
	dataHashBytes, ok := params["dataHash"].([32]byte)
	dataHash := "0x" + common.Bytes2Hex(dataHashBytes[:])
	if dataHash == "0x" || !ok {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusBadRequest, "交易参数错误", err)
	}

	// 根据 hash 取 metadata
	metadata, err := dao.GetMetadataDao().GetMetadataOverviewByDataHash(dataHash)
	if err != nil {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusServiceUnavailable, "获取Metadata详细信息失败", err)
	}
	// 检查metadata当前是否可共享
	if metadata.IsSharable == false || metadata.IsHidden == true {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusForbidden, "抱歉，该基因型数据当前非共享", err)
	}

	// 定义事件签名和哈希
	eventSignatures := map[string]string{
		"NewViewAccess":     "0x" + common.Bytes2Hex(crypto.Keccak256([]byte("NewViewAccess(address,bytes32,uint256)"))),
		"RenewalViewAccess": "0x" + common.Bytes2Hex(crypto.Keccak256([]byte("RenewalViewAccess(address,bytes32,uint256)"))),
	}

	var expiry *big.Int
	var eventName string
	// 解析回执中的事件日志
	for _, logEntry := range receipt.Logs {
		if len(logEntry.Topics) < 3 {
			continue
		}

		// 根据事件签名哈希解析事件类型
		eventHash := logEntry.Topics[0]

		switch eventHash {
		case eventSignatures["NewViewAccess"]:
			//新建查看许可事件
			eventName = "NewViewAccess"
		case eventSignatures["RenewalViewAccess"]:
			//续约查看许可事件
			eventName = "RenewalViewAccess"
		default:
			continue
		}

		// 解析事件数据
		if len(logEntry.Data[2:]) != 64 {
			return dto.NewViewAccessResp{}, appErrors.New(http.StatusInternalServerError, "解析事件数据失败，请检查")
		}
		expiry = new(big.Int)
		_, success := expiry.SetString(logEntry.Data[2:], 16)
		if !success {
			return dto.NewViewAccessResp{}, appErrors.New(http.StatusInternalServerError, "解析事件数据失败，请检查")
		}

		// 从 Topics 中提取参数
		viewer := "0x" + logEntry.Topics[1][26:]
		if userAddress != viewer {
			return dto.NewViewAccessResp{}, appErrors.New(http.StatusForbidden, "用户不匹配,多次进行该操作将会被封禁", err)
		}
		break
	}
	geneSharingAddress, ok := params["geneSharingAddress"].(common.Address)
	if !ok {
		return dto.NewViewAccessResp{}, appErrors.New(http.StatusInternalServerError, "解析事件数据失败，请检查")
	}
	//链下数据库
	activity := &models.Activity{
		Id:              uuid.New().String(),
		Metadata:        dataHash,
		TransactionHash: receipt.TransactionHash,
		UserAddress:     userAddress,
		Time:            time.Now(),
		Expiry:          expiry.Int64(),
		From:            geneSharingAddress.Hex(),
		To:              userAddress,
		GeneSharing:     metadata.GeneSharingContractAddress,
	}
	if eventName == "NewViewAccess" {
		err = dao.GetActivityDao().NewViewAccess(activity)
	} else {
		err = dao.GetActivityDao().RenewalViewAccess(activity)
	}
	if err != nil {
		return dto.NewViewAccessResp{}, appErrors.New(503, "服务繁忙，请稍后再试", err)
	}

	//生成短链接
	shortURL, err := DownloadServ.GenerateDownloadLink(dataHash, userAddress, pubKey)
	if err != nil {
		return dto.NewViewAccessResp{}, appErrors.New(503, "生成短链接失败", err)
	}
	return dto.NewViewAccessResp{DownloadURL: shortURL}, nil
}
