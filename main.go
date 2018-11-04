package main

import (
	"fmt"
	"net"
	commandFunc "redigo/command"
	"strings"
)

func parseCommand(buf []byte, cnt int) (command string, args []string, err error) {
	rawStrings := strings.Split(strings.TrimSpace(string(buf[0:cnt])), "\r\n")
	argLen := (len(rawStrings) - 3) / 2
	fmt.Printf("%d: %v\n", argLen, rawStrings)
	for i := 0; i < argLen; i++ {
		args = append(args, rawStrings[i*2+4])
	}
	return rawStrings[2], args, nil
}

func reformatResponse(resp string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(resp), resp)
}

func connHandler(conn net.Conn) {
	if conn == nil {
		println("empty Conn")
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := conn.Read(buf)
		if err != nil || cnt == 0 {
			conn.Close()
			break
		}
		command, args, err := parseCommand(buf, cnt)
		if err != nil {
			println("Parse Args Error")
			return
		}
		fmt.Printf("Command %s, Args %v\n", command, args)

		var result string
		switch strings.ToUpper(command) {
		case "PING":
			result = commandFunc.Ping(command, args)
		case "ECHO":
			result = commandFunc.Echo(command, args)
		default:
			result = commandFunc.Default(command, args)
		}
		conn.Write([]byte(reformatResponse(result)))
	}
}

func main() {
	println("redigo client start")
	server, err := net.Listen("tcp", ":6379")
	if err != nil {
		println("bind port error")
		return
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			println("connect fail")
			break
		}
		go connHandler(conn)
	}
	return
}
