package fastdfs

import (
	"strings"
	"github.com/fastdfs-client-go/src/main/common"
	"bytes"
	"strconv"
	"fmt"
	"os"
)

type TrackerClient struct {
	trackerGroup    *TrackerGroup
	lastCallErrorNo int
	errNo           byte
}

/**
   * constructor
   */
func NewTrackerClient() *TrackerClient {
	trackerClient := new(TrackerClient)
	trackerClient.trackerGroup = GetConfig().trackGroup
	return trackerClient
}

/**
   * constructor with specified tracker group
   */
func NewTrackerClientWithSpecifiedGroup(trackerGroup TrackerGroup) *TrackerClient {
	trackerClient := new(TrackerClient)
	trackerClient.trackerGroup = &trackerGroup
	trackerClient.lastCallErrorNo = 0
	return trackerClient
}

//get a connection to tracker server
func (tc *TrackerClient) GetConnection() *TrackerServer {
	return tc.trackerGroup.GetConnectedTrackerServer()
}

/**
   * get the error code of last call
   *
   * @return the error code of last call
   */
func (tc *TrackerClient) GetLastCallErrorNo() int {
	return tc.lastCallErrorNo
}

//query storage server to upload file
func (tc *TrackerClient) GetStorageServer(trackerServer *TrackerServer) *StorageServer {
	return tc.GetStorageServerWithGroup(trackerServer, "")
}

// query storage server for upload file
func (tc *TrackerClient) GetStorageServerWithGroup(trackerServer *TrackerServer, groupName string) *StorageServer {
	if trackerServer == nil {
		trackerServer = tc.GetConnection()
		if trackerServer == nil {
			return nil
		}
	}

	conn := trackerServer.conn
	defer conn.Close()
	var cmd byte
	var length int

	if strings.Compare(strings.TrimSpace(groupName), "") == 0 || len(strings.TrimSpace(groupName)) == 0 {
		cmd = common.TrackerProtoCmdServiceQueryStoreWithoutGroupOne
		length = 0
	} else {
		cmd = common.TrackerProtoCmdServiceQueryStoreWithGroupOne
		length = common.FdfsGroupNameMaxLen
	}

	header := common.PackHeader(cmd, int64(length), 0)
	conn.Write(header)

	//has group name
	if strings.Compare(strings.TrimSpace(groupName), "") != 0 && len(strings.TrimSpace(groupName)) > 0 {
		bs := []byte(groupName)
		groupNameByte := make([]byte, common.FdfsGroupNameMaxLen)
		copy(groupNameByte, bs)
		conn.Write(groupNameByte)
	}

	pkg := common.ParseReceivePackage(*conn, int(common.TrackerProtoCmdResp), int64(common.TrackerQueryStorageStoreBodyLen))
	if pkg.ErrorNo != 0 {
		return nil
	}

	ipAddress := strings.TrimSpace(string(pkg.Body[common.FdfsGroupNameMaxLen : common.FdfsIpaddrSize-2+common.FdfsGroupNameMaxLen]))
	port := common.Buff2Int64(pkg.Body, common.FdfsGroupNameMaxLen+common.FdfsIpaddrSize-1)
	storePath := pkg.Body[common.TrackerQueryStorageStoreBodyLen-1]

	var buf bytes.Buffer
	buf.WriteString(ipAddress)
	buf.WriteString(":")
	buf.WriteString(strconv.FormatInt(int64(port), 10))
	return NewStorageServer(buf.String(), int(storePath))
}

// query storage server to upload file
func (tc *TrackerClient) GetStorageServers(trackerServer *TrackerServer, groupName string) []*StorageServer {
	if trackerServer == nil {
		trackerServer = tc.GetConnection()
		if trackerServer == nil {
			return nil
		}
	}

	conn := trackerServer.conn
	defer conn.Close()
	var cmd byte
	var length int

	if strings.Compare(strings.TrimSpace(groupName), "") == 0 || len(strings.TrimSpace(groupName)) == 0 {
		cmd = common.TrackerProtoCmdServiceQueryStoreWithoutGroupAll
		length = 0
	} else {
		cmd = common.TrackerProtoCmdServiceQueryStoreWithGroupAll
		length = common.FdfsGroupNameMaxLen
	}

	header := common.PackHeader(cmd, int64(length), 0)
	conn.Write(header)

	//has group name
	if strings.Compare(strings.TrimSpace(groupName), "") != 0 && len(strings.TrimSpace(groupName)) > 0 {
		bs := []byte(groupName)
		groupNameByte := make([]byte, common.FdfsGroupNameMaxLen)
		copy(groupNameByte, bs)
		conn.Write(groupNameByte)
	}

	pkg := common.ParseReceivePackage(*conn, int(common.TrackerProtoCmdResp), -1)
	tc.errNo = pkg.ErrorNo
	if pkg.ErrorNo != 0 {
		return nil
	}

	if len(pkg.Body) < common.TrackerQueryStorageStoreBodyLen {
		tc.errNo = common.ErrNoEinval
		return nil
	}

	ipPortLen := len(pkg.Body) - common.FdfsGroupNameMaxLen - 1
	recordLen := common.FdfsIpaddrSize - 1 + common.FdfsProtoPkgLenSize

	if ipPortLen%recordLen != 0 {
		tc.errNo = common.ErrNoEinval
		return nil;
	}

	serverCount := ipPortLen / recordLen
	if serverCount > 16 {
		tc.errNo = common.ErrNoEnospc
		return nil;
	}

	result := make([]*StorageServer, serverCount)

	storePath := pkg.Body[len(pkg.Body)-1];
	offset := common.FdfsGroupNameMaxLen

	for i := 0; i < serverCount; i++ {
		ipAddress := strings.TrimSpace(string(pkg.Body[offset : common.FdfsIpaddrSize-2+offset]))
		offset += common.FdfsIpaddrSize - 1

		port := common.Buff2Int64(pkg.Body, offset);
		offset += common.FdfsProtoPkgLenSize

		var buf bytes.Buffer
		buf.WriteString(ipAddress)
		buf.WriteString(":")
		buf.WriteString(strconv.FormatInt(int64(port), 10))
		result[i] = NewStorageServer(buf.String(), int(storePath))
	}
	return result
}

