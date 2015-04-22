package main

import "fmt"

func main() {
    var m map[string]int //map 是k v 集合
    fmt.Println(m)//零值为nil 不能赋值
    m=make(map[string]int)//map 使用前必须make
    m["a"]=3
    fmt.Println(m)
    fmt.Println(m["a"])
    var m1 = map[string]int{//可以类似结构体一样初始化，必须包含键名
        "a":3,
        "b":4,//,是必须得
    }
    m1["a"]=0//修改键值
    m1["c"]=6
    fmt.Println(m1)
    elem,ok:=m1["o"]//m1["o"]存在则ok为true
    if ok{
        fmt.Println(elem)
    }
    fmt.Println(m1["q"])//不存在时在为类型的零值
}