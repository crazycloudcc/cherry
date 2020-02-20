package nettcp

/*
 * network tcp service.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"fmt"
	"net"

	"github.com/crazycloudcc/cherry/base"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// Service struct.
type Service struct {
	listener   net.Listener
	conf       Config
	router     *Router
	clientMngr *ClientMngr
}

/************************************************************************/
// export functions.
/************************************************************************/

// NewService 创建新的服务.
func NewService(conf Config) *Service {
	serv := new(Service)
	serv.conf = conf
	serv.router = newRouter(conf.Max)
	serv.clientMngr = newClientMngr(conf.Max)
	return serv
}

// SetConnCallback register connection close event callback.
func (owner *Service) SetConnCallback(cbAdd, cbDel ConnEventCallback) {
	owner.clientMngr.callbackAdd = cbAdd
	owner.clientMngr.callbackDel = cbDel
}

// Start network service.
func (owner *Service) Start() {
	owner.doListen()
	go owner.doAccept()
	base.LogInfo("nettcp.Service Started with:", owner.conf.String())
}

// Close network service.
func (owner *Service) Close() {
	// TODO.
}

// RegHandler register net message handler.
func (owner *Service) RegHandler(msgid uint32, handler interface{}) {
	owner.router.RegHandler(msgid, handler)
}

// GetRouter get router Instance.
func (owner *Service) GetRouter() *Router {
	return owner.router
}

// GetClientMngr get TcpConnMngr Instance.
func (owner *Service) GetClientMngr() *ClientMngr {
	return owner.clientMngr
}

/************************************************************************/
// moudule functions.
/************************************************************************/

func (owner *Service) doListen() {
	url := fmt.Sprintf("%s:%d", owner.conf.Host, owner.conf.Port)
	listener, err := net.Listen(owner.conf.Type, url)
	if err != nil {
		base.LogFatal("doListen net.Listen:", err)
		return
	}
	owner.listener = listener
}

// goroutine.
func (owner *Service) doAccept() {
	for {
		conn, err := owner.listener.Accept()
		if err != nil {
			base.LogWarn("doAccept owner.listener.Accept:", err)
			continue
		}

		owner.clientMngr.addConnection(newClient(conn, owner))
		base.LogDebug("New Client From:", conn.RemoteAddr().String())
	}
}

/************************************************************************/
// unit tests.
/************************************************************************/
