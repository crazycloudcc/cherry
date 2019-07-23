/*
 * go map data.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */
package base

/************************************************************************/
// constants, variables, structs, interfaces.
/************************************************************************/

type Hash struct {
	table map[interface{}]interface{}
	count int32
}

/************************************************************************/
// export functions.
/************************************************************************/

// create new hash.
func NewHash(count int32) *Hash {
	m := new(Hash)
	m.table = make(map[interface{}]interface{})
	m.count = count
	return m
}

// add.
func (this *Hash) Add(key interface{}, value interface{}) {
	this.table[key] = value
}

// set.
func (this *Hash) Set(key interface{}, value interface{}) {
	this.table[key] = value
}

// get.
func (this *Hash) Get(key interface{}) interface{} {
	ret, ok := this.table[key]
	if ok {
		return ret
	}
	return nil
}

// delete.
func (this *Hash) Del(key interface{}) {
	delete(this.table, key)
}

// length.
func (this *Hash) Len() int {
	return len(this.table)
}

// for range.
func (this *Hash) ForRange(f func(key interface{}, value interface{})) {
	for k, v := range this.table {
		f(k, v)
	}
}

/************************************************************************/
// moudule functions.
/************************************************************************/

/************************************************************************/
// unit tests.
/************************************************************************/
