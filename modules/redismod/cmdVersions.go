package redismod

var (
	cmdVersions = map[string]string{
		"append":                        "2.0.0",
		"auth":                          "1.0.0",
		"bgrewriteaof":                  "1.0.0",
		"bgsave":                        "1.0.0",
		"bitcount":                      "2.6.0",
		"bitfield":                      "3.2.0",
		"bitop":                         "2.6.0",
		"bitpos":                        "2.8.7",
		"blpop":                         "2.0.0",
		"brpop":                         "2.0.0",
		"brpoplpush":                    "2.2.0",
		"bzpopmin":                      "5.0.0",
		"bzpopmax":                      "5.0.0",
		"client id":                     "5.0.0",
		"client kill":                   "2.4.0",
		"client list":                   "2.4.0",
		"client pause":                  "2.9.50",
		"client reply":                  "3.2.0",
		"client setname":                "2.6.9",
		"cluster addslots":              "3.0.0",
		"cluster count-failure-reports": "3.0.0",
		"cluster countkeysinslot":       "3.0.0",
		"cluster delslots":              "3.0.0",
		"cluster failover":              "3.0.0",
		"cluster forget":                "3.0.0",
		"cluster getkeysinslot":         "3.0.0",
		"cluster info":                  "3.0.0",
		"cluster keyslot":               "3.0.0",
		"cluster meet":                  "3.0.0",
		"cluster nodes":                 "3.0.0",
		"cluster replicate":             "3.0.0",
		"cluster reset":                 "3.0.0",
		"cluster saveconfig":            "3.0.0",
		"cluster set-config-epoch":      "3.0.0",
		"cluster setslot":               "3.0.0",
		"cluster slaves":                "3.0.0",
		"cluster slots":                 "3.0.0",
		"command":                       "2.8.13",
		"command count":                 "2.8.13",
		"command getkeys":               "2.8.13",
		"command info":                  "2.8.13",
		"config get":                    "2.0.0",
		"config rewrite":                "2.8.0",
		"config set":                    "2.0.0",
		"config resetstat":              "2.0.0",
		"dbsize":                        "1.0.0",
		"debug object":                  "1.0.0",
		"debug segfault":                "1.0.0",
		"decr":                          "1.0.0",
		"decrby":                        "1.0.0",
		"del":                           "1.0.0",
		"discard":                       "2.0.0",
		"dump":                          "2.6.0",
		"echo":                          "1.0.0",
		"eval":                          "2.6.0",
		"evalsha":                       "2.6.0",
		"exec":                          "1.2.0",
		"exists":                        "1.0.0",
		"expire":                        "1.0.0",
		"expireat":                      "1.2.0",
		"flushall":                      "1.0.0",
		"flushdb":                       "1.0.0",
		"geoadd":                        "3.2.0",
		"geodist":                       "3.2.0",
		"geohash":                       "3.2.0",
		"geopos":                        "3.2.0",
		"georadius":                     "3.2.0",
		"georadiusbymember":             "3.2.0",
		"get":                           "1.0.0",
		"getbit":                        "2.2.0",
		"getrange":                      "2.4.0",
		"getset":                        "1.0.0",
		"hdel":                          "2.0.0",
		"hexists":                       "2.0.0",
		"hget":                          "2.0.0",
		"hgetall":                       "2.0.0",
		"hincrby":                       "2.0.0",
		"hincrbyfloat":                  "2.6.0",
		"hkeys":                         "2.0.0",
		"hlen":                          "2.0.0",
		"hmget":                         "2.0.0",
		"hmset":                         "2.0.0",
		"hset":                          "2.0.0",
		"hsetnx":                        "2.0.0",
		"hvals":                         "2.0.0",
		"incr":                          "1.0.0",
		"incrby":                        "1.0.0",
		"incrbyfloat":                   "2.6.0",
		"info":                          "1.0.0",
		"keys":                          "1.0.0",
		"lastsave":                      "1.0.0",
		"lindex":                        "1.0.0",
		"linsert":                       "2.2.0",
		"llen":                          "1.0.0",
		"lpop":                          "1.0.0",
		"lpush":                         "1.0.0",
		"lpushx":                        "2.2.0",
		"lrange":                        "1.0.0",
		"lrem":                          "1.0.0",
		"lset":                          "1.0.0",
		"ltrim":                         "1.0.0",
		"mget":                          "1.0.0",
		"migrate":                       "2.6.0",
		"monitor":                       "1.0.0",
		"move":                          "1.0.0",
		"mset":                          "1.0.1",
		"msetnx":                        "1.0.1",
		"multi":                         "1.2.0",
		"object":                        "2.2.3",
		"persist":                       "2.2.0",
		"pexpire":                       "2.6.0",
		"pexpireat":                     "2.6.0",
		"pfadd":                         "2.8.9",
		"pfcount":                       "2.8.9",
		"pfmerge":                       "2.8.9",
		"ping":                          "1.0.0",
		"psetex":                        "2.6.0",
		"psubscribe":                    "2.0.0",
		"pttl":                          "2.6.0",
		"publish":                       "2.0.0",
		"pubsub":                        "2.8.0",
		"punsubscribe":                  "2.0.0",
		"quit":                          "1.0.0",
		"randomkey":                     "1.0.0",
		"readonly":                      "3.0.0",
		"readwrite":                     "3.0.0",
		"rename":                        "1.0.0",
		"renamenx":                      "1.0.0",
		"restore":                       "2.6.0",
		"role":                          "2.8.12",
		"rpop":                          "1.0.0",
		"rpoplpush":                     "1.2.0",
		"rpush":                         "1.0.0",
		"rpushx":                        "2.2.0",
		"sadd":                          "1.0.0",
		"save":                          "1.0.0",
		"scard":                         "1.0.0",
		"script debug":                  "3.2.0",
		"script exists":                 "2.6.0",
		"script flush":                  "2.8.0",
		"script kill":                   "2.6.0",
		"script load":                   "2.6.0",
		"sdiff":                         "1.0.0",
		"sdiffstore":                    "1.0.0",
		"select":                        "1.0.0",
		"set":                           "1.0.0",
		"setbit":                        "2.2.0",
		"setex":                         "2.0.0",
		"setnx":                         "1.0.0",
		"setrange":                      "2.2.0",
		"shutdown":                      "1.0.0",
		"sinter":                        "1.0.0",
		"sinterstore":                   "1.0.0",
		"sismember":                     "1.0.0",
		"slaveof":                       "1.0.0",
		"slowlog":                       "2.2.12",
		"smembers":                      "1.0.0",
		"smove":                         "1.0.0",
		"sort":                          "1.0.0",
		"spop":                          "1.0.0",
		"srandmember":                   "1.0.0",
		"srem":                          "1.0.0",
		"strlen":                        "2.2.0",
		"subscribe":                     "2.0.0",
		"sunion":                        "1.0.0",
		"sunionstore":                   "1.0.0",
		"swapdb":                        "4.0.0",
		"sync":                          "1.0.0",
		"time":                          "2.6.0",
		"touch":                         "3.2.1",
		"ttl":                           "1.0.0",
		"type":                          "1.0.0",
		"unsubscribe":                   "2.0.0",
		"unlink":                        "4.0.0",
		"wait":                          "3.0.0",
		"xack":                          "5.0.0",
		"xadd":                          "5.0.0",
		"xautoclaim":                    "6.2.0",
		"xclaim":                        "5.0.0",
		"xdel":                          "5.0.0",
		"xgroup":                        "5.0.0",
		"xinfo":                         "5.0.0",
		"xlen":                          "5.0.0",
		"xpending":                      "5.0.0",
		"xrange":                        "5.0.0",
		"xread":                         "5.0.0",
		"xreadgroup":                    "5.0.0",
		"xrevrange":                     "5.0.0",
		"xtrim":                         "5.0.0",
		"zadd":                          "1.2.0",
		"zcard":                         "1.2.0",
		"zcount":                        "2.0.0",
		"zincrby":                       "1.2.0",
		"zinterstore":                   "2.0.0",
		"zlexcount":                     "2.8.9",
		"zrange":                        "1.2.0",
		"zrangebylex":                   "2.8.9",
		"zrangebyscore":                 "1.0.5",
		"zrank":                         "2.0.0",
		"zrem":                          "1.2.0",
		"zremrangebylex":                "2.8.9",
		"zremrangebyrank":               "2.0.0",
		"zremrangebyscore":              "1.0.5",
		"zrevrange":                     "1.2.0",
		"zrevrangebylex":                "2.8.9",
		"zrevrangebyscore":              "2.0.0",
		"zrevrank":                      "2.0.0",
		"zscore":                        "1.2.0",
		"zunionstore":                   "2.0.0",
	}
)
