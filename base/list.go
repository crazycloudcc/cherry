/*
 * go container/list.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package base

import (
	"container/list"
	"errors"
	"fmt"
)

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type List struct {
	listDatas *list.List
}

/************************************************************************/
// export functions.
/************************************************************************/

// create new list.
func NewList(max int32) *List {
	l := new(List)
	l.listDatas = list.New()
	return l
}

// pushback.
func (this *List) PushBack(element interface{}) error {
	this.listDatas.PushBack(element)
	return nil
}

// pushfront.
func (this *List) PushFront(element interface{}) error {
	this.listDatas.PushFront(element)
	return nil
}

// insert.
func (this *List) Insert(index uint32, element interface{}) error {
	if index == 0 {
		this.PushFront(element)
		return nil
	}
	if index == uint32(this.listDatas.Len()) {
		this.PushBack(element)
		return nil
	}
	if index >= uint32(this.listDatas.Len()) {
		return errors.New("Insert index error!")
	}
	mark := this.listDatas.Front()
	var i uint32
	for i = 0; i < index; i++ {
		mark = mark.Next()
	}
	this.listDatas.InsertBefore(element, mark)
	return nil
}

// delete.
func (this *List) Del(index uint32) (interface{}, error) {
	if index >= uint32(this.listDatas.Len()) {
		return nil, errors.New("Del index error!")
	}
	mark := this.listDatas.Front()
	var i uint32
	for i = 0; i < index; i++ {
		mark = mark.Next()
	}
	ret := this.listDatas.Remove(mark)
	return ret, nil
}

// get.
func (this *List) Get(index uint32) (interface{}, error) {
	if index >= uint32(this.listDatas.Len()) {
		return nil, errors.New("Get index error!")
	}
	mark := this.listDatas.Front()
	var i uint32
	for i = 0; i < index; i++ {
		mark = mark.Next()
	}
	return mark.Value, nil
}

// length.
func (this *List) Len() int {
	return this.listDatas.Len()
}

// for range.
func (this *List) ForRange(f func(index int32, value interface{})) {
	len := int32(this.listDatas.Len())
	if len > 0 {
		mark := this.listDatas.Front()
		var i int32
		for i = 0; i < len; i++ {
			f(i, mark.Value)
			mark = mark.Next()
		}
	}
}

// data to string.
func (this *List) String() string {
	ret := "List Data: "
	len := this.listDatas.Len()
	ret += fmt.Sprintf("len[%d]", len)
	if len > 0 {
		mark := this.listDatas.Front()
		for i := 0; i < len; i++ {
			ret += fmt.Sprintf("{index[%d],data[%v]},", i, mark.Value)
			mark = mark.Next()
		}
	}
	ret += " End"
	return ret
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
