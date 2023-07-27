package config

import (
	"io"
	"net"
	"net/http"
	"strings"
)

type Swagger struct {
	IpAddressType string `json:"ipAddressType"`
	ExteralIp     string `json:"externalIp"`
}

func (s *Swagger) String() string {
	return "ipAddressType: " + s.IpAddressType + ", ExteralIp: " + s.ExteralIp
}

func (s *Swagger) GetUsingIp() string {
	switch s.IpAddressType {
	case "external", "External":
		return s.ExteralIp
	case "public", "Public":
		return GetPublicIp()
	case "auto", "Auto":
		return GetIpWithMyexternalip()
	case "local", "localhost", "Localhost":
		return "localhost"
	default:
		return "localhost"
	}
}

func GetIpWithMyexternalip() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

func GetPublicIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
