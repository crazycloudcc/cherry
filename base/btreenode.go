/*
 * binary tree node.
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package base

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// 树节点.
type BTreeNode struct {
	data   interface{} // 数据.
	parent *BTreeNode  // 父节点.
	lChild *BTreeNode  // 左子节点.
	rChild *BTreeNode  // 右子节点.
}

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// 查询节点数据.
func (this *BTreeNode) GetData() interface{} {
	return this.data
}

// 查询父节点.
func (this *BTreeNode) GetParent() *BTreeNode {
	return this.parent
}

// 查询左子节点.
func (this *BTreeNode) GetLChild() *BTreeNode {
	return this.lChild
}

// 查询右子节点.
func (this *BTreeNode) GetRChild() *BTreeNode {
	return this.rChild
}

/************************************************************************/
// 模块内功能实现
/************************************************************************/

/************************************************************************/
// 模块内功能调试
/************************************************************************/
