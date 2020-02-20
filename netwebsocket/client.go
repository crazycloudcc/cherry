package netwebsocket

/*
 * netwebsocket connection.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"sync"

	"github.com/crazycloudcc/cherry/base"

	"github.com/gorilla/websocket"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// Client TODO.
type Client struct {
	sync.RWMutex
	id      int32
	conn    *websocket.Conn
	serv    *Service
	chanRes chan []byte
}

/************************************************************************/
// export functions.
/************************************************************************/

/************************************************************************/
// moudule functions.
/************************************************************************/

// new connection.
func newClient(id int32, c *websocket.Conn, serv *Service) *Client {
	tc := new(Client)
	tc.id = id
	tc.conn = c
	tc.serv = serv
	tc.chanRes = make(chan []byte)

	go tc.read()
	go tc.write()

	return tc
}

// close
func (owner *Client) close() {
	if owner.id == -1 {
		base.LogDebug("client.go closed -------", owner.id, owner.chanRes)
		return
	}
	base.LogDebug("client.go close -------", owner.id, owner.chanRes)
	owner.id = -1
	close(owner.chanRes)
	owner.conn.Close()
	owner.serv.GetClientMngr().unregister <- owner
}

// send data to connection.
func (owner *Client) send(buf []byte) {
	owner.Lock()
	defer owner.Unlock()
	owner.chanRes <- buf
}

func (owner *Client) read() {
	defer func() {
		base.LogError("client read unregister: ", owner.id)
		owner.close()
	}()

	for {
		_, buf, err := owner.conn.ReadMessage()
		if err != nil {
			base.LogError("client read error: ", owner.id, err)
			break
		}
		base.LogDebug("client read data: ", owner.id, len(buf))
		msg := bytesToMsg(buf)
		owner.serv.GetRouter().Route(owner.id, msg.ID, msg.MetaData)
	}
}

func (owner *Client) write() {
	defer func() {
		base.LogError("client write error: ", owner.id)
		owner.close()
	}()

	for {
		select {
		case buf := <-owner.chanRes:
			if owner.id == -1 || buf == nil || len(buf) < 5 {
				break
			}
			base.LogDebug("client write data: ", owner.id, buf)
			owner.conn.WriteMessage(websocket.BinaryMessage, buf)
		}
	}
}

/************************************************************************/
// unit tests.
/************************************************************************/
