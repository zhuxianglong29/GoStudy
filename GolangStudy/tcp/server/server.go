package main

import (
	"fmt"
	"net"
	"strconv"
	_ "strings"
)

func process(conn net.Conn) {
	//循环接收
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		fmt.Printf("服务器在等待客户%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出err=%v", err)
			return
		}
		//3.显示客户端发送内容
		fmt.Print(string(buf[:n]))
		fmt.Printf("格式%T", buf)
		//**数据转换并处理
		strnub := string(buf[:n-1])
		fmt.Println("strnub=", strnub)
		k, _ := strconv.Atoi(strnub)
		fmt.Println("k=", k)
		j := strconv.Itoa(jisheng(k))
		//j:=jisheng(k)

		nub := []byte(j)
		fmt.Println("阶乘", jisheng(k))
		conn.Write(nub)
		fmt.Printf("j%T %v", j, j)
	}
}

func jisheng(n int) (result int) {
	if n > 1 {
		result = n * jisheng(n-1)
		return result
	}
	return 1
}

func main() {
	fmt.Println("服务器开始监")
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen err= ", err)
		return
	}
	defer listen.Close()

	//等待链接
	for {
		fmt.Println("等待连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Acceept err=", err)
		} else {
			fmt.Printf("Accept suc conn=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//协程-客户�?
		go process(conn)

	}

}
