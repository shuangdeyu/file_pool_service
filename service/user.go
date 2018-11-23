package service

import "file_pool_service/model"

/**
 * 获取用户基本信息
 * @param UserId int 用户id
 */
type GetUserInfoArgs struct {
	UserId int
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
