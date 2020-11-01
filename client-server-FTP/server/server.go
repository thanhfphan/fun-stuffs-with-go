package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:1202")
	fmt.Println("Listening on port 1202")
	if err != nil {
		fmt.Println("Error " + err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		s := string(buf[0:n])
		agrs := strings.Split(s, " ")
		command := agrs[0]

		if command == CD {
			chdir(conn, agrs[1])
		} else if command == DIR {
			dirList(conn)
		} else if command == PWD {
			pwd(conn)
		}
	}
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("CD error"))
	}
}

func dirList(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".") // . stand for current dir
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}

	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte("PWD error \r\n"))
		return
	}

	conn.Write([]byte(s))
}
