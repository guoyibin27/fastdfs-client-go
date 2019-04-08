package fastdfs

import (
	"net"
)

type ServerInfo struct {
	ipAddr string
}

func NewServerInfo(ipAddr string) *ServerInfo {
	return &ServerInfo{ipAddr: ipAddr}
}

func (si *ServerInfo) Connect() (conn *net.TCPConn, e error) {
	return GetConnection(si.ipAddr)
}
