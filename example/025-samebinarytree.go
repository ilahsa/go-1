package main

import (
    "fmt"
    "code.google.com/p/go-tour/tree"
)

func walk(t *tree.Tree,ch chan int) {
    ch<-t.Value//写入本节点
    if t.Left !=nil{//写入左节点
        walk(t.Left,ch)
    }
    if t.Right!=nil{//写入右节点
        walk(t.Right,ch)
    }
}

func same(t1,t2 *tree.Tree) bool {
    ch:=make(chan int)
    ch1:=make(chan int)
    go walk(t1,ch)  
    go walk(t2,ch1)
    for{//循环比较
        v1:=<-ch
        v2:=<-ch1
        fmt.Println(v1,v2)
        if v1!=v2{
            return false
        }
    } 
    return true
}

func main() {
    t1:=tree.New(1)
    t2:=tree.New(1)
    fmt.Println(t1)
    fmt.Println(t2)
    fmt.Println(same(t1,t2))
}