package fastdfs

import (
	"testing"
)

func TestTrackerGroup_GetConnectedTrackerServer(t *testing.T) {
	InitClientConfig("../../resources/fastdfs_client_config.json")
	trackerGroup := GetConfig().trackGroup
	trackerServer := trackerGroup.GetConnectedTrackerServer()
	trackerServer.Close()
}
