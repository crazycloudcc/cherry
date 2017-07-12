// 队列.
package base

import "math/rand"

/************************************************************************/
// 常量, 变量, 结构体, 接口定义.
/************************************************************************/

func Int() int {
    return rand.Int()
}

//[0,n)
func Intn(n int) int {
    return rand.Intn(n)
}

func Int32() int32 {
    return rand.Int31()
}

//[0,n)
func Int32n(n int32) int32 {
    return rand.Int31n(n)
}
