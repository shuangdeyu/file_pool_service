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

/**
 * 获取用户文档池列表
 */
func (d *FilePoolService) GetUserPoolList(ctx context.Context, args *service.GetUserPoolListArgs, reply *ServiceReply) error {
	reply.Out = service.GetUserPoolList(args)
	return nil
}

/**
 * 获取用户回收站文档池列表(只显示自己创建的池)
 */
func (d *FilePoolService) GetUserRecyclePoolList(ctx context.Context, args *service.GetUserRecyclePoolListArgs, reply *ServiceReply) error {
	reply.Out = service.GetUserRecyclePoolList(args)
	return nil
}

/**
 * 删除文档池
 */
func (d *FilePoolService) DeleteUserPoolById(ctx context.Context, args *service.DeleteUserPoolByIdArgs, reply *ServiceReply) error {
	reply.Out = service.DeleteUserPoolById(args)
	return nil
}

/**
 * 恢复删除的文档池
 */
func (d *FilePoolService) RestoreUserPoolById(ctx context.Context, args *service.RestoreUserPoolByIdArgs, reply *ServiceReply) error {
	reply.Out = service.RestoreUserPoolById(args)
	return nil
}

/**
 * 根据用户文档池id获取文档池信息
 */
func (d *FilePoolService) GetPoolInfoByPoolUserId(ctx context.Context, args *service.GetPoolInfoByPoolUserIdArgs, reply *ServiceReply) error {
	reply.Out = service.GetPoolInfoByPoolUserId(args)
	return nil
}

/**
 * 获取文档池下的文档列表
 */
func (d *FilePoolService) GetFileListByPoolId(ctx context.Context, args *service.GetFileListByPoolIdArgs, reply *ServiceReply) error {
	reply.Out = service.GetFileListByPoolId(args)
	return nil
}

/**
 * 获取文档内容
 */

/**
 * 新增文档
 */

/**
 * 修改文档内容
 */

/**
 * 删除文档
 */

/**
 * 恢复删除的文档
 */
