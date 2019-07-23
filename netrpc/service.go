package netrpc

/*
 * 文件描述.
 */

import (
	"cherry/base"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/************************************************************************/
// 常量, 变量, 结构体, 接口定义.
/************************************************************************/

// Service TODO
type Service struct {
	listener net.Listener
	conn     net.Conn
	conf     Config // 配置文件.
	proxy    *Proxy // proxy.
}

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// NewService TODO
func NewService(conf Config) *Service {
	s := new(Service)
	s.conf = conf
	s.proxy = new(Proxy)
	return s
}

// Start TODO
func (owner *Service) Start() {
	url := fmt.Sprintf("%s:%d", owner.conf.Host, owner.conf.Port)
	l, err := net.Listen(owner.conf.Type, url)
	if err != nil {
		base.LogError("netrpc service.go - Start:", err)
		return
	}
	owner.listener = l

	regErr := rpc.Register(owner.proxy)
	if regErr != nil {
		base.LogError("netrpc service.go - Register proxy:", regErr)
		return
	}

	for {
		conn, err := owner.listener.Accept()
		if err != nil {
			base.LogError("netrpc service.go - Start Accept Error:", err)
			break
		}
		owner.conn = conn
		go jsonrpc.ServeConn(owner.conn)
	}
}

// Stop TODO
func (owner *Service) Stop() {
	owner.listener.Close()
	owner.conn.Close()
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
