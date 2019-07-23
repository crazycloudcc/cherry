/*
 * stack.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package base

import (
	"container/list"
	"errors"
	"sync"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type Stack struct {
	sync.Mutex
	listDatas *list.List
}

/************************************************************************/
// export functions.
/************************************************************************/

func NewStack(max int32) *Stack {
	s := new(Stack)
	s.listDatas = list.New()
	return s
}

func (this *Stack) Pop() (interface{}, error) {
	if this.listDatas.Len() <= 0 {
		return nil, errors.New("stack is empty!")
	}
	ret := this.listDatas.Remove(this.listDatas.Front())
	return ret, nil
}

func (this *Stack) Push(data interface{}) error {
	this.listDatas.PushBack(data)
	return nil
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
