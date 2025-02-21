package services

type (
	IUserInfoService interface {
		GetUserInfo(ctx context.Context, id any) (string, error)
		EditUserInfo(ctx context.Context) error
	}
	IUserInfo interface {
		// Get 根据编号读取活动信息
		Get(ctx context.Context, id any) (out *entity.UserInfo, err error)
		// Gets 根据编号读取读取多条活动信息
		Gets(ctx context.Context, id any) (list []*entity.UserInfo, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.UserInfoListInput) (out []*entity.UserInfo, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserInfoListInput) (out *do.UserInfoListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.UserInfo) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.UserInfo) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// PassWordEdit 修改密码
		PassWordEdit(ctx context.Context, userId uint, userPassword string) (bool, error)
		// GetUserData 获取用户详细信息
		GetUserData(ctx context.Context, userId uint) (userInfoOutput *model.UserInfoOutput, err error)
		// AddTags 批量设置标签
		AddTags(ctx context.Context, userIds string, tagIds string) (bool, error)
		// AddVouchers 添加代金券
		AddVouchers(ctx context.Context, userIds []uint, activityId uint) error
		// GetList 获取用户信息列表
		GetList(ctx context.Context, in *do.UserInfoListInput) (out *model.UserInfoListOutput, err error)
		// EditUser 编辑用户
		EditUser(ctx context.Context, userInfo *model.UserInfo) (affected int64, err error)
		// RemoveUser 删除用户
		RemoveUser(ctx context.Context, userId uint) (res bool, err error)
	}
)
