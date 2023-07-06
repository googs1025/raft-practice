package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	c "github.com/practice/raft_practice/pkg/cache"
	"log"
	"sync"
	"testing"
)

var cache *c.Bcache
var cacheOnce sync.Once

// local cache
func getCache() *c.Bcache {
	cacheOnce.Do(func() {
		cache = c.NewBcache("../tmp")
	})
	return cache
}

//测试 本地KV
func Benchmark_AddKV(t *testing.B) {
	for i := 0; i < t.N; i++ {
		key := fmt.Sprintf("key%d", i)
		value := key + "value"
		err := getCache().SetItem(key, value)
		if err != nil {
			log.Println("kv error:", err.Error())
		}
	}

}

var rcache redis.Conn
var rcacheOnce sync.Once

func getRedisCache() redis.Conn {
	rcacheOnce.Do(func() {
		c, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println("Connect to redis error", err)
			return
		}
		rcache = c
	})
	return rcache
}

// 测试 redis 插入
func Benchmark_AddRedis(t *testing.B) {
	for i := 0; i < t.N; i++ {
		key := fmt.Sprintf("key%d", i)
		value := key + "value"
		_, err := getRedisCache().Do("set", key, value)
		if err != nil {
			log.Println("redis set error", err)
		}
	}
}
