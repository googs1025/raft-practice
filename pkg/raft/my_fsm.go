package raft

import (
	"encoding/json"
	"io"
	"log"

	"github.com/hashicorp/raft"
	"github.com/practice/raft_practice/pkg/cache"
)

var _ raft.FSM = &MyFSM{}

// MyFSM raft中最重要的struct，
// 定义如何将数据持久化日志、如何存储快照、如何备份快照等操作
type MyFSM struct {
	CacheI  cache.Cache
}

func NewMyFSM() *MyFSM {
	m := &MyFSM{
		CacheI: cache.NewMapCache(),
	}
	return m
}

// Apply 持久化数据方法
func (m *MyFSM) Apply(log *raft.Log) interface{} {
	req := NewCacheRequest()
	err := json.Unmarshal(log.Data, req)
	m.CacheI.SetItem(req.Key, req.Value) // 数据保存
	return err
}

// Snapshot 实现存储快照的逻辑
func (m *MyFSM) Snapshot() (raft.FSMSnapshot, error) {
	return NewMySnapshot(), nil
}

// Restore 快照恢复逻辑
func (m *MyFSM) Restore(reader io.ReadCloser) error {
	err := json.NewDecoder(reader).Decode(cache.GetCache(cache.NewMapCache().(*cache.MapCache)))
	if err != nil {
		log.Println("restore error:", err)
		return err
	}
	return nil
}

var _ raft.FSMSnapshot = &MySnapshot{}

type MySnapshot struct{}

func NewMySnapshot() *MySnapshot {
	return &MySnapshot{}
}

func (ms *MySnapshot) Persist(sink raft.SnapshotSink) error {
	b, err := json.Marshal(cache.NewMapCache()) //写json快照
	if err != nil {
		return err
	} else {
		_, err := sink.Write(b)
		if err != nil {
			log.Println("snapshot write error: ", err)
			return sink.Cancel()
		}
	}
	return nil
}

// Release 保存后可执行的回调方法
func (ms *MySnapshot) Release() {
	log.Println("snapshot write successful...")
}
