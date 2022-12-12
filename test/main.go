package main

import (
	"NSockets"
	"fmt"
)

func main() {
	fmt.Println(NSockets.ConnectUDP("192.168.15.74:60000"))
	defer NSockets.CloseSession()
	fmt.Println(NSockets.SendMsg("hello world"))
}
