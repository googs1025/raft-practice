package pkg

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	c "golanglearning/new_project/raft_practice/pkg/config"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var RaftNode *raft.Raft
var SysConfig *c.Config

func BootStrap(path string) error {
	sysConfig, err := c.LoadConfig(path)
	if err != nil {
		return err
	}
	SysConfig = sysConfig
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(sysConfig.ServerID)
	config.Logger = hclog.New(&hclog.LoggerOptions{
		Name:   sysConfig.ServerName,
		Level:  hclog.LevelFromString("DEBUG"),
		Output: os.Stderr,
	})

	config.SnapshotInterval = time.Second * 5
	config.SnapshotThreshold = 2

	//logStore保存配置
	dir, _ := os.Getwd()
	root := strings.Replace(dir, "\\", "/", -1)
	logStore, err := raftboltdb.NewBoltStore(root + sysConfig.LogStore)
	if err != nil {
		return err
	}

	//保存节点信息
	stableStore, err := raftboltdb.NewBoltStore(root + sysConfig.StableStore)
	if err != nil {
		return err
	}

	// 保存文件快照
	snapshotStore, err := raft.NewFileSnapshotStore(root+SysConfig.Snapshot, 1, nil)
	if err != nil {
		return err
	}

	// 节点之间的通信
	addr, err := net.ResolveTCPAddr("tcp", sysConfig.Transport)
	transport, err := raft.NewTCPTransport(addr.String(), addr, 5, time.Second*10, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	fsm := &MyFSM{}

	RaftNode, err = raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return err
	}

	servers := make([]raft.Server, 0) //改动点， 循环构建
	for _, s := range SysConfig.Servers {
		servers = append(servers, raft.Server{ID: s.ID, Address: s.Address})
	}

	configuration := raft.Configuration{
		Servers: servers,
	}

	RaftNode.BootstrapCluster(configuration)
	return nil
}
