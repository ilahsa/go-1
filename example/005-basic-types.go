package main

import "fmt"

func main() {
    var(
        result bool=true//bool
        str string="I love you"//string
        i int=1//int int8 int16 int32 int64 int会根据系统来变在 32bit os中int 等于int32
        j uint=1//uint uint8 uint16 uint32 uint64 uintptr uint 类似int 由操作系统来决定
        b byte=2//byte uint8的别名
        u rune=5678//int32 的别名 代表一个unicode 码
        f float32=4.5//float32 float64
        f1 float32=float32(i)//T(v)可将v转换为类型T
        c complex64//complex64 complex128 表示复数
    )
    k:=5//可省略类型，使用:=来进行类型推到，自动使用右侧类型

    fmt.Println(result,str,i,k,j,b,u,f,f1,c)
}