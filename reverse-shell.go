package main

import (
    "fmt"
    "net"
)

func handleConnection(conn ) {

}

func main() {
    listener, err := net.Listen("tcp", ":4444") //starts a listener on tcp port 4444

    if err != nil {
        fmt.Println("Add code to handle errors here")
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Add code to handle errors here")
        }

        go handleConnection(conn)
    }
}
