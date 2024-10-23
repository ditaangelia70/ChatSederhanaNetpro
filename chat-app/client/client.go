package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    go listenForMessages(conn)

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter message: ")
        msg, _ := reader.ReadString('\n')
        conn.Write([]byte(msg))
    }
}

func listenForMessages(conn net.Conn) {
    for {
        msg := make([]byte, 256)
        n, err := conn.Read(msg)
        if err != nil {
            fmt.Println("Error reading from server:", err)
            return
        }
        fmt.Print(string(msg[:n]))
    }
}
