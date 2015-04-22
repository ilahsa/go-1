package main

import "fmt"

func main() {
    for i:=1;i<10;i++{//go 中只有for 一种循环  for 后强制无须()
        fmt.Println(i)
    }
    j:=1
    for j<10 {//可以让前置后置语句为空
        fmt.Println(j)
        j=j+1
    }
    for{//死循环
        fmt.Println(1)
    }
}