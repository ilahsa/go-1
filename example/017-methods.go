package main

import (
    "fmt"
    "math"
)

type Vertex struct{
    x,y float64
}

func (v *Vertex) Abs()float64{//可为结构体添加方法
    return math.Sqrt(v.x*v.x+v.y*v.y)
}

/*
你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
但是，不能对来自其他包的类型或基础类型定义方法。
*/

func main() {
    v:= &Vertex{3,4}
    fmt.Println(v.Abs())
}