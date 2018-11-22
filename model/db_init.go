package model

import (
	"file_pool_service/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
)

var Engine *xorm.Engine

type Arr map[string]interface{}

func DbInit() *xorm.Engine {
	if Engine == nil {
		db_conf := conf.GetConfig()
		// 连接数据库
		engine, err := xorm.NewEngine("mysql", db_conf.DbUser+":"+db_conf.DbPassword+"@/"+db_conf.DbName+"?charset=utf8")
		if err != nil {
			log.Println("数据库连接出错: ", err.Error())
		}
		// 设置表前缀映射
		if db_conf.DbPreFix != "" {
			tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, db_conf.DbPreFix)
			engine.SetTableMapper(tbMapper)
		}
		Engine = engine
	}
	return Engine
}
