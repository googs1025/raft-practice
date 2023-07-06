package raft

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/practice/raft_practice/pkg/cache"
	"net/http/httputil"
	"net/url"
	"time"
)

type CacheRequest struct {
	Key   string `json:"key" binding:"required,min=1"`
	Value string `json:"value" binding:"omitempty,min=1"`
}

func NewCacheRequest() *CacheRequest {
	return &CacheRequest{}
}

func CacheMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				context.JSON(400, gin.H{"message": e})
			}
		}()
		if IsLeader() {
			context.Next()
		} else {
			leaderHttp := GetLeaderHttp()
			addr, _ := url.Parse(leaderHttp)
			p := httputil.NewSingleHostReverseProxy(addr)
			p.ServeHTTP(context.Writer, context.Request)
			context.Abort()
		}

	}
}

func Error(err error) {
	if err != nil {
		panic(err)
	}

}

func CacheServer() *gin.Engine {
	r := gin.New()
	r.Use(CacheMiddleware())

	r.Handle("POST", "/get", func(context *gin.Context) {
		req := NewCacheRequest()
		Error(context.BindJSON(req))

		// 使用内存map缓存
		if v, err := cache.NewMapCache().GetItem(req.Key); v != "" && err == nil {
			context.JSON(200, req)
		} else {
			Error(fmt.Errorf("find no cache"))
		}

		// TODO 使用本地db缓存
		//if v, err := cache.GetBcache().GetItem(req.Key); err == nil {
		//	req.Value = v
		//	context.JSON(200, req)
		//} else {
		//	Error(fmt.Errorf("find no cache"))
		//}
	})

	r.Handle("POST", "/set", func(context *gin.Context) {
		req := NewCacheRequest()
		Error(context.BindJSON(req))

		// 应该在fsm中保存，不是在接口层
		// Set(req.Key,req.Value) //往我们的sync.Map里插值
		// context.JSON(200, gin.H{"message": "OK"})
		//bc := cache.GetBcache()
		//bc.SetItem(req.Key, req.Value)

		// 不需要额外执行set操作，直接使用raft能力
		reqBytes, _ := json.Marshal(req)
		future := RaftNode.Apply(reqBytes, time.Second)
		if e := future.Error(); e != nil {
			Error(e)
		} else {
			context.JSON(200, gin.H{"message": "OK"})
		}

	})

	return r
}
