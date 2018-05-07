package main
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "time"
    "os"
    "strings"
    "net"
    "golang.org/x/net/proxy"
)

/*
简单的网页抓取，相当于wget
*/
func httpGet(){
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

/*
网页抓取，post请求
*/
func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func Socks5Client(add string, auth ...*proxy.Auth) (client *http.Client, err error){
    dialer , err := proxy.SOCKS5("tcp", "127.0.0.1:9090", 
        nil,
        &net.Dialer{
            Timeout: 30 * time.Second,
            KeepAlive: 30 * time.Second,
        },
    )
    if err != nil {
        return
    }
    transport := &http.Transport{
        Proxy: nil,
        Dial: dialer.Dial,
        TLSHandshakeTimeout: 10 * time.Second,
    }
    client = &http.Client {Transport: transport}
    return
}

/*
使用代理服务器进行抓取
这里假设代理服务器的IP 为 127.0.0.1，端口为 8080
*/
func crawerByPoxy() {
    client, err := Socks5Client("127.0.0.1:8080")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    resp, err := client.Get("http://www.baidu.com")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(b))
    }
}
