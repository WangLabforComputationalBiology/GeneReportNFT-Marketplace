package services

var (
	orderService *OrderService
)

type OrderService struct {
	iUserBase
}

// 订单基础接口
type iOrderBase interface {
}

func RegisterOrderService() {
	orderService = &OrderService{}
}

/*fill your method here*/
