/*
 * a star path finding.
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package base

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

// 路径节点. 包含基本的索引ID, 坐标, 是否阻挡信息.
type PathNode struct {
	id      int32 // 唯一ID.
	x, y    int32 // 坐标.
	isBlock bool  // 是否阻挡.
}

func (this *PathNode) SetBlock(flag bool) {
	this.isBlock = flag
}

type AStar struct {
}

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// 查找路径.
func (this *AStar) GetPath() error {
	return nil
}

/************************************************************************/
// 模块内功能实现
/************************************************************************/

// 从open表里取出一个节点.
func (this *AStar) popOpenList() *PathNode {
	return nil
}

// 将一个节点放入close表.
func (this *AStar) pushCloseList(pn *PathNode) error {
	return nil
}

/************************************************************************/
// 模块内功能调试
/************************************************************************/
