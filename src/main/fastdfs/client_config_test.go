package fastdfs

import (
	"testing"
	"fmt"
)

func TestInitClientConfig(t *testing.T) {
	InitClientConfig("../../resources/fastdfs_client_config.json")
	PrintConfigInfo()
}

func TestGetConfig(t *testing.T) {
	InitClientConfig("../../resources/fastdfs_client_config.json")
	fmt.Printf("%d", GetConfig().connectTimeout)
	fmt.Printf("%s", GetConfig().secretKey)
	fmt.Printf("%s", GetConfig().trackGroup)
}
