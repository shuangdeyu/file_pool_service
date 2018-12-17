package service

import (
	"file_pool_service/model"
	"helper_go/timehelper"
)

const (
	POOL_TYPE_ALL    = "ALL"
	POOL_TYPE_MINE   = "MINE"
	POOL_TYPE_JOIN   = "JOIN"
	POOL_TYPE_DELETE = "DELETE"
)

/**
 * 获取用户文档池列表
 * @param UserId int 用户id
 * @param Offset int 偏移量
 * @param Limit int 请求数
 */
type GetUserPoolListArgs struct {
	UserId int
	Search string
	Offset int
	Limit  int
	Type   string
}

func GetUserPoolList(params *GetUserPoolListArgs) *Out {
	if !(params.UserId > 0) {
		return NewOut(ERROR_USER_NOT_EXITS)
	}

	if params.Offset <= 0 {
		params.Offset = 0
	}
	if params.Limit <= 0 {
		params.Limit = 12
	}

	where := "pu.user_id = ? "
	param := []interface{}{params.UserId}
	if params.Search != "" {
		where += " and p.name like '%" + params.Search + "%'"
	}
	switch params.Type {
	case POOL_TYPE_MINE: // 我创建的
		where += " and pu.is_manager = 'Y' and pu.delete_time is null"
	case POOL_TYPE_JOIN: // 我参与的
		where += " and pu.is_manager = 'N' and pu.delete_time is null"
	case POOL_TYPE_DELETE: // 回收站
		where += " and pu.is_manager = 'Y' and pu.delete_time is not null"
	default:
		where += " and pu.delete_time is null"
	}
	param = append(param, params.Offset, params.Limit)

	list, _ := model.DefaultPoolUser.Query(`select
	pu.*,p.name,p.icon,p.manager_id 
from
	f_pool_user pu
	join f_pool p on pu.pool_id = p.id 
where `+where+" order by pu.create_time desc limit ?,?", param)
	retCount, _ := model.DefaultPoolUser.Query(`select
	count(pu.id) total_count 
from
	f_pool_user pu 
	join f_pool p on pu.pool_id = p.id 
where `+where, []interface{}{params.UserId})
	count := "0"
	if len(retCount) > 0 {
		count = retCount[0]["total_count"]
	}

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
		return NewOut(ERROR_POOL_NOT_EXITS)
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
		return NewOut(ERROR_POOL_NOT_EXITS)
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
		return NewOut(ERROR_POOL_NOT_EXITS)
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
