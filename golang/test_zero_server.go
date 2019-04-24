package main

import (
    "fmt"
    "zero"
)

func main() {
    fmt.Println("test")

    host := "127.0.0.1:18787"

    ss, err := zero.NewSocketService(host)
    if err != nil {
        return
    }
    // ss.SetHeartBeat(5*time.Second, 30*time.Second)
    ss.RegMessageHandler(HandleMessage)
    ss.RegConnectHandler(HandleConnect)
    ss.RegDisconnectHandler(HandleDisconnect)

    fmt.Println("service running on " + host)
    ss.Serv()
}

func HandleMessage(s *zero.Session, msg *zero.Message) {
    fmt.Println("receive msgID:", msg)
    fmt.Println("receive data:", string(msg.GetData()))
}

func HandleDisconnect(s *zero.Session, err error) {
    fmt.Println(s.GetConn().GetName() + " lost.")
}

func HandleConnect(s *zero.Session) {
    fmt.Println(s.GetConn().GetName() + " connected.")
}

