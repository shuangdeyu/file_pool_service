package service

import (
	"file_pool_service/model"
	"helper_go/timehelper"
)

/**
 * 获取用户文档池列表
 * @param UserId int 用户id
 * @param Offset int 偏移量
 * @param Limit int 请求数
 */
type GetUserPoolListArgs struct {
	UserId int
	Offset int
	Limit  int
}

func GetUserPoolList(params *GetUserPoolListArgs) *Out {
	if !(params.UserId > 0) {
		return NewOut(ERROR_USER_NOT_EXITS)
	}
	list, _ := model.DefaultPoolUser.Query(`SELECT
	pu.*,p.name,p.icon,p.manager_id,u.name manager,u.head_pic 
FROM
	f_pool_user pu
	INNER JOIN f_pool p ON pu.pool_id = p.id
	LEFT JOIN f_user u ON p.manager_id = u.id 
WHERE
	pu.user_id = ? 
	AND pu.delete_time IS NULL`, []interface{}{params.UserId})
	if len(list) > 0 {
		return NewOut(list)
	} else {
		return NewOut([]map[string]string{})
	}
}

/**
 * 获取用户回收站文档池列表
 * @param UserId int 用户id
 * @param Offset int 偏移量
 * @param Limit int 请求数
 */
type GetUserRecyclePoolListArgs struct {
	UserId int
	Offset int
	Limit  int
}

func GetUserRecyclePoolList(params *GetUserRecyclePoolListArgs) *Out {
	if !(params.UserId > 0) {
		return NewOut(ERROR_USER_NOT_EXITS)
	}
	list, _ := model.DefaultPoolUser.Query(`SELECT
	pu.*,p.name,p.icon
FROM
	f_pool_user pu
	INNER JOIN f_pool p ON pu.pool_id = p.id 
WHERE
	pu.user_id = ? 
	AND pu.is_manager = 'Y' 
	AND pu.delete_time IS NOT NULL`, []interface{}{params.UserId})
	if len(list) > 0 {
		return NewOut(list)
	} else {
		return NewOut([]map[string]string{})
	}
}

/**
 * 删除用户文档池
 */
type DeleteUserPoolByIdArgs struct {
	PoolUserId int
}

func DeleteUserPoolById(params *DeleteUserPoolByIdArgs) *Out {
	if !(params.PoolUserId > 0) {
		return NewOut(POOL_NOT_EXITS)
	}
	now := timehelper.CurrentTime()
	err := model.DefaultPoolUser.Update(model.Arr{"delete_time": now}, model.Arr{"id": params.PoolUserId})
	if err != nil {
		return NewOut(err)
	}
	return NewOut("")
}

/**
 * 还原用户文档池
 */
type RestoreUserPoolByIdArgs struct {
	PoolUserId int
}

func RestoreUserPoolById(params *RestoreUserPoolByIdArgs) *Out {
	if !(params.PoolUserId > 0) {
		return NewOut(POOL_NOT_EXITS)
	}
	err := model.DefaultPoolUser.Update(model.Arr{"delete_time": nil}, model.Arr{"id": params.PoolUserId})
	if err != nil {
		return NewOut(err)
	}
	return NewOut("")
}

/**
 * 根据用户文档池id获取文档池信息
 */
type GetPoolInfoByPoolUserIdArgs struct {
	PoolUserId int
}

func GetPoolInfoByPoolUserId(params *GetPoolInfoByPoolUserIdArgs) *Out {
	if !(params.PoolUserId > 0) {
		return NewOut(POOL_NOT_EXITS)
	}
	list, _ := model.DefaultPoolUser.Query(`SELECT
	pu.*,p.name,p.icon 
FROM
	f_pool_user pu
	INNER JOIN f_pool p ON pu.pool_id = p.id 
WHERE
	pu.id = ? 
LIMIT 1`, []interface{}{params.PoolUserId})
	if len(list) > 0 {
		return NewOut(list[0])
	} else {
		return NewOut([]map[string]string{})
	}
}
