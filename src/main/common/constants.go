package common

const (
	FdfsProtoCmdQuit                                byte = 82
	TrackerProtoCmdServerListGroup                  byte = 91
	TrackerProtoCmdServerListStorage                byte = 92
	TrackerProtoCmdServerDeleteStorage              byte = 93
	TrackerProtoCmdServiceQueryStoreWithoutGroupOne byte = 101
	TrackerProtoCmdServiceQueryFetchOne             byte = 102
	TrackerProtoCmdServiceQueryUpdate               byte = 103
	TrackerProtoCmdServiceQueryStoreWithGroupOne    byte = 104
	TrackerProtoCmdServiceQueryFetchAll             byte = 105
	TrackerProtoCmdServiceQueryStoreWithoutGroupAll byte = 106
	TrackerProtoCmdServiceQueryStoreWithGroupAll    byte = 107
	TrackerProtoCmdResp                             byte = 100
	FdfsProtoCmdActiveTest                          byte = 111
	StorageProtoCmdUploadFile                       byte = 11
	StorageProtoCmdDeleteFile                       byte = 12
	StorageProtoCmdSetMetadata                      byte = 13
	StorageProtoCmdDownloadFile                     byte = 14
	StorageProtoCmdGetMetadata                      byte = 15
	StorageProtoCmdUploadSlaveFile                  byte = 21
	StorageProtoCmdQueryFileInfo                    byte = 22
	StorageProtoCmdUploadAppenderFile               byte = 23 //create appender file
	StorageProtoCmdAppendFile                       byte = 24 //append file
	StorageProtoCmdModifyFile                       byte = 34 //modify appender file
	StorageProtoCmdTruncateFile                     byte = 36 //truncate appender file
	StorageProtoCmdResp                             byte = TrackerProtoCmdResp
	FdfsStorageStatusInit                           byte = 0
	FdfsStorageStatusWaitSync                       byte = 1
	FdfsStorageStatusSyncing                        byte = 2
	FdfsStorageStatusIpChanged                      byte = 3
	FdfsStorageStatusDeleted                        byte = 4
	FdfsStorageStatusOffline                        byte = 5
	FdfsStorageStatusOnline                         byte = 6
	FdfsStorageStatusActive                         byte = 7
	FdfsStorageStatusNone                           byte = 99
	/**
	 * for overwrite all old metadata
	 */
	StorageSetMetadataFlagOverwrite byte = 'O'
	/**
	 * for replace, insert when the meta item not exist, otherwise update it
	 */
	StorageSetMetadataFlagMerge     byte   = 'M'
	FdfsProtoPkgLenSize             int    = 8
	FdfsProtoCmdSize                int    = 1
	FdfsGroupNameMaxLen             int    = 16
	FdfsIpaddrSize                  int    = 16
	FdfsDomainNameMaxSize           int    = 128
	FdfsVersionSize                 int    = 6
	FdfsStorageIdMaxSize            int    = 16
	FdfsRecordSeparator             string = "\u0001"
	DfsFieldSeparator               string = "\u0002"
	TrackerQueryStorageFetchBodyLen        = FdfsGroupNameMaxLen + FdfsIpaddrSize - 1 + FdfsProtoPkgLenSize
	TrackerQueryStorageStoreBodyLen        = FdfsGroupNameMaxLen + FdfsIpaddrSize + FdfsProtoPkgLenSize
	FdfsFileExtNameMaxLen           byte   = 6
	FdfsFilePrefixMaxLen            byte   = 16
	FdfsFilePathLen                 byte   = 10
	FdfsFilenameBase64Length        byte   = 27
	FdfsTrunkFileInfoLen            byte   = 16
	ErrNoEnoent                     byte   = 2
	ErrNoEio                        byte   = 5
	ErrNoEbusy                      byte   = 16
	ErrNoEinval                     byte   = 22
	ErrNoEnospc                     byte   = 28
	ECONNREFUSED                    byte   = 61
	ErrNoEalready                   byte   = 114
	InfiniteFileSize                int64  = 256 * 1024 * 1024 * 1024 * 1024 * 1024
	AppenderFileSize                int64  = InfiniteFileSize
	TrunkFileMarkSize               int64  = 512 * 1024 * 1024 * 1024 * 1024 * 1024
	NormalLogicFilenameLength       int64  = int64(FdfsFilePathLen + FdfsFilenameBase64Length + FdfsFileExtNameMaxLen + 1)
	TrunkLogicFilenameLength        int64  = NormalLogicFilenameLength + int64(FdfsTrunkFileInfoLen)
	ProtoHeaderCmdIndex             int    = FdfsProtoPkgLenSize
	ProtoHeaderStatusIndex          int    = FdfsProtoPkgLenSize + 1

	//config
	ConfKeyConnectTimeout      = "connect_timeout"
	ConfKeyNetworkTimeout      = "network_timeout"
	ConfKeyCharset             = "charset"
	ConfKeyHttpAntiStealToken  = "http.anti_steal_token"
	ConfKeyHttpSecretKey       = "http.secret_key"
	ConfKeyHttpTrackerHttpPort = "http.tracker_http_port"
	ConfKeyTrackerServer       = "tracker_server"

	DefaultConnectTimeout      = 5  //second
	DefaultNetworkTimeout      = 30 //second
	DefaultCharset             = "UTF-8"
	DefaultHttpTrackerHttpPort = 80
)
