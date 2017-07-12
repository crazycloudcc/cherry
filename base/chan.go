/*
 * go channel.
 * author: CC
 * email : crazycloudcc@gmail.com
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
// 模块对外接口.
/************************************************************************/

// 创建通道.
// 参数列表: 1-元素数量, 2-通道数据接收函数.
func NewChan(count int32, listener ChanListener) *Chan {
	ret := new(Chan)
	ret.ch = make(chan interface{}, count)
	ret.closeFlag = false
	ret.listener = listener
	ret.count = count
	go ret.listen()
	return ret
}

// 关闭通道.
func (this *Chan) Close() {
	this.closeFlag = true
	this.ch <- nil
	close(this.ch)
}

// 向通道写入数据.
func (this *Chan) Write(data interface{}) {
	this.ch <- data
}

/************************************************************************/
// 模块内功能实现
/************************************************************************/

// goroutine 读取数据.
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
