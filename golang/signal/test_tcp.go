package main

import (
    "flag"
    "fmt"
    "net"
    "os"
    "os/exec"
    "os/signal"
    "syscall"
)

var (
    graceful = flag.Bool("graceful", false, "-graceful")
)

// Accepted accepted connection
type Accepted struct {
    conn net.Conn
    err  error
}

func handleConnection(conn net.Conn) {
    conn.Write([]byte("hello"))
    conn.Close()
}

func listenAndServe(ln net.Listener, sig chan os.Signal) {
    accepted := make(chan Accepted, 1)
    go func() {
        for {
            conn, err := ln.Accept()
            accepted <- Accepted{conn, err}
        }
    }()

    for {
        select {
        case a := <-accepted:
            if a.err == nil {
                fmt.Println("handle connection")
                go handleConnection(a.conn)
            }
        case _ = <-sig:
            fmt.Println("gonna fork and run")
            forkAndRun(ln)
            break
        }
    }
}

func gracefulListener() net.Listener {
    ln, err := net.FileListener(os.NewFile(3, "graceful server"))
    if err != nil {
        fmt.Println(err)
    }

    return ln
}

func firstBootListener() net.Listener {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
    }

    return ln
}

func forkAndRun(ln net.Listener) {
    l := ln.(*net.TCPListener)
    newFile, _ := l.File()
    fmt.Println(newFile.Fd())

    cmd := exec.Command(os.Args[0], "-graceful")
    cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
    cmd.ExtraFiles = []*os.File{newFile}
    cmd.Run()
}

func main() {
    flag.Parse()
    fmt.Printf("given args: %t, pid: %d\n", *graceful, os.Getpid())
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGUSR1)

    var ln net.Listener
    if *graceful {
        ln = gracefulListener()
    } else {
        ln = firstBootListener()
    }

    listenAndServe(ln, c)
}