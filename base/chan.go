/*
 * go channel.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package base

import (
	"sync"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type ChanListener func(interface{})

type Chan struct {
	sync.Mutex
	ch        chan interface{}
	closeFlag bool
	listener  ChanListener
	count     int32
}

/************************************************************************/
// export functions.
/************************************************************************/

// NewChan create and open channel.
func NewChan(count int32, listener ChanListener) *Chan {
	ret := new(Chan)
	ret.ch = make(chan interface{}, count)
	ret.closeFlag = false
	ret.listener = listener
	ret.count = count
	go ret.listen()
	return ret
}

// Close close and destroy channel.
func (owner *Chan) Close() {
	if owner.closeFlag {
		return
	}
	owner.closeFlag = true
	owner.ch <- nil
	close(owner.ch)
}

// Write write data to channel.
func (owner *Chan) Write(data interface{}) {
	if owner.closeFlag {
		return
	}
	owner.ch <- data
}

/************************************************************************/
// moudule functions.
/************************************************************************/

// goroutine read data from channel.
func (owner *Chan) listen() {
	for {
		data := <-owner.ch
		if owner.closeFlag {
			// LogDebug("chan.go - listen exit.")
			break
		}
		owner.listener(data)
	}
}
