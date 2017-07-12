/*
 * go map data.
 * author: CC
 * email : crazycloudcc@gmail.com
 * date  : 2017.06.17
 */
package base

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type Map struct {
	table map[interface{}]interface{}
	count int32
}

/************************************************************************/
// 模块对外接口.
/************************************************************************/

// 创建map.
func NewMap(count int32) *Map {
	m := new(Map)
	m.table = make(map[interface{}]interface{})
	m.count = count
	return m
}

// 添加.
func (this *Map) Add(key interface{}, value interface{}) {
	this.table[key] = value
}

// 设置.
func (this *Map) Set(key interface{}, value interface{}) {
	this.table[key] = value
}

// 获取.
func (this *Map) Get(key interface{}) interface{} {
	ret, ok := this.table[key]
	if ok {
		return ret
	}
	return nil
}

// 删除.
func (this *Map) Del(key interface{}) {
	delete(this.table, key)
}

// 长度.
func (this *Map) Len() int {
	return len(this.table)
}

// 遍历.
func (this *Map) ForRange(f func(key interface{}, value interface{})) {
	for k, v := range this.table {
		f(k, v)
	}
}

/************************************************************************/
// 模块内功能实现
/************************************************************************/
