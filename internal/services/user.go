package services

var (
	userService *UserService
)

type UserService struct {
	iUserBase interface{}
}

// 用户基础接口
type iUserBase interface {
}

func RegisterUserService() {
	userService = &UserService{}
}

/*fill your method here*/
