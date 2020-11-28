package getip

import (
	"fmt"
	"net"
	"strings"
)

func GetOutBoundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return strings.Split(localAddr.IP.String(), ":")[0], nil
}
