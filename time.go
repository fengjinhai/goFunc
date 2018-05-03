package main

import (
    "fmt"
    "time"
)

func main() {
    //获取当前时间戳
    timestamp := time.Now().Unix()
    fmt.Println(timestamp)

    //时间戳格式化
    tm := time.Unix(timestamp, 0)
    fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
    fmt.Println(tm.Format("02/01/2006"))

    //格式化时间转时间戳
    t, _ := time.Parse("01/02/2006", "06/21/2017")
    fmt.Println(t.Unix())
}
