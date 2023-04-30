package pkg

import (
	"encoding/json"
	"github.com/hashicorp/raft"
	"io"
	"log"
)

type MyFSM struct {
}

// Apply 持久化数据方法
func (m *MyFSM) Apply(log *raft.Log) interface{} {
	req := NewCacheRequest()
	err := json.Unmarshal(log.Data, req)
	Set(req.Key, req.Value) // 数据保存
	return err
}

// Snapshot 实现存储快照的逻辑
func (m *MyFSM) Snapshot() (raft.FSMSnapshot, error) {
	return NewMySnapshot(), nil
}

// Restore 快照恢复逻辑
func (m *MyFSM) Restore(reader io.ReadCloser) error {
	err := json.NewDecoder(reader).Decode(getCache())
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
	b, err := json.Marshal(getCache()) //写json快照
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
