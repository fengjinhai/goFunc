package main
import (
    "fmt"
    "net/http"
    "io/ioutil"
)


func main(){
    resp, err := http.Get("http://www.baidu.com")
    if err != nil {
        fmt.Println(err)
    }
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(b))
}
