package service

const (
	SUCCESS = 100000

	ERROR_INSERT_FAILED = 100001 // 新增失败
	ERROR_INVILD_PARAMS = 100002 // 参数错误

	ERROR_USER_NOT_EXITS          = 100101 // 用户不存在
	ERROR_USER_NAME_PASSWORD      = 100102 // 用户名或密码不正确
	ERROR_USER_NAME_ALERDAY_EXITS = 100103 // 用户名已被注册

	ERROR_POOL_NOT_EXITS = 100201 // 文档池不存在
)

type Out struct {
	OutData interface{}
	OutMsg  string
}

// 返回数据
func NewOut(outData interface{}) *Out {
	return &Out{outData, ToString(outData)}
}

// 错误信息匹配
func ToString(errorCode interface{}) string {
	switch code := errorCode.(type) {
	case error:
		if code != nil {
			return code.Error()
		}
	case int:
		if code == ERROR_INSERT_FAILED {
			return "新增失败"
		} else if code == ERROR_INVILD_PARAMS {
			return "参数不合法"
		} else if code == ERROR_USER_NOT_EXITS {
			return "用户不存在"
		} else if code == ERROR_USER_NAME_PASSWORD {
			return "用户名或密码不正确"
		} else if code == ERROR_USER_NAME_ALERDAY_EXITS {
			return "该用户名已被注册"
		} else if code == ERROR_POOL_NOT_EXITS {
			return "文档池不存在"
		}
	}
	return ""
}
