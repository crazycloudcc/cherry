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

// create and open channel.
func NewChan(count int32, listener ChanListener) *Chan {
	ret := new(Chan)
	ret.ch = make(chan interface{}, count)
	ret.closeFlag = false
	ret.listener = listener
	ret.count = count
	go ret.listen()
	return ret
}

// close and destroy channel.
func (this *Chan) Close() {
	if this.closeFlag {
		return
	}
	this.closeFlag = true
	this.ch <- nil
	close(this.ch)
}

// write data to channel.
func (this *Chan) Write(data interface{}) {
	if this.closeFlag {
		return
	}
	this.ch <- data
}

/************************************************************************/
// moudule functions.
/************************************************************************/

// goroutine read data from channel.
func (this *Chan) listen() {
	for {
		data := <-this.ch
		if this.closeFlag {
			// LogDebug("chan.go - listen exit.")
			break
		}
		this.listener(data)
	}
}
