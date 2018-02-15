package main

import(
  "fmt"
  "bufio"
  "net"
  "io/ioutil"
)

func check(err error, message string) {
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", message)
}

func main() {
  logo, err := ioutil.ReadFile("logo.ans")

  ln, err := net.Listen("tcp", ":8080")
  check(err, "Server ready!")

  for {
    conn, err := ln.Accept()
    check(err, "Accepted connection!")

    go func() {
      buf := bufio.NewReader(conn)

      conn.Write([]byte("\033[H\033[2J"))
      conn.Write([]byte("\033[33m"));
      conn.Write([]byte("Welcome to The Ape Machine BBS.\n\n"))
      conn.Write([]byte(logo))

      for {
        conn.Write([]byte(">"))
        name, err := buf.ReadString('\n')

        if err != nil {
          fmt.Printf("Client disconnected.\n")
          break
        }

        conn.Write([]byte("Hello, " + name))
      }
    }()
  }

  // fmt.Printf("GoGo(gadget)CloudTube!\n")
  // fmt.Printf("Written by Daniel Owen van Dommelen\n")
}
