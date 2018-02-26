package main

import(
  "fmt"
  "bufio"
  "net"
  "io/ioutil"
  "strings"
  "time"
)

func check(err error, message string) {
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", message)
}

type ClientJob struct {
  cmd string
  conn net.Conn
}

func generateResponses(clientJobs chan ClientJob) {
  for {
    clientJob := <-clientJobs

    for start := time.Now(); time.Now().Sub(start) < time.Second; {
    }

    if strings.Compare("logout", clientJob.cmd) == 0 {
      clientJob.conn.Write([]byte("Bye!"))
      fmt.Printf("Client logged out.\n")
      break
    } else {
      fmt.Printf(clientJob.cmd)
      clientJob.conn.Write([]byte(clientJob.cmd))
    }
  }
}

func main() {
  logo, err  := ioutil.ReadFile("logo.ans")
  clientJobs := make(chan ClientJob)
  
  go generateResponses(clientJobs)

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
        cmd, err := buf.ReadString('\n')
        cmd       = strings.Replace(cmd, "\r\n", "", -1)

        if err != nil {
          fmt.Printf("Client disconnected.\n")
          break
        }

        clientJobs <- ClientJob{cmd, conn}
      }
    }()
  }
}
