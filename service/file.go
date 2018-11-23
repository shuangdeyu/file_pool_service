package service

import "file_pool_service/model"

/**
 * 获取文档池下的文档列表
 */
type GetFileListByPoolIdArgs struct {
	PoolId int
}

func GetFileListByPoolId(params *GetFileListByPoolIdArgs) *Out {
	if !(params.PoolId > 0) {
		return NewOut(POOL_NOT_EXITS)
	}
	list, _ := model.DefaultFile.QueryByMap(model.Arr{"pool_id": params.PoolId})
	if len(list) > 0 {
		return NewOut(list)
	} else {
		return NewOut([]map[string]interface{}{})
	}
}
