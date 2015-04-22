package main

import "fmt"

func main() {
    defer fmt.Println("GO")//defer 会延迟执行直到上层函数返回，延迟函数的参数会立刻生成
    //defer栈 多条defer语句会压入一个栈后进先出 后边的defer 更早执行
    //defer 一般用在 打开资源前 defer file.close()这样即使后续出现异常也可以保证文件正常关闭
    fmt.Println("Hello")
}