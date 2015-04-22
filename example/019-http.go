package main

import(
    "fmt"
    "log"
    "net/http"
)

type Hello struct{

}

/* http 中的 接口 Handler
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
*/
func (h Hello) ServeHTTP(//Hello 实现了http.Handler 的接口 实现方法 ServeHTTP
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w,"Hello Go")
}
func main() {
    var h Hello
    err:=http.ListenAndServe("localhost:4000",h)
    if err!=nil{
        log.Fatal(err)
    }
}