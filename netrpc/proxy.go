package netrpc

/*
 * rpc代理入口.
 */

import (
	"cherry/base"
)

/************************************************************************/
// 常量, 变量, 结构体, 接口定义.
/************************************************************************/

// Proxy TODO
type Proxy struct {
}

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// Write 写入数据.
func (owner *Proxy) Write(arg *RPCArgs, reply *RPCReply) error {
	base.LogDebug("proxy.go - Write:", arg, reply)
	reply.Data = arg.Data
	return nil
}

// Read 读取数据.
func (owner *Proxy) Read(arg *RPCArgs, reply *RPCReply) error {
	base.LogDebug("proxy.go - Read:", arg, reply)
	reply.Data = arg.Data
	return nil
}

/************************************************************************/
// 模块内功能实现.
/************************************************************************/

/************************************************************************/
// 模块内功能调试.
/************************************************************************/

/************************************************************************/
// 模块单元测试.
/************************************************************************/
