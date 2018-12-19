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
		where += " and pu.is_manager = 'Y' and pu.delete_time is null and p.delete_time is null"
	case POOL_TYPE_JOIN: // 我参与的
		where += " and pu.is_manager = 'N' and pu.delete_time is null and p.delete_time is null"
	case POOL_TYPE_DELETE: // 回收站
		where += " and pu.is_manager = 'Y' and pu.delete_time is not null"
	default:
		where += " and pu.delete_time is null and p.delete_time is null"
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
		return NewOut(map[string]interface{}{
			"count": count,
			"list":  list,
		})
	} else {
		return NewOut(map[string]interface{}{
			"count": "0",
			"list":  []map[string]string{},
		})
	}
}

/**
 * 删除用户文档池
 */
type DeleteUserPoolByIdArgs struct {
	PoolUserId int
	PoolId     int
	Manager    bool
}

func DeleteUserPoolById(params *DeleteUserPoolByIdArgs) *Out {
	if !(params.PoolUserId > 0) {
		return NewOut(ERROR_POOL_NOT_EXITS)
	}
	now := timehelper.CurrentTime()
	// 删除文档池用户
	err := model.DefaultPoolUser.Update(model.Arr{"delete_time": now}, model.Arr{"id": params.PoolUserId})
	if params.Manager {
		// 删除文档池
		err = model.DefaultPool.Update(model.Arr{"delete_time": now}, model.Arr{"id": params.PoolId})
	}
	if err != nil {
		return NewOut(err)
	}
	return NewOut(SUCCESS)
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
	return NewOut(SUCCESS)
}

/**
 * 获取池信息
 */
type GetPoolInfoArgs struct {
	PoolId int
}

func GetPoolInfo(params *GetPoolInfoArgs) *Out {
	if !(params.PoolId > 0) {
		return NewOut(ERROR_POOL_NOT_EXITS)
	}

	list, _ := model.DefaultPool.QueryByMap(model.Arr{"id": params.PoolId})
	if len(list) > 0 {
		return NewOut(list[0])
	} else {
		return NewOut(ERROR_POOL_NOT_EXITS)
	}
}

/**
 * 根据用户文档池id获取文档池信息
 */
type GetPoolInfoByIdArgs struct {
	PoolUserId int
}

func GetPoolInfoById(params *GetPoolInfoByIdArgs) *Out {
	if !(params.PoolUserId > 0) {
		return NewOut(ERROR_POOL_NOT_EXITS)
	}
	list, _ := model.DefaultPoolUser.Query(`SELECT
	pu.*,p.name,p.icon,p.desc,p.permit 
FROM
	f_pool_user pu
	INNER JOIN f_pool p ON pu.pool_id = p.id 
WHERE
	pu.id = ? 
LIMIT 1`, []interface{}{params.PoolUserId})
	if len(list) > 0 {
		return NewOut(list[0])
	} else {
		return NewOut(ERROR_POOL_NOT_EXITS)
	}
}

/**
 * 新建池
 */
type CreateNewPoolArgs struct {
	Name   string // 池名称
	Desc   string // 描述
	Icon   string // icon 地址
	Permit string // 权限值
	UserId int    // 用户id
}

func CreateNewPool(params *CreateNewPoolArgs) *Out {
	now := timehelper.CurrentTime()

	// 新增池信息
	insert_p := &model.Pool{
		Name:       params.Name,
		ManagerId:  params.UserId,
		Desc:       params.Desc,
		Icon:       params.Icon,
		CreateTime: now,
		Permit:     params.Permit,
	}
	err := insert_p.InsertByStructure("delete_time")
	if err != nil {
		return NewOut(ERROR_INSERT_FAILED)
	}
	pool_id := insert_p.Id
	if !(pool_id > 0) {
		return NewOut(ERROR_INSERT_FAILED)
	}
	// 新增用户池信息
	insert_pu := &model.PoolUser{
		UserId:     params.UserId,
		PoolId:     pool_id,
		IsManager:  "Y",
		CreateTime: now,
	}
	err = insert_pu.InsertByStructure("delete_time")
	if err != nil {
		return NewOut(ERROR_INSERT_FAILED)
	}

	return NewOut(SUCCESS)
}

/**
 * 修改池信息
 */
type EditPoolInfoArgs struct {
	Id     int
	Permit string
}

func EditPoolInfo(params *EditPoolInfoArgs) *Out {
	if !(params.Id > 0) {
		return NewOut(ERROR_POOL_NOT_EXITS)
	}
	err := model.DefaultPool.Update(model.Arr{"permit": params.Permit}, model.Arr{"id": params.Id})
	if err != nil {
		return NewOut(err)
	}
	return NewOut(SUCCESS)
}

/**
 * 获取池成员列表
 */
type GetPoolMembersArgs struct {
	PoolId int
}

func GetPoolMembers(params *GetPoolMembersArgs) *Out {
	if !(params.PoolId > 0) {
		return NewOut(ERROR_POOL_NOT_EXITS)
	}
	list, _ := model.DefaultPoolUser.Query(`select 
	pu.*,u.name user_name,u.head_pic 
from 
	f_pool_user pu 
	join f_user u on pu.user_id = u.id 
where 
	pu.pool_id = ? 
	and pu.delete_time is null`, []interface{}{params.PoolId})
	if len(list) > 0 {
		return NewOut(list)
	} else {
		return NewOut([]map[string]interface{}{})
	}
}

/**
 * 添加池成员
 */
type AddPoolMembersArgs struct {
	UserId    int
	PoolId    int
	IsManager string
}

func AddPoolMembers(params *AddPoolMembersArgs) *Out {
	if params.UserId < 1 || params.PoolId < 1 {
		return NewOut(ERROR_INVILD_PARAMS)
	}
	if params.IsManager != "Y" && params.IsManager != "N" {
		params.IsManager = "N"
	}
	now := timehelper.CurrentTime()
	insert := &model.PoolUser{
		UserId:     params.UserId,
		PoolId:     params.PoolId,
		IsManager:  params.IsManager,
		CreateTime: now,
	}
	err := insert.InsertByStructure("delete_time")
	if err != nil {
		return NewOut(err)
	}
	return NewOut(insert.Id)
}

/**
 * 删除池成员
 */
type DeletePoolMembersArgs struct {
	PoolId int
	UserId int
}

func DeletePoolMembers(params *DeletePoolMembersArgs) *Out {
	if params.UserId < 1 || params.PoolId < 1 {
		return NewOut(ERROR_INVILD_PARAMS)
	}
	now := timehelper.CurrentTime()
	err := model.DefaultPoolUser.Update(model.Arr{"delete_time": now}, model.Arr{"user_id": params.UserId, "pool_id": params.PoolId})
	if err != nil {
		return NewOut(err)
	}
	return NewOut(SUCCESS)
}
