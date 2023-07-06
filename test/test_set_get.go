package main

import (
	"fmt"
	"github.com/practice/raft_practice/pkg"
	"log"
)

func main() {

	cache := pkg.NewBcache("tmp")
	err := cache.SetItem("aaa", "dddd")
	if err != nil {
		log.Println(err)
		return
	}
	getValue, err := cache.GetItem("aaa")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(getValue)

}
