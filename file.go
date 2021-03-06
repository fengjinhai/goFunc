package main

import (
    "bufio"
    "fmt"
    "os"
    "io"
    "strings"
)

func check(e error){
    if e != nil {
        panic(e)
    }
}

//文件读取到切片里面，输入文件名
func read(filename string) []string{
    var datalines []string
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return datalines 
    }
    buf := bufio.NewReader(f)
    defer f.Close()
    for {
        line, err := buf.ReadString('\n')
        line = strings.TrimSpace(line)
        if err != nil {
            if err == io.EOF {
                return datalines 
            }
            fmt.Println(err)
            return datalines 
        }
        datalines = append(datalines, line)
    }
    return datalines 
}


//把切片里面的文件写入到本地
func write(filename string, datalines []string){
    f, err:= os.Create(filename)
    check(err)
    for i := range datalines {
        fmt.Println(datalines[i])
        n, err := f.WriteString(datalines[i] + "\n")
        check(err)
        fmt.Printf("write %d bytes\n", n)
    }
    defer f.Close()
}

func main(){
    var datalines []string
    datalines = read("./test.log")
    write("./test2.log", datalines)
}
