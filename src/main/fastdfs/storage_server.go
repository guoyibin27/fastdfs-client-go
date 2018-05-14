package fastdfs

import "github.com/fastdfs-client-go/src/main/common"

type StorageServer struct {
	trackerServer  *TrackerServer
	storePathIndex int
}

func NewStorageServer(addr string, storePath int) *StorageServer {
	storageServer := new(StorageServer)
	storageServer.storePathIndex = storePath

	if s, e := common.GetConnection(addr); e != nil {
		storageServer.trackerServer = NewTrackerServer(s)
	} else {
		common.CheckError(e)
	}

	return storageServer
}

func (ss *StorageServer) GetStorePathIndex() int {
	return ss.storePathIndex
}
