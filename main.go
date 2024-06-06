package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	PORT := 8082

	addr := net.UDPAddr{
		Port: PORT,
		IP:   net.ParseIP("0.0.0.0"),
	}

	server, err := net.ListenUDP("udp", &addr)

	if err != nil {
		fmt.Println("Listen error ", err)
		os.Exit(-1)
	}

	fmt.Println("Listen at ", addr.String())

	// Goroutine for handle send message
	go handleSendMessage(server)

	for {
		// Received message
		buffer := make([]byte, 1024)
		n, addr, err := server.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Read message received error ", err)
			continue
		}
		msg := string(buffer[:n])
		fmt.Printf("%s: %s\n", addr, msg)

		// Send message

	}
}

func handleSendMessage(server *net.UDPConn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read message input error ", err)
			continue
		}

		// client address
		receiveAddr := net.UDPAddr{
			IP:   net.ParseIP("10.2.49.53"),
			Port: 8080,
		}

		_, err = server.WriteToUDP([]byte(text), &receiveAddr)
		if err != nil {
			fmt.Print("-> Send message error\n", err)
		}
	}
}
