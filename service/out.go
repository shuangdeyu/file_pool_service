package service

const (
	Success              = 100000
	ERROR_USER_NOT_EXITS = 100001 // 用户不存在
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
		if code == ERROR_USER_NOT_EXITS {
			return "用户不存在"
		}
	}
	return ""
}
