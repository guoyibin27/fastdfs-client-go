package fastdfs

import (
	"github.com/fastdfs-client-go/src/main/common"
	"fmt"
	"strings"
)

type ClientConfig struct {
	connectTimeout  int
	networkTimeout  int
	charset         string
	antiStealToken  bool
	secretKey       string
	trackerHttpPort int
	trackGroup      *TrackerGroup
}

var cc ClientConfig

func InitClientConfig(filePath string) {
	reader := common.NewConfigReader()
	reader.LoadConfigFile(filePath)
	cc.connectTimeout = reader.GetIntValue(common.ConfKeyConnectTimeout, common.DefaultConnectTimeout) * 1000
	cc.networkTimeout = reader.GetIntValue(common.ConfKeyNetworkTimeout, common.DefaultNetworkTimeout) * 1000
	cc.charset = reader.GetStringValue(common.ConfKeyCharset, common.DefaultCharset)
	cc.antiStealToken = reader.GetBoolValue(common.ConfKeyHttpAntiStealToken, false)
	cc.trackerHttpPort = reader.GetIntValue(common.ConfKeyHttpTrackerHttpPort, common.DefaultHttpTrackerHttpPort)
	if cc.antiStealToken {
		cc.secretKey = reader.GetStringValue(common.ConfKeyHttpSecretKey, "")
	}

	trackerServers := reader.GetValues(common.ConfKeyTrackerServer)
	cc.trackGroup = NewTrackerGroup(trackerServers)
}

func PrintConfigInfo() {
	if &cc == nil {
		return
	}
	if cc.trackGroup == nil {
		return
	}

	trackerServers := cc.trackGroup.trackerServers
	fmt.Printf("{\n  connectTimeout(ms) = %d "+
		"\n  networkTimeout(ms) = %d"+
		"\n  charset = %s"+
		"\n  antiStealToken = %t"+
		"\n  secretKey = %s"+
		"\n  trackerHttpPort = %d"+
		"\n  trackerServerAddrs = %s "+
		"\n}", cc.connectTimeout, cc.networkTimeout, cc.charset, cc.antiStealToken,
		cc.secretKey, cc.trackerHttpPort, strings.Join(trackerServers, ","))
}

func GetConfig() ClientConfig {
	return cc
}
