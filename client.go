


package main

import (
"fmt"
"net"
)

func main() {
	fmt.Println("开始了.........")
	listener,err:= net.Listen("tcp","localhost:5001")
	if err != nil {
		fmt.Println("错误",err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("出现错误", err.Error())
			return
		}
		go doServerStuff(conn)

	}

}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("程序错误",err.Error())
			return
		}
		fmt.Println("接收到的是：%v", string(buf[:len]))
	}

}

