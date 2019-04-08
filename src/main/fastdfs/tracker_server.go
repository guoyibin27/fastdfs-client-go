package fastdfs

import (
	"net"
)

type TrackerServer struct {
	conn *net.TCPConn
}

func NewTrackerServer(conn *net.TCPConn) *TrackerServer {
	server := new(TrackerServer)
	server.conn = conn
	return server
}

func (tr *TrackerServer) Close() {
	if tr.conn != nil {
		CloseConnection(tr.conn)
	}
}

func (tr *TrackerServer) GetConn() *net.TCPConn {
	return tr.conn
}
