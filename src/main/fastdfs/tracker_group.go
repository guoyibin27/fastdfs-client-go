package fastdfs

import (
	"fmt"
	"os"
	"sync"
)

type TrackerGroup struct {
	trackerServers     []string
	trackerServerIndex int
	mutex              sync.Mutex
}

/**
   * Constructor
   *
   * @param trackerServers tracker servers
   */
func NewTrackerGroup(trackServers []string) *TrackerGroup {
	trackGroup := new(TrackerGroup)
	trackGroup.trackerServerIndex = 0
	trackGroup.trackerServers = trackServers
	return trackGroup
}

/**
   * return connected tracker server with index
   *
   * @return connected tracker server, null for fail
   */
func (tr *TrackerGroup) getConnectedTrackerServerWithIndex(index int) (ts *TrackerServer, err error) {
	if conn, e := GetConnection(tr.trackerServers[index]); e != nil {
		return nil, e
	} else {
		return NewTrackerServer(conn), nil
	}
}

/**
   * return connected tracker server
   * @return connected tracker server, null for fail
   */
func (tr *TrackerGroup) GetConnectedTrackerServer() *TrackerServer {
	var currentIndex int

	tr.mutex.Lock()
	defer tr.mutex.Unlock()
	tr.trackerServerIndex ++
	if tr.trackerServerIndex >= len(tr.trackerServers) {
		tr.trackerServerIndex = 0
	}

	currentIndex = tr.trackerServerIndex

	if ts, err := tr.getConnectedTrackerServerWithIndex(currentIndex); err == nil {
		return ts
	} else {
		fmt.Fprintf(os.Stderr, "connect to server : %s  fail, cause : %s", tr.trackerServers[currentIndex], err.Error())
	}

	for index := 0; index < len(tr.trackerServers); index++ {
		if index == currentIndex {
			continue
		}

		if ts, err := tr.getConnectedTrackerServerWithIndex(index); err != nil {
			fmt.Fprintf(os.Stderr, "connect to server : %s  fail, cause : %s", tr.trackerServers[currentIndex], err.Error())
		} else {
			tr.mutex.Lock()
			if tr.trackerServerIndex == currentIndex {
				tr.trackerServerIndex = index
			}
			tr.mutex.Unlock()
			return ts
		}
	}

	return nil
}
