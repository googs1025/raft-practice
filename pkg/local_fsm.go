package pkg

import (
	"encoding/json"
	"github.com/hashicorp/raft"
	"io"
)

type LocalFSM struct {
}

//真正 持久化数据
func (lf *LocalFSM) Apply(log *raft.Log) interface{} {
	req := NewCacheRequest()
	err := json.Unmarshal(log.Data, req)
	if err != nil {
		return err
	}
	return LocalCache.SetItem(req.Key, req.Value)

}
func (lf *LocalFSM) Snapshot() (snapshot raft.FSMSnapshot, err error) {
	return nil, nil

}
func (lf *LocalFSM) Restore(reader io.ReadCloser) error {
	return nil
}
