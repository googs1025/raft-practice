package main

import (
	"flag"
	"golanglearning/new_project/raft_practice/pkg"
	"log"
)

func main() {

	cfile := ""
	flag.StringVar(&cfile, "c", "", "your config file ")
	flag.Parse()
	if cfile == "" {
		log.Fatal("config file error")
	}
	err := pkg.BootStrap(cfile)
	if err != nil {
		log.Fatal(err)
	}

	//for {
	//	fmt.Println("master node:", pkg.RaftNode.Leader())
	//	time.Sleep(time.Second * 1)
	//}

	// 启动gin server
	pkg.CacheServer().Run(":" + pkg.SysConfig.Port)

}
