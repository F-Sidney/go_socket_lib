package NSockets

import (
	"fmt"
	"net"
)

var NSession net.Conn

func ConnectUDP(ipAndPort string) (result string) {
	if NSession != nil {
		NSession.Close()
	}

	curs, err := net.Dial("udp", ipAndPort)
	if err != nil {
		return fmt.Sprintf("connect error: %s, remote ip is: %s", err.Error(), ipAndPort)
	}

	NSession = curs
	return fmt.Sprintf("connect server: %s", ipAndPort)
}

func SendMsg(msg string) (reply string) {
	if NSession != nil {
		_, err := NSession.Write([]byte(msg))
		if err != nil {
			reply = fmt.Sprintf("error: %s", err.Error())
			return
		}

		rp := make([]byte, 1024)
		_, err = NSession.Read(rp)
		if err != nil {
			reply = fmt.Sprintf("error: %s", err.Error())
			return
		}
		reply = fmt.Sprintf("msg send:%s\nmsg got:%s", msg, string(rp))
		return
	} else {
		reply = fmt.Sprintf("no session, can not send message")
		return
	}
}

func CloseSession() (reply string) {
	if NSession != nil {
		err := NSession.Close()
		if err != nil {
			reply = fmt.Sprintf("error: %s", err.Error())
		} else {
			reply = fmt.Sprintf("Disconnected!\n")
		}
	} else {
		reply = fmt.Sprintf("No session need to disconnect!")
	}

	return
}
