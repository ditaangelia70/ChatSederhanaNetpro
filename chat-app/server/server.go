package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

var clients []net.Conn

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server is running on port 8080")

    // Goroutine untuk menerima input dari server sendiri
    go handleServerInput()

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        clients = append(clients, conn)
        go handleClient(conn)
    }
}

// Fungsi untuk menangani input dari server sendiri
func handleServerInput() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Server: Enter message: ")
        msg, _ := reader.ReadString('\n')

        // Kirim pesan dari server ke semua klien
        broadcast("Server: " + msg, nil)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from client:", err)
            return
        }

        // Menampilkan pesan dari klien di server
        fmt.Print("Received from client: ", msg)

        // Kirim pesan dari klien ke semua klien lain
        broadcast(msg, conn)
    }
}

// Fungsi untuk mengirim pesan ke semua klien
func broadcast(msg string, sender net.Conn) {
    for _, client := range clients {
        if client != sender { // Jangan kirim kembali ke pengirim
            client.Write([]byte(msg))
        }
    }
}
