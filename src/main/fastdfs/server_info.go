package fastdfs

import (
	"github.com/fastdfs-client-go/src/main/common"
	"net"
)

type ServerInfo struct {
	ipAddr string
}

func NewServerInfo(ipAddr string) *ServerInfo {
	return &ServerInfo{ipAddr: ipAddr}
}

func (si *ServerInfo) Connect() (conn *net.TCPConn, e error) {
	return common.GetConnection(si.ipAddr)
}
