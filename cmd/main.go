package main

import (
	"flag"
	"github.com/practice/raft_practice/pkg"
	"log"
)

func main() {

	cfile := ""
	// ex: go run cmd/main.go -c n1.yaml
	flag.StringVar(&cfile, "c", "", "config file ")
	flag.Parse()
	if cfile == "" {
		log.Fatal("config file error")
	}
	// raft流程
	err := pkg.BootStrap(cfile)
	if err != nil {
		log.Fatal(err)
	}

	//for {
	//	fmt.Println("master node:", pkg.RaftNode.Leader())
	//	time.Sleep(time.Second * 1)
	//}

	// 启动 gin server
	pkg.CacheServer().Run(":" + pkg.SysConfig.Port)

}
