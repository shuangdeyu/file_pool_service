package service

import (
	"file_pool_service/model"
)

/**
 * 获取文档列表
 */
type GetFileListArgs struct {
	Search   string
	Offset   int
	Limit    int
	PoolId   int
	Open     string
	IsDelete string
}

func GetFileList(params *GetFileListArgs) *Out {
	if params.Offset <= 0 {
		params.Offset = 0
	}
	if params.Limit <= 0 {
		params.Limit = 12
	}

	where := ""
	param := []interface{}{}
	param_count := []interface{}{}
	if params.Search != "" {
		where += " And name LIKE '%" + params.Search + "%'"
	}
	if params.PoolId > 0 {
		where += " AND pool_id = ?"
		param = append(param, params.PoolId)
		param_count = append(param_count, params.PoolId)
	}
	if params.Open != "" {
		where += " AND open = ?"
		param = append(param, params.Open)
		param_count = append(param_count, params.Open)
	}
	if params.IsDelete == "Y" {
		where += " AND delete_time IS NOT NULL"
	} else {
		where += " AND delete_time IS NULL"
	}
	param = append(param, params.Offset, params.Limit)

	list, _ := model.DefaultFile.Query("SELECT * FROM f_file WHERE 1=1 "+where+" ORDER BY edit_time DESC LIMIT ?,?", param)
	retCount, _ := model.DefaultFile.Query("SELECT count(id) as total_count FROM f_file WHERE 1=1 "+where, param_count)
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
 * 获取文档点赞数
 */
type GetFilePraiseCountArgs struct {
	FileId int
}

func GetFilePraiseCount(params *GetFilePraiseCountArgs) *Out {
	count := "0"
	ret, _ := model.DefaultFileCollect.Query("select count(id) as total_count from f_file_collect where praise = 'Y' and file_id = ?", []interface{}{params.FileId})
	if len(ret) > 0 {
		count = ret[0]["total_count"]
	}
	return NewOut(count)
}

/**
 * 获取文档收藏数
 */
func GetFileCollectCount(params *GetFilePraiseCountArgs) *Out {
	count := "0"
	ret, _ := model.DefaultFileCollect.Query("select count(id) as total_count from f_file_collect where collect = 'Y' and file_id = ?", []interface{}{params.FileId})
	if len(ret) > 0 {
		count = ret[0]["total_count"]
	}
	return NewOut(count)
}
