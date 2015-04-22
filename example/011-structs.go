package main

import "fmt"

type Point struct{//type xxx struct{}声明一个结构体
    X int
    Y int
}

func main() {
    fmt.Println(Point{1,2})
    fmt.Println(Point{X:1})//name 赋值可指定部分字段其余字段用初始值
    v:=Point{5,6}
    fmt.Println(v.X)//结构体的字段用.来访问
    p:=&v
    fmt.Println(p)
    (*p).X=9
    p.X=8//如果 *p.X 可以访问 则可速记为 p.X
    fmt.Println(v.X)
}