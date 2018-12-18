package model

import (
	"fmt"
	"helper_go/comhelper"
	"log"
)

type User struct {
	Id         int    `xorm:"int(11) NOT NULL autoincr" json:"id"`
	Name       string `xorm:"varchar(32) NOT NULL" json:"name"`
	Mobile     string `xorm:"varchar(32) NOT NULL" json:"mobile"`
	Password   string `xorm:"varchar(128) NOT NULL" json:"password"`
	HeadPic    string `xorm:"varchar(255) DEFAULT NULL" json:"head_pic"`
	Job        string `xorm:"varchar(32) DEFAULT NULL" json:"job"`
	Desc       string `xorm:"text" json:"desc"`
	Email      string `xorm:"varchar(50) DEFAULT NULL" json:"email"`
	CreateTime string `xorm:"datetime NOT NULL DEFAULT '2001-01-01 00:00:00'" json:"create_time"`
}

var DefaultUser = &User{}

/**
 * 执行原生sql查询，返回string类型的map
 */
func (m *User) Query(args ...interface{}) ([]map[string]string, error) {
	// 基础sql语句
	sql := ""
	switch val := args[0].(type) {
	case string:
		sql = val
	}
	// 映射参数
	params := []interface{}{}
	if len(args) > 1 {
		switch val := args[1].(type) {
		case []interface{}:
			params = val
		}
	}
	// order 语句拼接
	if len(args) > 2 {
		switch val := args[2].(type) {
		case string:
			sql += " " + val
		}
	}
	// limit 语句拼接
	if len(args) > 3 {
		switch val := args[3].(type) {
		case []int:
			sql += " limit " + comhelper.IntToString(val[0]) + "," + comhelper.IntToString(val[1])
		}
	}

	ret, err := DbInit().SQL(sql, params...).QueryString()
	if err != nil {
		log.Print(err.Error())
	}
	return ret, nil
}

/**
 * 执行原生sql，返回定义的结构体类型
 */
func (m *User) QueryStructure(args ...interface{}) ([]User, error) {
	// 基础sql语句
	sql := ""
	switch val := args[0].(type) {
	case string:
		sql = val
	}
	// 映射参数
	params := []interface{}{}
	if len(args) > 1 {
		switch val := args[1].(type) {
		case []interface{}:
			params = val
		}
	}
	// order 语句拼接
	if len(args) > 2 {
		switch val := args[2].(type) {
		case string:
			sql += " " + val
		}
	}
	// limit 语句拼接
	if len(args) > 3 {
		switch val := args[3].(type) {
		case []int:
			sql += " limit " + comhelper.IntToString(val[0]) + "," + comhelper.IntToString(val[1])
		}
	}

	result := []User{}
	err := DbInit().SQL(sql, params...).Find(&result)
	if err != nil {
		log.Print(err.Error())
	}
	return result, nil
}

/**
 * 通过参数构造sql查询，返回string类型的map
 */
func (m *User) QueryByMap(args ...interface{}) ([]map[string]string, error) {
	sql := "select * from f_user where 1=1 "
	// 拼接where语句
	var params []interface{}
	switch val := args[0].(type) {
	case Arr:
		for k, v := range val {
			switch v_type := v.(type) {
			case string:
				if v_type == "" {
					continue
				}
			}
			sql += fmt.Sprintf("and f_user.%s = ? ", k)
			params = append(params, v)
		}
	}
	// 拼接order语句
	if len(args) > 1 {
		switch val := args[1].(type) {
		case string:
			if val != "" {
				sql += val + " "
			}
		}
	}
	// 拼接limit语句
	if len(args) > 2 {
		switch val := args[2].(type) {
		case []int:
			sql += "limit " + comhelper.IntToString(val[0]) + "," + comhelper.IntToString(val[1])
		}
	}

	ret, err := DbInit().SQL(sql, params...).QueryString()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return ret, nil
}

/**
 * 通过参数构造sql查询，返回定义的结构体类型
 */
func (m *User) QueryStructureByMap(args ...interface{}) ([]User, error) {
	result := []User{}

	sql := "select * from f_user where 1=1 "
	// 拼接where语句
	var params []interface{}
	switch val := args[0].(type) {
	case Arr:
		for k, v := range val {
			switch v_type := v.(type) {
			case string:
				if v_type == "" {
					continue
				}
			}
			sql += fmt.Sprintf("and f_user.%s = ? ", k)
			params = append(params, v)
		}
	}
	// 拼接order语句
	if len(args) > 1 {
		switch val := args[1].(type) {
		case string:
			if val != "" {
				sql += val + " "
			}
		}
	}
	// 拼接limit语句
	if len(args) > 2 {
		switch val := args[2].(type) {
		case []int:
			sql += "limit " + comhelper.IntToString(val[0]) + "," + comhelper.IntToString(val[1])
		}
	}

	err := DbInit().SQL(sql, params...).Find(&result)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}

/**
 * 获取count
 */
func (m *User) Count() (int64, error) {
	ret, err := DbInit().Count(m)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return ret, nil
}

/**
 * 删除，通过参数构造sql
 */
func (m *User) Delete(args Arr) error {
	sql := "delete from f_user where 1=1 "
	var params []interface{}
	for k, v := range args {
		sql += fmt.Sprintf("and f_user.%s = ? ", k)
		params = append(params, v)
	}
	_, err := DbInit().SQL(sql, params...).QueryString()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 删除，绑定结构体
 */
func (m *User) DeleteByStructure(id int) error {
	_, err := DbInit().Id(id).Delete(m)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 更新，通过参数构造sql
 */
func (m *User) Update(set, args Arr) error {
	sql := "update f_user set "
	var params []interface{}
	for k, v := range set {
		sql += fmt.Sprintf("f_user.%s = ?,", k)
		params = append(params, v)
	}

	sql = sql[:len(sql)-1]
	sql += " where 1=1 "
	for k, v := range args {
		sql += fmt.Sprintf("and f_user.%s = ? ", k)
		params = append(params, v)
	}

	_, err := DbInit().SQL(sql, params...).QueryString()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 更新，绑定结构体
 */
func (m *User) UpdateByStructure(args *User) error {
	_, err := DbInit().Update(m, args)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 新增，绑定结构体
 */
func (m *User) InsertByStructure(args ...string) error {
	_, err := DbInit().Omit(args...).Insert(m)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
