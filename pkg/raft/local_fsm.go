package raft

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/hashicorp/raft"
	"github.com/practice/raft_practice/pkg/cache"
)

type LocalFSM struct {
	CacheI cache.Cache
}

var (
	dOnce sync.Once
)

func getOnceBcache() {
	dOnce.Do(func() {

	})
}

func NewLocalFSM(path string) *LocalFSM {
	m := &LocalFSM{
		CacheI: cache.NewBcache(path),
	}
	return m
}

// Apply 持久化数据
func (lf *LocalFSM) Apply(log *raft.Log) interface{} {
	req := NewCacheRequest()
	err := json.Unmarshal(log.Data, req)
	if err != nil {
		return err
	}
	return cache.LocalCache.SetItem(req.Key, req.Value)
}
func (lf *LocalFSM) Snapshot() (snapshot raft.FSMSnapshot, err error) {
	return nil, nil

}
func (lf *LocalFSM) Restore(reader io.ReadCloser) error {
	return nil
}
