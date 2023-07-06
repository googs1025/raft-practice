package main

import (
	"fmt"
	"github.com/practice/raft_practice/pkg"
	"log"
)

func main() {
	cache := pkg.NewBcache("tmp")
	err := cache.SetItem("/ddd/aaa", "dddd")
	_ = cache.SetItem("/ddd/ccc", "dddddsdd")
	_ = cache.SetItem("/ddd/ssss", "dddddsdd")

	if err != nil {
		log.Println(err)
		return
	}
	prefix, err := cache.KeysWithPrefix(10, "/ddd")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(prefix)
}
