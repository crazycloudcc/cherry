package netwebsocket

/*
 * netwebsocket service.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"fmt"
	"net/http"

	"github.com/crazycloudcc/cherry/base"

	"github.com/gorilla/websocket"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// Service struct.
type Service struct {
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
	serv.clientMngr = newClientMngr(serv)

	return serv
}

// Start netwebsocket service.
func (owner *Service) Start() {
	base.LogInfo("netwebsocket.Service Started with:", owner.conf.String())
	go owner.clientMngr.start()
	http.HandleFunc("/ws", owner.doAccept)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", owner.conf.Host, owner.conf.Port), nil)
}

// Close netwebsocket service.
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

func (owner *Service) doAccept(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}

	owner.clientMngr.register <- conn

	// go client.read()
	// go client.write()
}

/************************************************************************/
// unit tests.
/************************************************************************/
// func main() {
// 	fmt.Println("Starting application...")
// 	go manager.Start()
// 	http.HandleFunc("/ws", wsPage)
// 	http.ListenAndServe(":8070", nil)
// }
