package cache

import (
	"fmt"
	"github.com/practice/raft_practice/pkg/common"
	"log"
	"testing"
	"time"
)

func TestLocalCacheSetGet(t *testing.T) {
	wd := common.GetWd()
	cache := NewBcache(wd + "/tmp")
	defer cache.Close()
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

func TestLocalCacheSetTTL(t *testing.T) {
	wd := common.GetWd()
	cache := NewBcache(wd + "/tmp")
	defer cache.Close()
	for i := 0; i < 10; i++ {
		err := cache.SetItemWithTTl(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i), time.Second*10)
		if err != nil {
			log.Println(err)
		} else {
			time.Sleep(time.Millisecond * 500)
			log.Println("key", i, "设置成功")
		}
	}
	for i := 0; i < 20; i++ {
		fmt.Println(cache.Keys(20))
		time.Sleep(time.Second * 1)
	}
}

func TestLocalCacheKeyPrefix(t *testing.T) {
	wd := common.GetWd()
	cache := NewBcache(wd + "/tmp")
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
