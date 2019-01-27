package main

import (
    "net/http"
    "crypto/tls"
    "time"
    "net"
    "fmt"
    "sync"
    "io/ioutil"
)


var (
    DialTimeout               = 200 * time.Millisecond
    TLSHandshakeTimeout       = 300 * time.Millisecond
    OutNetDialTimeout         = 10 * time.Second
    OutNetTLSHandshakeTimeout = 2 * time.Second
    ResponseHeaderTimeout     = 700 * time.Millisecond
    SleepTime                 = 100 * time.Microsecond
    KeepAlive                 = 60 * time.Second
    MaxIdleConnections        = 0
)



type Client struct {
    cli *http.Client
}


func New() *Client {
    c := new(Client)
    c.cli = &http.Client{
        Transport: &http.Transport{
            Dial: (&net.Dialer{
                Timeout: DialTimeout,
                KeepAlive: KeepAlive,
            }).Dial,
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
            TLSHandshakeTimeout: TLSHandshakeTimeout,
            // ResponseHeaderTimeout: ResponseHeaderTimeout,
            MaxIdleConnsPerHost: MaxIdleConnections,
            DisableKeepAlives:true,
        },
    }
    return c
}


func (c *Client) GET(url string, retry int) (resp *http.Response, err error) {
    if retry < 1 {
        retry = 1
    }

    for i := 0; i < retry; i++ {
        resp, err = c.cli.Get(url)
        if err == nil {
            break
        }
        fmt.Println("------GET-------retry---", i, c.cli, url)
        // time.Sleep(SleepTime)
    }
    return
}

// var Conf = &httpclient.Conf{RequestTimeout: 3, MaxIdleConnsPerHost: 10}

var Url = "http://testkeepaliveserving.wordstore.knowbox.cn/sleep"
// var Url = "http://10.9.71.78:6601"

// var c = CreateHttpClient(10,5)
var c = New()

func SendMessageAPI() (err error) {
    //把post表单发送给目标服务器
    // c := &http.Client{
    //     Timeout: 7 * time.Second,
    // }
    resp, err := c.GET(Url,1)

    if err != nil {
        fmt.Println(err.Error())
        return err
    }
    defer resp.Body.Close()
    buf, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf(string(buf), err)    
    }
    return nil
}

func main() {
    fmt.Println("-------start-----")
    MaxSend := 3
    stop := make(chan bool, MaxSend)
    m := 0
    n := 0
    time_1 := 0
    time_2 := 0
    time_3 := 0
    time_4 := 0
    time_5 := 0
    time_6 := 0
    // var wg sync.WaitGroup
    var beginWg sync.WaitGroup
    f := func(i int) {
        flag := false
        count := 0
        beginWg.Wait()
        for {
            if flag {
                break
            }
            select {
            case <-stop:
                // fmt.Println("stop pop one element", i)
                flag = true
            default:
                // fmt.Println("in f()")
                t1 := time.Now()
                err := SendMessageAPI()
                elapsed := time.Since(t1)
                if err != nil {
                    fmt.Println(err)
                    n += 1
                } else {
                    m += 1
                }
                if elapsed < 500*time.Millisecond {
                    time_1 += 1
                } else if elapsed < 1*time.Second {
                    time_2 += 1
                } else if elapsed < 3*time.Second {
                    time_3 += 1
                } else if elapsed < 5*time.Second {
                    time_4 += 1
                } else if elapsed < 10*time.Second {
                    time_5 += 1
                } else {
                    time_6 += 1
                }
            }

            count++
            if count>2 {
                break
            }
        }
    }
    // f(1)
    for i := 0; i < MaxSend; i++ {
        fmt.Println("--------------",i,"--------------")
        beginWg.Add(1)
    }
    fmt.Println("working...")
    for i := 0; i < MaxSend; i++ {
        go f(i)
        beginWg.Done()
    }

    time.Sleep(30 * time.Second)
    fmt.Println("Time:", 300, "s success:", m, "failure:", n)
    fmt.Println("\t cost Time 500ms :", time_1, "\t <1s", time_2, "\t <3s:", time_3, "\t >3s:", time_4, "\t >5s:", time_5, "\t >10s:", time_6)
    for i := 0; i < MaxSend; i++ {
        stop <- true
    }
    time.Sleep(5 * time.Second)
}