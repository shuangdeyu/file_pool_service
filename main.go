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
 * 登录
 */
func (d *FilePoolService) Login(ctx context.Context, args *service.LoginArgs, reply *ServiceReply) error {
	reply.Out = service.Login(args)
	return nil
}

/**
 * 注册
 */
func (d *FilePoolService) Register(ctx context.Context, args *service.RegisterArgs, reply *ServiceReply) error {
	reply.Out = service.Register(args)
	return nil
}

/**
 * 获取用户信息
 */
func (d *FilePoolService) GetUserInfo(ctx context.Context, args *service.GetUserInfoArgs, reply *ServiceReply) error {
	reply.Out = service.GetUserInfo(args)
	return nil
}

/**
 * 根据用户名获取用户信息
 */
func (d *FilePoolService) GetUserInfoByName(ctx context.Context, args *service.GetUserInfoByNameArgs, reply *ServiceReply) error {
	reply.Out = service.GetUserInfoByName(args)
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
 * 获取文档池信息
 */
func (d *FilePoolService) GetPoolInfo(ctx context.Context, args *service.GetPoolInfoArgs, reply *ServiceReply) error {
	reply.Out = service.GetPoolInfo(args)
	return nil
}

/**
 * 根据用户文档池id获取文档池信息
 */
func (d *FilePoolService) GetPoolInfoById(ctx context.Context, args *service.GetPoolInfoByIdArgs, reply *ServiceReply) error {
	reply.Out = service.GetPoolInfoById(args)
	return nil
}

/**
 * 新建池
 */
func (d *FilePoolService) CreateNewPool(ctx context.Context, args *service.CreateNewPoolArgs, reply *ServiceReply) error {
	reply.Out = service.CreateNewPool(args)
	return nil
}

/**
 * 编辑池信息
 */
func (d *FilePoolService) EditPoolInfo(ctx context.Context, args *service.EditPoolInfoArgs, reply *ServiceReply) error {
	reply.Out = service.EditPoolInfo(args)
	return nil
}

/**
 * 获取池成员列表
 */
func (d *FilePoolService) GetPoolMembers(ctx context.Context, args *service.GetPoolMembersArgs, reply *ServiceReply) error {
	reply.Out = service.GetPoolMembers(args)
	return nil
}

/**
 * 添加池成员列表
 */
func (d *FilePoolService) AddPoolMembers(ctx context.Context, args *service.AddPoolMembersArgs, reply *ServiceReply) error {
	reply.Out = service.AddPoolMembers(args)
	return nil
}

/**
 * 删除池成员
 */
func (d *FilePoolService) DeletePoolMembers(ctx context.Context, args *service.DeletePoolMembersArgs, reply *ServiceReply) error {
	reply.Out = service.DeletePoolMembers(args)
	return nil
}

/**
 * 获取文档列表
 */
func (d *FilePoolService) GetFileList(ctx context.Context, args *service.GetFileListArgs, reply *ServiceReply) error {
	reply.Out = service.GetFileList(args)
	return nil
}

/**
 * 获取文档点赞数
 */
func (d *FilePoolService) GetFilePraiseCount(ctx context.Context, args *service.GetFilePraiseCountArgs, reply *ServiceReply) error {
	reply.Out = service.GetFilePraiseCount(args)
	return nil
}

/**
 * 获取文档收藏数
 */
func (d *FilePoolService) GetFileCollectCount(ctx context.Context, args *service.GetFilePraiseCountArgs, reply *ServiceReply) error {
	reply.Out = service.GetFileCollectCount(args)
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
