package main

import (
	"flag"
	"log"

	"github.com/practice/raft_practice/pkg/config"
	"github.com/practice/raft_practice/pkg/raft"
)

var (
	configFile string
)

func main() {

	// ex: go run cmd/main.go --config n1.yaml
	flag.StringVar(&configFile, "config", "", "config file")
	flag.Parse()
	if configFile == "" {
		log.Fatal("config file error")
	}
	// raft流程
	err := raft.BootStrap(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// 启动 gin server
	raft.CacheServer().Run(":" + config.SysConfig.Port)

}
