package nettcp

/*
 * 网络连接管理.
 * 管理Conn类型实例.
 */

import (
	"cherry/base"
	"errors"
	"sync"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// ConnEventCallback type callback
type ConnEventCallback func(id int32, uid uint64)

// ClientMngr struct.
type ClientMngr struct {
	sync.RWMutex
	seed        int32
	callbackAdd ConnEventCallback
	callbackDel ConnEventCallback
	hashClients *base.Hash // connection id -> *Client.
	hashUIDs    *base.Hash // connection id -> uid.
}

/************************************************************************/
// export functions.
/************************************************************************/

// AddUID add hashUIDs data. (login verify success, logic call this function.)
func (owner *ClientMngr) AddUID(id int32, uid uint64) {
	owner.Lock()
	defer owner.Unlock()
	owner.hashUIDs.Add(id, uid)
	owner.callbackAdd(id, uid)
}

// DelUID del hashUIDs data. (logic call this function.)
func (owner *ClientMngr) DelUID(id int32) {
	owner.Lock()
	defer owner.Unlock()
	conn := owner.hashClients.Get(id)
	conn.(*Client).doClose()

	uid := owner.hashUIDs.Get(id)
	if uid != nil {
		owner.callbackDel(id, uid.(uint64))
	}
	owner.delConnection(id)
}

// CheckConnection function.
func (owner *ClientMngr) CheckConnection(id int32) error {
	owner.RLock()
	defer owner.RUnlock()
	uid := owner.hashUIDs.Get(id)
	if uid == nil {
		return errors.New("invalid connection")
	}
	return nil
}

// GetUIDByConnection function.
func (owner *ClientMngr) GetUIDByConnection(id int32) uint64 {
	owner.RLock()
	defer owner.RUnlock()
	uid := owner.hashUIDs.Get(id)
	if uid != nil {
		return uid.(uint64)
	}
	return 0
}

// SendToClient function.
func (owner *ClientMngr) SendToClient(id int32, msg *Msg) {
	conn := owner.hashClients.Get(id)
	if conn == nil {
		base.LogWarn("SendToClient:", id)
		return
	}
	conn.(*Client).Send(msg)
}

// SendToAll function.
func (owner *ClientMngr) SendToAll(msg *Msg) {
	owner.hashClients.ForRange(func(id, conn interface{}) {
		conn.(*Client).Send(msg)
	})
}

// RefreshHBTime function.
func (owner *ClientMngr) RefreshHBTime(id int32) {
	conn := owner.hashClients.Get(id)
	if conn == nil {
		return
	}
	conn.(*Client).refreshHBTime()
}

/************************************************************************/
// moudule functions.
/************************************************************************/

// create new ClientMngr.
func newClientMngr(max int32) *ClientMngr {
	tcm := new(ClientMngr)
	tcm.hashClients = base.NewHash(1)
	tcm.hashUIDs = base.NewHash(1)
	return tcm
}

// new connection event.
func (owner *ClientMngr) addConnection(c *Client) int32 {
	owner.Lock()
	defer owner.Unlock()
	defer owner.seedTick()
	id := owner.seed
	owner.hashClients.Add(id, c)
	c.Run(id)
	return id
}

// del hashClients data.
func (owner *ClientMngr) delConnection(id int32) {
	owner.hashUIDs.Del(id)
	owner.hashClients.Del(id)
}

// seed tick.
func (owner *ClientMngr) seedTick() {
	owner.seed++
}

/************************************************************************/
// unit tests.
/************************************************************************/
