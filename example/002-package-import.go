package main//每个go程序都是由包组成的，程序运行的入口是包main
//包名与导入之间按照惯例空行
import(//括号可以导入多个包,称作打包导入语句也可以分开 import,
    "fmt"
    "math/rand"
)
/*
这是多行注释

按照惯例一般包名与导入路径的最后一个目录一致
例如：import "math/rand" 的包名一般为 package rand
导入包中首字母大写的会被导出
*/
//导入与函数之间按照惯例空行
func main(){
    fmt.Println("My favorite number is",rand.Intn(10))
}