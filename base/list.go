/*
 * 链表. 对container/list的封装.
 * 默认以单向链表方式使用.
 * 可扩展至双向链表.
 */
package base

import (
	"container/list"
	"errors"
	"fmt"
)

/************************************************************************/
// 常量, 变量, 结构体, 接口定义.
/************************************************************************/

type List struct {
	listDatas *list.List
}

/************************************************************************/
// 模块对外接口.
/************************************************************************/

func NewList(max int32) *List {
	l := new(List)
	l.listDatas = list.New()
	return l
}

// 添加到列表尾部.
func (this *List) PushBack(element interface{}) error {
	this.listDatas.PushBack(element)
	return nil
}

// 添加到列表头部.
func (this *List) PushFront(element interface{}) error {
	this.listDatas.PushFront(element)
	return nil
}

// 插入.
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

// 删除.
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

// 获取.
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

// 长度.
func (this *List) Len() int {
	return this.listDatas.Len()
}

// 遍历.
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

// 移动到目标元素之前
func (this *List) MoveBefore(startIdx uint32, endIdx uint32) error {
	if startIdx >= uint32(this.listDatas.Len()) || startIdx >= uint32(this.listDatas.Len()) {
		return errors.New("Move index error!")
	}
	if startIdx == endIdx {
		return errors.New("Move index repeat!")
	}
	startEle := this.getEle(startIdx)
	endEle := this.getEle(endIdx)
	this.listDatas.MoveBefore(startEle, endEle)
	return nil
}

// 移动到目标元素之前
func (this *List) MoveAfter(startIdx uint32, endIdx uint32) error {
	if startIdx >= uint32(this.listDatas.Len()) || startIdx >= uint32(this.listDatas.Len()) {
		return errors.New("Move index error!")
	}
	if startIdx == endIdx {
		return errors.New("Move index repeat!")
	}
	startEle := this.getEle(startIdx)
	endEle := this.getEle(endIdx)
	this.listDatas.MoveAfter(startEle, endEle)
	return nil
}

// 字符串化.
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
// 模块内功能实现.
/************************************************************************/

//获取目标元素
func (this *List) getEle(index uint32) *list.Element {
	mark := this.listDatas.Front()
	var i uint32
	for i = 0; i < index; i++ {
		mark = mark.Next()
	}
	return mark
}

/************************************************************************/
// 模块内功能调试.
/************************************************************************/

/************************************************************************/
// 模块单元测试.
/************************************************************************/
