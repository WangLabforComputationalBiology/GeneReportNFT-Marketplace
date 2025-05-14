package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/internal/models"
	"GeneReport_platform/pkg/appErrors"
)

var (
	CollectionServ *CollectionService
)

type CollectionService struct {
	iCollectionBase
}

// collection基础接口
type iCollectionBase interface {
}

func RegisterCollectionService() {
	CollectionServ = &CollectionService{}
}

// GetCollectionInfoByID 将[]dao.CollectionWithGNFT 整理为 dto.GetCollectionInfoResp
func (c *CollectionService) GetCollectionInfoByID(collectionID string) (dto.GetCollectionWithGNFTsByIDResp, error) {
	results, err := dao.GetCollectionDao().GetCollectionWithGNFTByID(collectionID)
	if err != nil {
		return dto.GetCollectionWithGNFTsByIDResp{}, appErrors.New(503, "获取集合信息失败", err)
	}
	collection := results[0].Collection
	var gnfts []models.GNFT
	for _, result := range results {
		gnfts = append(gnfts, result.GNFT)
	}
	return dto.GetCollectionWithGNFTsByIDResp{
		Collection: collection,
		GNFTs:      gnfts,
	}, nil
}

func (c *CollectionService) GetCollectionInfosByCreator(creator string) (dto.GetCollectionsWithGNFTsByCreatorResp, error) {
	//获取用户创建的集合信息
	targetCollections, err := dao.GetCollectionDao().GetCollectionsInfoByCreator(creator)
	if err != nil {
		return dto.GetCollectionsWithGNFTsByCreatorResp{}, appErrors.New(503, "获取集合信息失败", err)
	}

	//创建一个collectionID到index的映射，用于优化后续的遍历
	targetCollectionIdsMap := make(map[string]int)

	//初始化返回值并构建映射
	toResp := dto.GetCollectionsWithGNFTsByCreatorResp{
		Collections: make([]dto.GetCollectionWithGNFTsByIDResp, len(targetCollections)),
	}
	for i, collection := range targetCollections {
		toResp.Collections[i].Collection = collection
		targetCollectionIdsMap[collection.Id] = i
	}

	//获取联表查询结果
	results, err := dao.GetCollectionDao().GetCollectionWithGNFTByCreator(creator)
	if err != nil {
		return dto.GetCollectionsWithGNFTsByCreatorResp{}, appErrors.New(503, "获取集合信息失败", err)
	}

	//遍历
	for _, result := range results {
		index := targetCollectionIdsMap[result.Collection.Id]
		toResp.Collections[index].GNFTs = append(toResp.Collections[index].GNFTs, result.GNFT)
	}
	return toResp, nil
}
