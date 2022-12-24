package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	defer conn.Close() //关闭连接

	inputReader := bufio.NewReader(os.Stdin)

	for {
		//读取用户输入
		input,_:=inputReader.ReadString('\n')
		//去掉\n
		inputinfo:=strings.Trim(input,"\n")
		//输入q就推出
		if strings.ToUpper(inputinfo)=="Q" {
			return
		}
		_,err:=conn.Write([]byte(inputinfo))

		if err!=nil {
			return 
		}

	}
}
