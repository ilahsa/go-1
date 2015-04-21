package main

import "fmt"

//func 函数名 参数 返回值
func add1(x int,y int) int {
    return x+y
}

//多个同类型参数可以同时声明 x,y int
func add2(x,y int) int{
    return x+y
}

//函数可返回多个返回值
func add3(x,y int)(int,int) {
    return y,x
}

//函数的返回值可以被命名
func add4(x,y int) (r1,r2 int){
    r1=x+1
    r2=y+1
    return//没有参数的return 语句返回结果的当前值，仅在短函数中使用，否则容易影响代码可读性
    //return x,y
}

//函数之间惯例空行保持整洁
func main() {
    fmt.Println(add1(5,6))
    fmt.Println(add2(5,6))
    fmt.Println(add3(5,6))
    fmt.Println(add4(5,6))
}