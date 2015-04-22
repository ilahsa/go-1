package main

import "fmt"

var c,python,java bool//变量可定义在包或函数级别

func main() {
    var i int//类型在变量名称后，变量定义不赋值会输出零值，即初始值，数值类型零值：0，字符串零值：”“，bool零值: false
    var j int=1//变量声明的时候可被赋值
    var k,l,o =2,5,"ccc"//可同时为多个声明变量赋值,可同时声明多种类型变量 直接赋值省略类型，不赋值无法同时声明多种类型
    q:=3//:=是一种语法糖，右侧为表达式时可免去声明类型，仅能在函数中使用，函数外必须一var func 等开头
    fmt.Println(c,python,java,i,j,k,l,o,q)
    var (
        x,y int
        cc,dd string
    )//可用括号混合声明
    var(
        xx,yy=1,3
        ccc,ddd="ccc","ccccddd"
    )//混合声明并赋值
    fmt.Println(x,y,cc,dd,xx,yy,ccc,ddd)
}