//query storage server to download file
func (tc *TrackerClient) getStorages(trackerServer *TrackerServer, cmd byte, groupName string, filename string) []*ServerInfo {
	if trackerServer == nil {
		trackerServer = tc.GetConnection()
		if trackerServer == nil {
			return nil
		}
	}

	conn := trackerServer.conn
	defer conn.Close()
	bs := []byte(groupName)
	groupNameByte := make([]byte, common.FdfsGroupNameMaxLen)
	filenameByte := []byte(filename)
	copy(groupNameByte, bs)
	header := common.PackHeader(cmd, int64(common.FdfsGroupNameMaxLen+len(filenameByte)), 0)
	wholePkg := make([]byte, len(header)+len(groupNameByte), len(filenameByte))
	var buffer bytes.Buffer
	buffer.Write(header)
	buffer.Write(groupNameByte)
	buffer.Write(filenameByte)
	copy(wholePkg, buffer.Bytes())
	conn.Write(wholePkg)

	pkgInfo := common.ParseReceivePackage(*conn, int(common.TrackerProtoCmdResp), -1)
	tc.errNo = pkgInfo.ErrorNo
	if pkgInfo.ErrorNo != 0 {
		return nil
	}

	if len(pkgInfo.Body) < common.TrackerQueryStorageFetchBodyLen {
		fmt.Fprintf(os.Stderr, "Invalid body legth: %d", len(pkgInfo.Body))
		return nil
	}

	if (len(pkgInfo.Body)-common.TrackerQueryStorageFetchBodyLen)%(common.FdfsIpaddrSize-1) != 0 {
		fmt.Fprintf(os.Stderr, "Invalid body length: %d", len(pkgInfo.Body))
	}

	serverCount := 1 + (len(pkgInfo.Body)-common.TrackerQueryStorageFetchBodyLen)/(common.FdfsIpaddrSize-1)
	ipAddr := strings.TrimSpace(string(pkgInfo.Body[common.FdfsGroupNameMaxLen:(common.FdfsGroupNameMaxLen + common.FdfsIpaddrSize - 1)]))
	offset := common.FdfsGroupNameMaxLen + common.FdfsIpaddrSize - 1

	port := common.Buff2Int64(pkgInfo.Body, offset)
	offset += common.FdfsProtoPkgLenSize

	servers := make([]*ServerInfo, serverCount)
	servers[0] = NewServerInfo(fmt.Sprintf("%s:%d", ipAddr, int(port)))

	for i := 1; i < serverCount; i++ {
		ipAddr = strings.TrimSpace(string(pkgInfo.Body[offset:(offset + common.FdfsIpaddrSize - 1)]))
		servers[i] = NewServerInfo(fmt.Sprintf("%s:%d", ipAddr, int(port)))
		offset += common.FdfsIpaddrSize - 1
	}
	return servers
}

//query storage server to download file
func (tc *TrackerClient) GetFetchStorage(trackerServer *TrackerServer, groupName string, fileName string) *StorageServer {
	servers := tc.getStorages(trackerServer, common.TrackerProtoCmdServiceQueryFetchOne, groupName, fileName)
	if servers == nil {
		return nil
	} else {
		return NewStorageServer(servers[0].ipAddr, 0)
	}
}


func (tc *TrackerClient) GetFetchStorages(trackerServer *TrackerServer, groupName string, filename string) []*ServerInfo {
	return tc.getStorages(trackerServer, common.TrackerProtoCmdServiceQueryFetchAll, groupName, filename)
}