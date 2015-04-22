package main

import "fmt"

func main() {
    var p *int//*T声明一个类型T的指针，零值为nil
    i:=5
    p=&i//&从对象生成指针
    fmt.Println(p)//打印出内存地址
    *p=22//*从指针返回对象
    fmt.Println(*p)//打印出i的值此时i=22
    fmt.Println(i)
}