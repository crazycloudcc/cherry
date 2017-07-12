/*
 * binary tree.
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

type BTree struct {
	sync.RWMutex
}

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// 创建.
func NewBTree() *BTree {
	bt := new(BTree)
	return bt
}

// 插入.
func (this *BTree) Insert() error {
	return nil
}

// 查询.
func (this *BTree) Find() (*BTreeNode, error) {
	return nil, nil
}

// 删除.
func (this *BTree) Remove() error {
	return nil
}

// 前序遍历.
func (this *BTree) ForWithFront() []*BTreeNode {
	return nil
}

// 中序遍历.
func (this *BTree) ForWithMiddle() []*BTreeNode {
	return nil
}

// 后序遍历.
func (this *BTree) ForWithBack() []*BTreeNode {
	return nil
}

/************************************************************************/
// 模块内功能实现
/************************************************************************/

/************************************************************************/
// 模块内功能调试
/************************************************************************/
