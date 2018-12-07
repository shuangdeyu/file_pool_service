package conf

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"time"
)

// 配置项
type Configuration struct {
	DbUser         string        `yaml:"db_user"`
	DbPassword     string        `yaml:"db_password"`
	DbName         string        `yaml:"db_name"`
	DbPreFix       string        `yaml:"db_prefix"`
	RedisHost      string        `yaml:"redis_host"`     // redis地址
	RedisPassword  string        `yaml:"redis_password"` // redis密码
	RedisTimeout   time.Duration `yaml:"redis_timeout"`  // redis超时时间
	RedisDatabase  int           `yaml:"redis_database"` // redis库
	ServiceAddress string        `yaml:"service_address"`
	LocalLog       string        `yaml:"local_log"`
}

var AppConfig *Configuration

// 加载配置文件
func LoadConfig(path string) error {
	// 读取文件
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	// 解析文件数据，并存入结构体
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	// 将解析后的数据赋值给全局变量
	AppConfig = &config
	return err
}

// 获取配置
func GetConfig() *Configuration {
	return AppConfig
}
