package main

import (
	"context"
	"file_pool_service/conf"
	"file_pool_service/service"
	"flag"
	"github.com/smallnest/rpcx/server"
	"log"
	"runtime"
)

type FilePoolService string

type ServiceReply struct {
	Out *service.Out `msg:"Out"`
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configPath := flag.String("c", "conf/conf.yaml", "config file path") // 系统配置文件
	flag.Parse()

	// 加载系统配置文件
	if err := conf.LoadConfig(*configPath); err != nil {
		log.Println("加载系统配置出错！")
	}

	// 注意地址，127.0.0.1:6666只能本地访问，服务器上0.0.0.0:6666允许远程访问
	addr := flag.String("addr", conf.GetConfig().ServiceAddress, "server address") // rpc服务地址

	// 启动服务
	s := server.NewServer()
	s.RegisterName("FilePoolService", new(FilePoolService), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

// 服务接口注册

/**
 * 获取用户信息
 */
func (d *FilePoolService) GetUserInfo(ctx context.Context, args *service.GetUserInfoArgs, reply *ServiceReply) error {
	reply.Out = service.GetUserInfo(args)
	return nil
}
