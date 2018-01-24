package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:7777") //获取一个TCP地址信息,TCPAddr
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr) //创建一个TCP连接:TCPConn
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")) //发送HTTP请求头
	checkError(err)
	result, err := ioutil.ReadAll(conn) //获得返回数据
	checkError(err)
	fmt.Println(string(result))
	os.Exit(1)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}