package elea

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var encoder net.Conn
var decoder net.Conn

// 使用socket作为中转，需要记录两个conn之间的关系，并做转发
func Aleph(ip, port string) {
	log.Println("本地ip地址：" + ip)

	ln, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := ln.Accept()
		log.Println("New Client")
		if err != nil {
			log.Panicln(err)
		}

		go readStream(conn)
	}
}

func readStream(conn net.Conn) {
	// 获取远程客户端IP:Port - conn.RemoteAddr()
	for {
		buffer := make([]byte, 13)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
			return
		}

		// 判断是否为encoder连接
		if buffer[:n][0] == 255 {
			// fmt.Println("encoder 连接")
			if decoder != nil {
				fmt.Println(buffer[7:n])
				num, err := decoder.Write(buffer[7:n])
				fmt.Println(num)
				if err != nil {
					log.Println(err)
				}
			}
		}

		// 判断是否为decoder连接
		if string(buffer[:7]) == "decoder" {
			fmt.Println("decoder 连接")
			// 别再读了，直接跳出吧
			decoder = conn
			break
		}
	}
}

func writeStream(conn net.Conn) {
	for i := range []int{1, 2, 3, 4} {
		fmt.Println(i)
		conn.Write([]byte("elea service message " + strconv.Itoa(i+1)))
		time.Sleep(time.Second * 1)
	}
}
