package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func serverProcess(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("receive from client: ", recvStr)
		echo := "hi guys, your input is " + recvStr
		conn.Write([]byte(echo))
		if recvStr == "exit" {
			fmt.Println("close by client")
			break
		}
	}
}

func server() {

	addr := "127.0.0.1:10099"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("tcp listen failed, err: ", err)
		return
	}
	fmt.Println("tcp listen at ", addr)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed, err: ", err)
			continue
		}
		go serverProcess(conn)
	}
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:10099")
	if err != nil {
		fmt.Println("client error:", err)
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input,_ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [128]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("client receive fail, err:", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func SampleMain()  {
	// 同时测试server和client，需要另开一个协程运行server，否则client不会执行
	go server()
	client()
}