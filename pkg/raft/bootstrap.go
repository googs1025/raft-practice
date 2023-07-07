package raft

import (
	"k8s.io/klog"
	"net"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	c "github.com/practice/raft_practice/pkg/config"
)

var RaftNode *raft.Raft

// BootStrap 启动核心流程
func BootStrap(path string) error {

	// 1. 项目配置
	sysConfig, err := c.LoadConfig(path)
	if err != nil {
		klog.Errorf("load file err: %s", err)
		return err
	}
	c.SysConfig = sysConfig

	// 2. raft配置
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(sysConfig.ServerID)
	config.Logger = hclog.New(&hclog.LoggerOptions{
		Name:   sysConfig.ServerName,
		Level:  hclog.LevelFromString("DEBUG"),
		Output: os.Stderr,
	})

	// FIXME: 测试使用，需要看snapshot才设置的
	config.SnapshotInterval = time.Second * 5
	config.SnapshotThreshold = 2

	// 3. logStore保存配置
	dir, _ := os.Getwd()
	root := strings.Replace(dir, "\\", "/", -1)
	logStore, err := raftboltdb.NewBoltStore(root + sysConfig.LogStore)
	if err != nil {
		klog.Errorf("log store err: %s", err)
		return err
	}

	// 4. 保存节点信息
	stableStore, err := raftboltdb.NewBoltStore(root + sysConfig.StableStore)
	if err != nil {
		klog.Errorf("stable file err: %s", err)
		return err
	}

	// 5. 保存文件快照
	//snapshotStore, err := raft.NewFileSnapshotStore(root+SysConfig.Snapshot, 1, nil)
	//if err != nil {
	// return err
	//}

	// 不使用内部snap功能
	snapshotStore := raft.NewDiscardSnapshotStore()

	// 6. 节点之间的通信
	addr, err := net.ResolveTCPAddr("tcp", sysConfig.Transport)
	transport, err := raft.NewTCPTransport(addr.String(), addr, 5, time.Second*10, os.Stdout)
	if err != nil {
		klog.Errorf("tcp transport err: %s", err)
		return err
	}

	// 自定义fsm
	fsm := NewMyFSM()
	//fsm := NewLocalFSM(root + "/" + sysConfig.LocalCache)

	RaftNode, err = raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		klog.Errorf("new raft err: %s", err)
		return err
	}

	servers := make([]raft.Server, 0) //改动点， 循环构建
	for _, s := range c.SysConfig.Servers {
		servers = append(servers, raft.Server{ID: s.ID, Address: s.Address})
	}

	configuration := raft.Configuration{
		Servers: servers,
	}

	RaftNode.BootstrapCluster(configuration)
	return nil
}
