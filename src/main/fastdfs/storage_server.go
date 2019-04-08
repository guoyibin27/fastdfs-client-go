package fastdfs

type StorageServer struct {
	trackerServer  *TrackerServer
	storePathIndex int
}

func NewStorageServer(addr string, storePath int) *StorageServer {
	storageServer := new(StorageServer)
	storageServer.storePathIndex = storePath

	if s, e := GetConnection(addr); e != nil {
		storageServer.trackerServer = NewTrackerServer(s)
	} else {
		CheckError(e)
	}

	return storageServer
}

func (ss *StorageServer) GetStorePathIndex() int {
	return ss.storePathIndex
}
