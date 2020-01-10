package netwebsocket

import (
	"cherry/base"
	"sync"

	"github.com/gorilla/websocket"
)

/*
 * 网络连接管理.
 * 管理Conn类型实例.
 */

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// ClientMngr TODO.
type ClientMngr struct {
	sync.RWMutex
	serv       *Service
	seed       int32
	clients    map[int32]*Client
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *Client
}

/************************************************************************/
// export functions.
/************************************************************************/

// SendToUID 通过用户UID发送
func (owner *ClientMngr) SendToUID(uid int32, msg *Msg) {
	buf := msgToBytes(msg)
	base.LogDebug("SendToUID: ", uid, len(buf))
}

// SendToConn 通过连接ID发送
func (owner *ClientMngr) SendToConn(ConnID int32, msg *Msg) {
	buf := msgToBytes(msg)
	for id, client := range owner.clients {
		if id == ConnID {
			client.send(buf)
		}
	}
}

// SendToAll TODO
func (owner *ClientMngr) SendToAll(msg *Msg) {
	buf := msgToBytes(msg)
	for _, client := range owner.clients {
		client.send(buf)
	}
}

// SendWithOut TODO
func (owner *ClientMngr) SendWithOut(msg *Msg, ignore int32) {
	buf := msgToBytes(msg)
	for id, client := range owner.clients {
		if id != ignore {
			client.send(buf)
		}
	}
}

/************************************************************************/
// moudule functions.
/************************************************************************/

// create new ClientMngr.
func newClientMngr(serv *Service) *ClientMngr {
	tcm := new(ClientMngr)
	tcm.serv = serv
	tcm.clients = make(map[int32]*Client)
	tcm.broadcast = make(chan []byte)
	tcm.register = make(chan *websocket.Conn)
	tcm.unregister = make(chan *Client)
	return tcm
}

// start TODO
func (owner *ClientMngr) start() {
	for {
		select {
		case conn := <-owner.register:
			owner.seedTick()
			id := owner.seed
			base.LogDebug("A new socket has connected. ", id, conn)
			owner.clients[id] = newClient(id, conn, owner.serv)
		case client := <-owner.unregister:
			if _, ok := owner.clients[client.id]; ok {
				base.LogDebug("A socket has disconnected. ", client.id)
				delete(owner.clients, client.id)
			}
		case message := <-owner.broadcast:
			base.LogDebug("broadcast message: ", message)
			// for id, client := range owner.clients {
			// 	select {
			// 	case conn.send <- message:
			// 	default:
			// 		close(conn.send)
			// 		delete(owner.clients, id)
			// 	}
			// }
		}
	}
}

// seed tick.
func (owner *ClientMngr) seedTick() {
	owner.seed++
}

/************************************************************************/
// unit tests.
/************************************************************************/
