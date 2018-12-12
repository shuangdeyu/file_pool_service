package service

import (
	"file_pool_service/model"
	"helper_go/timehelper"
)

/**
 * 登录
 */
type LoginArgs struct {
	Name     string // 用户名
	Password string // 密码
}

func Login(params *LoginArgs) *Out {
	errorCode := SUCCESS
	// 判断用户名是否正确
	ret, _ := model.DefaultUser.QueryByMap(model.Arr{"name": params.Name})
	if !(len(ret) > 0) {
		errorCode = ERROR_USER_NAME_PASSWORD
	} else {
		// 验证密码是否正确
		info := ret[0]
		if params.Password != info["password"] {
			errorCode = ERROR_USER_NAME_PASSWORD
		}
	}
	// 登录失败错误次数记录，判断是否冻结
	if errorCode != SUCCESS {

		return NewOut(errorCode)
	}

	return NewOut(ret[0])
}

/**
 * 注册
 */
type RegisterArgs struct {
	Name     string // 用户名
	Password string // 密码
	Email    string // 邮箱
}

func Register(params *RegisterArgs) *Out {
	// 检查用户是否已经存在
	ret, _ := model.DefaultUser.QueryByMap(model.Arr{"name": params.Name})
	if len(ret) > 0 {
		return NewOut(ERROR_USER_NAME_ALERDAY_EXITS)
	}
	// 插入用户
	insert := &model.User{
		Name:       params.Name,
		Password:   params.Password,
		Email:      params.Email,
		CreateTime: timehelper.CurrentTime(),
	}
	err := insert.InsertByStructure()
	if err != nil {
		return NewOut(ERROR_INSERT_FAILED)
	}

	return NewOut(SUCCESS)
}

/**
 * 获取用户基本信息
 * @param UserId int 用户id
 */
type GetUserInfoArgs struct {
	UserId int // 用户id
}

func GetUserInfo(params *GetUserInfoArgs) *Out {
	if !(params.UserId > 0) {
		return NewOut(ERROR_USER_NOT_EXITS)
	}
	// 获取用户基本信息
	list, _ := model.DefaultUser.QueryByMap(model.Arr{"id": params.UserId})
	if len(list) > 0 {
		return NewOut(list[0])
	} else {
		return NewOut(ERROR_USER_NOT_EXITS)
	}
}

/**
 * 根据用户名获取用户信息
 */
type GetUserInfoByNameArgs struct {
	Name string // 用户名
}

func GetUserInfoByName(params *GetUserInfoByNameArgs) *Out {
	// 获取用户基本信息
	list, _ := model.DefaultUser.QueryByMap(model.Arr{"name": params.Name})
	if len(list) > 0 {
		return NewOut(list[0])
	} else {
		return NewOut(ERROR_USER_NOT_EXITS)
	}
}
