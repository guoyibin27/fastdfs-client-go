package fastdfs

import (
	"testing"
	"fmt"
	"os"
)

func TestTrackerClient_GetStorageServerWithGroup(t *testing.T) {
	InitClientConfig("../../resources/fastdfs_client_config.json")
	trackerGroup := GetConfig().trackGroup
	trackerServer := trackerGroup.GetConnectedTrackerServer()
	tc := NewTrackerClientWithSpecifiedGroup(*trackerGroup)
	storageServer := tc.GetStorageServer(trackerServer)
	fmt.Fprintln(os.Stderr, storageServer.GetStorePathIndex())
}

func TestTrackerClient_GetStorageServers(b *testing.T) {
	InitClientConfig("../../resources/fastdfs_client_config.json")
	trackerGroup := GetConfig().trackGroup
	trackerServer := trackerGroup.GetConnectedTrackerServer()
	tc := NewTrackerClientWithSpecifiedGroup(*trackerGroup)
	groupName := ""
	storageServers := tc.GetStorageServers(trackerServer, groupName)
	if len(storageServers) > 0 {
		ss := storageServers[0]
		fmt.Fprintln(os.Stderr, ss.GetStorePathIndex())
	} else {
		fmt.Fprintf(os.Stderr, "not found any storage server with %s", groupName)
	}
}
