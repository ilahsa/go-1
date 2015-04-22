package main

import "fmt"

type IDoor interface{//定义接口
    open()//接口中定义方法
}

type PwdDoor struct{

}
type NormalDoor struct{

}

func(nd NormalDoor) open() {//实现接口定义的方法
    fmt.Println("Open the door with out any thing")    
}

func (pd PwdDoor) open(){
    fmt.Println("Open the door with password")
}
func main() {
    PwdDoor{}.open()
    NormalDoor{}.open()
}