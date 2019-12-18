package main

import (
    "fmt"
    "net"
    "os/exec"
    //"time"
)

//receives a reference to a connection, spawns a bash shell over the tcp connection
func handleConnection(connection net.Conn) {
    fmt.Printf("received connection from %v\n", connection.RemoteAddr().String()) //RemoteAddr refers to the machine connecting to the listener, while LocalAddr refers to the address/port of the listener itself

    _, err := connection.Write([]byte("connection successful, bash session over tcp initiated\n")) //convert the string to a byte slice and send it over the connection
    if err != nil {
        fmt.Println("Something went wrong trying to write to the connection:", err)
    }

    cmd := exec.Command("/bin/bash")
    cmd.Stdin = connection //connection pointer is dereferenced to retrieve the connection data
    cmd.Stdout = connection
    cmd.Stderr = connection

    cmd.Run()
}

func main() {
    var listenPort string = "4444"
    listener, err := net.Listen("tcp", "localhost:" + listenPort) //starts a listener on tcp port 4444

    if err != nil {
        fmt.Printf("An error occurred while initializing the listener on %v: %v\n", listenPort, err)
    } else {
        fmt.Println("listening on tcp port " + listenPort + "...")
    }

    //By removing this loop, you could have the program mimic netcat and end after one connection completes
    for {
        connection, err := listener.Accept() //waits for and returns the next connection to the listener
        if err != nil {
            fmt.Printf("An error occurred during an attempted connection: %v\n", err)
        }

        go handleConnection(connection) //go handle the connection concurrently in a goroutine

        //example code that would close the connection after 5 seconds
        /*time.Sleep(5 * time.Second)
        err = connection.Close()
        if err != nil {
            fmt.Println("Something went wrong when closing the connection:", err)
        }*/
    }
}

/* Something we learned here: Don't use pointers to interfaces, that is not how they are supposed to work. */
