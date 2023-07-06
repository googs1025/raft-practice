package raft

import (
	"github.com/practice/raft_practice/pkg/config"
)

// IsLeader 判断自己是不是leader
func IsLeader() bool {
	if string(RaftNode.Leader()) == config.SysConfig.Transport {
		return true
	}
	return false
}

// GetLeaderHttp 得到leader的http地址
func GetLeaderHttp() string {
	for _, s := range config.SysConfig.Servers {
		if string(s.Address) == string(RaftNode.Leader()) {
			return s.Http
		}
	}
	return ""
}
