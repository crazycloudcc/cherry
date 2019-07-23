/*
 * queue.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package base

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type Queue struct {
	sync.Mutex
	chanData chan interface{}
	count    int32
	max      int32
}

/************************************************************************/
// export functions.
/************************************************************************/

func NewQueue(max int32) *Queue {
	q := new(Queue)
	q.chanData = make(chan interface{}, max)
	q.count = 0
	q.max = max
	return q
}

func (this *Queue) Push(data interface{}) error {
	if this.count >= this.max {
		return errors.New("chan is full!")
	}
	this.chanData <- data
	this.count++
	return nil
}

func (this *Queue) Pop() (interface{}, error) {
	if this.count <= 0 {
		return nil, errors.New("chan is empty!")
	}

	data := <-this.chanData
	this.count--
	return data, nil
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/

func UnitTest_Queue(max int32, pushSleep, popSleep time.Duration) {
	q := NewQueue(max)

	go func() {
		index := 0
		for {
			err := q.Push(fmt.Sprintf("hello_%d", index))
			if err != nil {
				LogError("UnitTest_Queue Push Test error: ", err)
				continue
			}
			index++
			time.Sleep(pushSleep)
		}
	}()

	go func() {
		for {
			time.Sleep(popSleep)
			data, err := q.Pop()
			if err != nil {
				LogError("UnitTest_Queue Pop Test error: ", err)
				continue
			}
			LogDebug("UnitTest_Queue data: ", data)
		}
	}()
}
