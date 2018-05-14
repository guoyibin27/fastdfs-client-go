package common

import (
	"fmt"
	"os"
	"net"
	"crypto/md5"
	"encoding/hex"
)

type Header struct {
	errorNo       byte
	contentLength int64
}

type Package struct {
	ErrorNo byte
	Body    []byte
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		panic(err)
	}
}

func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func CloseConnection(conn *net.TCPConn) {
	header := PackHeader(FdfsProtoCmdQuit, 0, 0);
	conn.Write(header)
	defer conn.Close()
}

func PackHeader(cmd byte, pkgLength int64, errorNo byte) []byte {
	var header = new([FdfsProtoPkgLenSize + 2]byte)
	var hexLength = Int2Buff(pkgLength)
	copy(hexLength[:], header[:])
	header[ProtoHeaderCmdIndex] = cmd
	header[ProtoHeaderStatusIndex] = errorNo
	return header[:]
}

func Int2Buff(n int64) [8]byte {
	var bs [8]byte
	bs[0] = (byte)((n >> 56) & 0xFF)
	bs[1] = (byte)((n >> 48) & 0xFF)
	bs[2] = (byte)((n >> 40) & 0xFF)
	bs[3] = (byte)((n >> 32) & 0xFF)
	bs[4] = (byte)((n >> 24) & 0xFF)
	bs[5] = (byte)((n >> 16) & 0xFF)
	bs[6] = (byte)((n >> 8) & 0xFF)
	bs[7] = (byte)(n & 0xFF)
	return bs
}

func Buff2Int64(bs []byte, offset int) uint8 {
	return (If(bs[offset] >= 0, bs[offset], 256+int64(bs[offset])).(uint8) << 56) |
		(If(bs[offset+1] >= 0, bs[offset+1], 256+int64(bs[offset+1])).(uint8) << 48) |
		(If(bs[offset+2] >= 0, bs[offset+2], 256+int64(bs[offset+2])).(uint8) << 40) |
		(If(bs[offset+3] >= 0, bs[offset+3], 256+int64(bs[offset+3])).(uint8) << 32) |
		(If(bs[offset+4] >= 0, bs[offset+4], 256+int64(bs[offset+4])).(uint8) << 24) |
		(If(bs[offset+5] >= 0, bs[offset+5], 256+int64(bs[offset+5])).(uint8) << 16) |
		(If(bs[offset+6] >= 0, bs[offset+6], 256+int64(bs[offset+6])).(uint8) << 8) |
		(If(bs[offset+7] >= 0, bs[offset+7], 256+int64(bs[offset+7])).(uint8))
}

func Buff2Int(bs []byte, offset int) int {
	return (If(bs[offset] >= 0, bs[offset], 256+int(bs[offset])).(int) << 24) |
		(If(bs[offset+1] >= 0, bs[offset+1], 256+int(bs[offset+1])).(int) << 16) |
		(If(bs[offset+2] >= 0, bs[offset+2], 256+int(bs[offset+2])).(int) << 8) |
		(If(bs[offset+3] >= 0, bs[offset+3], 256+int(bs[offset+3])).(int))
}

// md5 crypto
func Md5(origin string) string {
	crypto := md5.New()
	crypto.Write([]byte(origin))
	ret := crypto.Sum(nil)
	return hex.EncodeToString(ret)
}

func GetConnection(host string) (c *net.TCPConn, e error) {
	addr := resolveAddr(host)
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}
	conn.SetKeepAlive(true)
	return conn, nil
}

//parse received header info
func ParseReceiveHeader(conn net.TCPConn, expectCmd int, expectContentLength int64) Header {
	var header = new([ FdfsProtoPkgLenSize + 2]byte)

	count, err := conn.Read(header[:]);
	CheckError(err)
	if count != len(header) {
		fmt.Fprintf(os.Stderr, "received package size %d != %d", count, len(header))
	}

	if int(header[ProtoHeaderCmdIndex]) != expectCmd {
		fmt.Fprintf(os.Stderr, "received cmd: %d is not correct, expect cmd: %d", header[ProtoHeaderCmdIndex], expectCmd)
	}

	if header[ProtoHeaderStatusIndex] != 0 {
		return Header{header[ProtoHeaderStatusIndex], 0}
	}

	pkgLength := Buff2Int64(header[:], 0)

	if pkgLength < 0 {
		fmt.Fprintf(os.Stderr, "received package length %d < 0 !", pkgLength)
	}

	if expectContentLength > 0 && int64(pkgLength) != expectContentLength {
		fmt.Fprintf(os.Stderr, "received package length : %d is not correct, expect package length: %d", pkgLength, expectContentLength)
	}

	return Header{0, int64(pkgLength)}
}

// parser received package info
func ParseReceivePackage(conn net.TCPConn, expectCmd int, expectContentLength int64) Package {
	header := ParseReceiveHeader(conn, expectCmd, expectContentLength)
	if header.errorNo != 0 {
		return Package{header.errorNo, nil}
	}

	var body = make([]byte, header.contentLength)
	count, err := conn.Read(body[:])
	CheckError(err)
	if int64(count) != header.contentLength {
		fmt.Fprintf(os.Stderr, "received package size %d != %d", count, header.contentLength)
	}

	return Package{0, body[:]}
}

func resolveAddr(addr string) *net.TCPAddr {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	CheckError(err)
	return tcpAddr
}
