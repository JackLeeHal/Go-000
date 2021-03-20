package main

import (
    "bufio"
    "log"
    "net"
)

func main() {
    listen, err := net.Listen("tcp", "6666")
    if err != nil {
        log.Panicf("listen error: %v\n", err)
    }

    for {
        conn, err := listen.Accept()
        if err != nil {
            log.Printf("accept error: %v\n", err)
            continue
        }

        go handleConn(conn)
    }

}

func handleConn(conn net.Conn) {
    ch := make(chan []byte, 1024)
    go readConn(conn, ch)
    go writeConn(conn, ch)
}

func readConn(conn net.Conn, ch chan<- []byte) {
    defer close(ch)

    reader := bufio.NewReader(conn)

    for {
        data, _, err := reader.ReadLine()
        if err != nil {
            log.Printf("read error: %v\n", err)
            break
        }

        ch <- data
    }
}

func writeConn(conn net.Conn, ch <-chan []byte) {
    defer conn.Close()

    writer := bufio.NewWriter(conn)

    for {
        select {
        case data, ok := <-ch:
            if !ok {
                log.Println("writeConn ch closed")
                return
            }

            writer.WriteString("receive data: ")
            writer.Write(data)
            writer.Flush()
        }
    }
}
