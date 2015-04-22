package main

import "fmt"

func adder() func (int) int {//go 函数可以闭包，函数可以作为变量或参数返回值
    sum:=0
    return func (x int) int {
        sum+=x
        return sum
    }
    
}

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {//实现了 返回斐波纳契
    x:=0
    y:=1
    return func()int{
        sum:=x+y
        x,y=y,sum
        return sum
    }
}

func main() {
    pos:=adder()
    for i:=0;i<10;i++{
        fmt.Println(pos(i))
    }
    f:=fibonacci()
    for j:=0;j<10;j++{
        fmt.Println(f())
    }
}