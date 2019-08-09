package fastdfs

type StorageClient struct {
	storageServer *StorageServer
	trackerServer *TrackerServer
}

func NewStorageClient() *StorageClient {
	return NewStorageClientWithGivenParam(nil, nil)
}

func NewStorageClientWithGivenParam(trackerServer *TrackerServer, storageServer *StorageServer) *StorageClient {
	storageClient := new(StorageClient)
	storageClient.trackerServer = trackerServer
	storageClient.storageServer = storageServer
	return storageClient
}
