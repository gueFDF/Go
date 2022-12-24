package main

import (
	"Go/net/codec"
	"bufio"
	"fmt"
	"net"
)

func work(fd net.Conn) {
	defer fd.Close()

	for {
		reader := bufio.NewReader(fd)

		//var buf [1024]byte

		//接收客户端发来的消息
		//n,err:=reader.Read(buf[:])
		msg,err:=codec.Decode(reader)
		if err!=nil {
			return 
		}

		//n, err := fd.Read(buf[:])

		// if err != nil {
		// 	fmt.Println("read from client err")
		// 	break
		// }
		//recvstr := string(buf[:n])

		fmt.Println(msg)

	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println("listen is err")
		return
	}

	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept is err")
			continue
		}

		go work(conn)
	}
}
