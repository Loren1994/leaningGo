package learn

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Learn19() {
	exitChan := make(chan int)
	go server("127.0.0.1:7001", exitChan)
	code := <-exitChan

	os.Exit(code)
}

func server(address string, exitChan chan int) {
	//侦听地址
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listener>>>>>:" + err.Error())
		exitChan <- 1
	}
	fmt.Println("侦听成功:" + address)
	defer listen.Close()
	//侦听循环
	for {
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println(err1.Error())
			continue
		}
		//开启会话
		go handleSession(conn, exitChan)
	}
}

func handleSession(conn net.Conn, exitChan chan int) {
	read := bufio.NewReader(conn)
	//循环接收数据
	for {
		str, err := read.ReadString('\n')
		if err == nil {
			str = strings.TrimSpace(str)
			// 处理Telnet指令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println(err)
			conn.Close()
			break
		}
	}
}

func processTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("server shutdown")
		exitChan <- 1
		return false
	}
	fmt.Println("input string:" + str)
	return true
}
