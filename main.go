package main

import (
	"fmt"
	"net"
	"redigo/command"
	"strings"
)

func parseCommand(buf []byte, cnt int) (command string, args []string, err error) {
	rawStrings := strings.Split(strings.TrimSpace(string(buf[0:cnt])), "\r\n")
	argLen := (len(rawStrings) - 3) / 2
	for i := 0; i < argLen; i++ {
		args = append(args, rawStrings[i*2+4])
	}
	return strings.ToUpper(rawStrings[2]), args, nil
}

func reformatResponse(resp string, err error) string {
	if err != nil {
		return fmt.Sprintf("-ERR %s\r\n", resp)
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(resp), resp)
}

func connHandler(conn net.Conn) {
	println("New Conn Start")
	if conn == nil {
		println("empty Conn")
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := conn.Read(buf)
		if err != nil || cnt == 0 {
			println("Conn closed by client")
			conn.Close()
			break
		}
		upperCommand, args, err := parseCommand(buf, cnt)
		if err != nil {
			println("Parse Args Error")
			return
		}
		fmt.Printf("Command %s, Args %v\n", upperCommand, args)

		var result string
		switch upperCommand {
		case "PING":
			result, err = command.Ping(upperCommand, args)
		case "ECHO":
			result, err = command.Echo(upperCommand, args)
		case "GET":
			result, err = command.Get(upperCommand, args)
		case "SET":
			result, err = command.Set(upperCommand, args)
		default:
			result, err = command.Default(upperCommand, args)
		}

		conn.Write([]byte(reformatResponse(result, err)))
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
