package main

/* A simple tcp client. This is nowhere near functional or complete, I am simply keeping it here for now.*/

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    var tcpPort string = "4444"
    connection, err := net.Dial("tcp", "127.0.0.1:" + tcpPort) //connect to the socket
    if err != nil {
        fmt.Println("An error occurred trying to connect to the target:", err)
    }

    for {
        //receive reply from server and print
        message, _ := bufio.NewReader(connection).ReadString('\n') //waits and receives a reply from the server
        //fmt.Print("Message from server: " + message)
        fmt.Print(message)

        //read input from standard in
        reader := bufio.NewReader(os.Stdin)
        //fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')

        //write input to tcp socket
        fmt.Fprintf(connection, text + "\n") //formats and writes to a given io.Writer object, in this case the connection
    }
}

/* To make this work more optimally:

   Create a channel to handle all stdout operations (everything to be printed to your terminal).
   Instead of reading one line from your stdin, and then waiting for one line from stdout (from the connection),
   send all stdout to a channel. Have a separate thread which reads from that channel and displays the information
   to the terminal.

   IN A NUTSHELL:
   ---------------------------------
   READING FROM STDIN AND READING FROM STDOUT LIVE WITHIN THEIR OWN GOROUTINES.
   THAT IS WHY A NORMAL TERMINAL OPERATION DOES NOT FREEZE UP YOUR INPUTS WHEN THINGS ARE RUNNING.
   YOU MAY NEED TO IMPLEMENT A LOCK OF SOME SORT SO THAT THE SERVER DOES NOT TRY AND DO MORE THAN ONE OPERATION
   WHILE IT IS ALREADY SENDING DATA.
   HONESTLY, THIS IS KINDA INTERESTING. I DON'T REALLY KNOW IF I AM EVEN SAYING THIS RIGHT. WE WILL FIGURE IT OUT.
   */
