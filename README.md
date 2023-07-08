## raft_practice 简易型分布式kv存储服务


### 项目测试：
1. 进入项目根目录，启动三个进程 ./n1.sh ./n2.sh ./n3.sh
```bash
➜  raft_practice git:(main) ✗ ./n1.sh                                                                           
badger 2023/07/07 01:24:36 INFO: All 0 tables opened in 0s
badger 2023/07/07 01:24:36 INFO: Replaying file id: 0 at offset: 0
badger 2023/07/07 01:24:36 INFO: Replay took: 9.5µs
2023-07-07T01:24:36.503+0800 [INFO]  myraft-1: initial configuration: index=1 servers="[{Suffrage:Voter ID:1 Address:127.0.0.1:3001} {Suffrage:Voter ID:2 Address:127.0.0.1:3002} {Suffrage:Voter ID:3 Address:127.0.0.1:3003}]"
2023-07-07T01:24:36.504+0800 [INFO]  myraft-1: entering follower state: follower="Node at 127.0.0.1:3001 [Follower]" leader-address= leader-id=
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /get                      --> github.com/practice/raft_practice/pkg/raft.CacheServer.func1 (2 handlers)
[GIN-debug] POST   /set                      --> github.com/practice/raft_practice/pkg/raft.CacheServer.func2 (2 handlers)
[GIN-debug] DELETE /delete                   --> github.com/practice/raft_practice/pkg/raft.CacheServer.func3 (2 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.

```

2. 调用接口
```bash
1. set 存入缓存
2. get 获取缓存
3. delete 删除缓存
```