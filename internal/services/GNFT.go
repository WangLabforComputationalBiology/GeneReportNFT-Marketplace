package services

var (
	gnftService *GNFTService
)

type GNFTService struct {
	iGNFTBase interface{}
}

// GNFT基础接口
type iGNFTBase interface {
}

func RegisterGNFTService() {
	gnftService = &GNFTService{}
}

/*fill your method here*/
